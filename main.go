package main

import (
	"flag"
	"github.com/gocql/gocql"
	"log"
	"time"
    "os"
)

func DoQueryAndLog(){
    
}

func main() {
    var outFile, keyspace, hosts string
	flag.StringVar(&outFile, "out", "results.out", "Output")
    flag.StringVar(&keyspace, "k", "sanbox", "Keyspace to use")
	flag.StringVar(&hosts, "h", "127.0.0.1", "Hosts to Connect")
    flag.Parse()
	fl, err := os.OpenFile(outFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	HandleError(err)
	defer fl.Close()
	log.SetOutput(fl)
	log.Println("CassMark version: ", VERSION)
	conn = NewCassandraConnector(keyspace, hosts...)

}
