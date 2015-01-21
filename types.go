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

//W_CONSISTENCY := [...]string{"ALL","EACH_QUORUM","QUORUM","LOCAL_QUORUM","ANY","LOCAL_ONE","ONE","TWO","THREE"}
//R_CONSISTENCY := [...]string{"ALL","EACH_QUORUM","QUORUM","LOCAL_QUORUM","LOCAL_ONE","ONE","TWO","THREE"}

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
		log.Panicf("Cassandra Connection error: %s", err)
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
	for i := range gocql.Consistency {
		log.Println(i)
	}

}

func ReadWithConsistency() {
	for i := range gocql.Consistency {
		log.Println(i)
	}

}
