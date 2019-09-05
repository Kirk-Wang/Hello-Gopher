package main

import (
	"fmt"
	"github.com/Kirk-Wang/Hello-Gopher/history/17.6/crawler_distributed/config"
	"github.com/Kirk-Wang/Hello-Gopher/history/17.6/crawler_distributed/persist"
	"github.com/Kirk-Wang/Hello-Gopher/history/17.6/crawler_distributed/rpcsupport"
	"github.com/olivere/elastic"
	"log"
)

func main() {
	// 一旦挂了就强制退出
	log.Fatal(serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex))
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
