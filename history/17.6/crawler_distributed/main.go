package main

import (
	"fmt"
	"github.com/Kirk-Wang/Hello-Gopher/17.6/crawler/engine"
	"github.com/Kirk-Wang/Hello-Gopher/17.6/crawler/scheduler"
	"github.com/Kirk-Wang/Hello-Gopher/17.6/crawler/zhenai/parser"
	"github.com/Kirk-Wang/Hello-Gopher/17.6/crawler_distributed/config"
	"github.com/Kirk-Wang/Hello-Gopher/17.6/crawler_distributed/persist/client"
)

func main() {
	itemChan, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort)) // 调 rpc, 让远程服务器去 save
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
