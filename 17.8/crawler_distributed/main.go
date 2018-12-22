package main

import (
	"fmt"
	"github.com/Kirk-Wang/Hello-Gopher/17.8/crawler/engine"
	"github.com/Kirk-Wang/Hello-Gopher/17.8/crawler/scheduler"
	"github.com/Kirk-Wang/Hello-Gopher/17.8/crawler/zhenai/parser"
	"github.com/Kirk-Wang/Hello-Gopher/17.8/crawler_distributed/config"
	itemsaver "github.com/Kirk-Wang/Hello-Gopher/17.8/crawler_distributed/persist/client"
	worker "github.com/Kirk-Wang/Hello-Gopher/17.8/crawler_distributed/worker/client"
)

func main() {
	itemChan, err := itemsaver.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort)) // 调 rpc, 让远程服务器去 save
	if err != nil {
		panic(err)
	}

	processor, err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})
}
