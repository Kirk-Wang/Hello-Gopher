package main

import (
	"flag"
	"fmt"
	"github.com/Kirk-Wang/Hello-Gopher/17.9/crawler_distributed/rpcsupport"
	"github.com/Kirk-Wang/Hello-Gopher/17.9/crawler_distributed/worker"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), worker.CrawlService{}))
}
