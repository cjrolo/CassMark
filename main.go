package main

import (
	"flag"
	"fmt"
	_ "github.com/gocql/gocql"
	"log"
	"os"
	_ "time"
)

const VERSION = "0.4 BETA"

type stringslice []string

func (i *stringslice) String() string {
	return fmt.Sprintf("%s", *i)
}

func (i *stringslice) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var outFile, keyspace, u, p string
	var hosts stringslice
	var conn *CassandraConnector

	flag.StringVar(&outFile, "out", "results.out", "Output")
	flag.StringVar(&keyspace, "k", "sandbox", "Keyspace to use")
	flag.Var(&hosts, "h", "Hosts to Connect")
	flag.StringVar(&u, "u", "", "User to connect to the Cluster")
	flag.StringVar(&p, "p", "", "Password for the Cluster")
	flag.Parse()
	fl, err := os.OpenFile(outFile, os.O_CREATE|os.O_RDWR, 0666)
	HandleError(err)
	defer fl.Close()
	log.SetOutput(fl)
	log.Println("#CassMark version: ", VERSION)
	if len(u) > 0 {
		conn = NewCassandraAuthConnector(keyspace, u, p, hosts...)
	} else {
		conn = NewCassandraConnector(keyspace, hosts...)
	}
	log.Println("== WRITES ==")
	conn.WriteWithConsistency()
	log.Println("== READS ==")
	conn.ReadWithConsistency()
	log.Println("#DONE")
	// Cleanup
	conn.Session.Close()
}
