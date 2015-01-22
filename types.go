package main

import (
	"github.com/gocql/gocql"
	"log"
)

var CONNECTOR *CassandraConnector

type CassandraObject interface {
	Insert() error
	Delete() error
}

type CassandraConnector struct {
	Cluster *gocql.ClusterConfig
	Session *gocql.Session
}

/*
   Any Consistency = 1 + iota
   One
   Two
   Three
   Quorum
   All
   LocalQuorum
   EachQuorum
   Serial
   LocalSerial
   LocalOne
*/
var CONSISTENCY = [...]string{"Any", "One", "Two", "Three", "Quorum", "All", "LocalQuorum", "EachQuorum", "Serial", "LocalSerial", "LocalOne"}

func NewCassandraAuthConnector(k string, u string, p string, hosts ...string) *CassandraConnector {
	cluster := gocql.NewCluster(hosts...)
	cluster.Keyspace = k
	cluster.Compressor = gocql.SnappyCompressor{}
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: u,
		Password: p,
	}
	session, err := cluster.CreateSession()
	cluster.RetryPolicy = &gocql.SimpleRetryPolicy{NumRetries: 2}
	if err != nil {
		log.Panicf("Cassandra Auth Connection error: %s", err)
	}
	return &CassandraConnector{
		Cluster: cluster,
		Session: session,
	}
}

func NewCassandraConnector(k string, hosts ...string) *CassandraConnector {
	cluster := gocql.NewCluster(hosts...)
	cluster.Keyspace = k
	cluster.Compressor = gocql.SnappyCompressor{}
	session, err := cluster.CreateSession()
	cluster.RetryPolicy = &gocql.SimpleRetryPolicy{NumRetries: 2}
	if err != nil {
		log.Panicf("Cassandra Connection error: %s", err)
	}
	return &CassandraConnector{
		Cluster: cluster,
		Session: session,
	}
}

func (conn *CassandraConnector) WriteWithConsistency() {
	for i, _ := range gocql.ConsistencyNames {
		if gocql.Consistency(i).String() == "default" {
			// This are only supported for writes!
			continue
		}
		log.Printf("Consistency %s", gocql.Consistency(i))
		query := conn.Session.Query(`INSERT INTO data (id, data) VALUES (?, ?)`,
			gocql.TimeUUID(),
			"hello world")
		query.Consistency(gocql.Consistency(i))
		err := query.Exec()
		if err != nil {
			log.Printf("Error executing query! %s", err)
			continue
		}
		lat := query.Latency()
		log.Printf("Query Successfull! Time: %d ms (%d ns)", int(lat/1000000), lat)
	}

}

func (conn *CassandraConnector) ReadWithConsistency() {
	// First write something we know so we can read
	uuid := gocql.TimeUUID()
	err := conn.Session.Query(`INSERT INTO data (id, data) VALUES (?, ?)`,
		uuid,
		"hello world").Exec()
	if err != nil {
		log.Fatalln("Cannot write to cluster! Read operation aborted...", err)
	}
	// Something to store the results
	var id gocql.UUID
	var text string

	for i, _ := range gocql.ConsistencyNames {
		if gocql.Consistency(i).String() == "default" || gocql.Consistency(i).String() == "any" || gocql.Consistency(i).String() == "eachquorum" {
			// This are only supported for writes!
			continue
		}
		log.Printf("Consistency %s", gocql.Consistency(i))
		query := conn.Session.Query(`SELECT id, data FROM data WHERE id = ? LIMIT 1`,
			uuid).Consistency(gocql.Consistency(i))
		err = query.Scan(&id, &text)
		if err != nil {
			log.Printf("Error executing query! %s", err)
			continue
		}
		lat := query.Latency()
		log.Printf("Query Successfull! Time: %d ms (%d ns)", int(lat/1000000), lat)
	}

}
