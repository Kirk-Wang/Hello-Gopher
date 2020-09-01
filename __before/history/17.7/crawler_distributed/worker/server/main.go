package main

import (
	"fmt"
	"github.com/Kirk-Wang/Hello-Gopher/history/17.7/crawler_distributed/config"
	"github.com/Kirk-Wang/Hello-Gopher/history/17.7/crawler_distributed/rpcsupport"
	"github.com/Kirk-Wang/Hello-Gopher/history/17.7/crawler_distributed/worker"
	"log"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", config.WorkerPort0), worker.CrawlService{}))
}
