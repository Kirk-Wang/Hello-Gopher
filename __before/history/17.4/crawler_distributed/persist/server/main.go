package main

import (
	"github.com/Kirk-Wang/Hello-Gopher/history/17.4/crawler_distributed/persist"
	"github.com/Kirk-Wang/Hello-Gopher/history/17.4/crawler_distributed/rpcsupport"
	"github.com/olivere/elastic"
	"log"
)

func main() {
	// 一旦挂了就强制退出
	log.Fatal(serveRpc(":1234", "dating_profile"))
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
