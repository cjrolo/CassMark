package main

import (
	"flag"
	"github.com/gocql/gocql"
	"log"
	"os"
	"time"
)

const VERSION = "0.1 ALFA"

func DoQueryAndLog() {

}

func main() {
	var outFile, keyspace, hosts, u, p string
	flag.StringVar(&outFile, "out", "results.out", "Output")
	flag.StringVar(&keyspace, "k", "sandbox", "Keyspace to use")
	flag.StringVar(&hosts, "h", "127.0.0.1", "Hosts to Connect")
	flag.StringVar(&keyspace, "u", "", "User to connect to the Cluster")
	flag.StringVar(&keyspace, "p", "", "Password for the Cluster")
	flag.Parse()
	fl, err := os.OpenFile(outFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	HandleError(err)
	defer fl.Close()
	log.SetOutput(fl)
	log.Println("CassMark version: ", VERSION)
	if len(u) > 0 {
		conn := NewCassandraAuthConnector(keyspace, u, p, hosts...)
	} else {
		conn := NewCassandraConnector(keyspace, hosts...)
	}
}
