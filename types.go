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

func WriteWithConsistency() {
	for i, cons := range CONSISTENCY {
		log.Println(i+1, cons)
	}

}

func ReadWithConsistency() {
	for i, cons := range CONSISTENCY {
		log.Println(i+1, cons)
	}

}
