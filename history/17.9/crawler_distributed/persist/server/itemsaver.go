package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Kirk-Wang/Hello-Gopher/history/17.9/crawler/config"
	"github.com/Kirk-Wang/Hello-Gopher/history/17.9/crawler_distributed/persist"
	"github.com/Kirk-Wang/Hello-Gopher/history/17.9/crawler_distributed/rpcsupport"
	"github.com/olivere/elastic"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	// 一旦挂了就强制退出
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
