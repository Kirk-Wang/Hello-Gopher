package main

import (
	"github.com/Kirk-Wang/Hello-Gopher/history/17.9/crawler/engine"
	"github.com/Kirk-Wang/Hello-Gopher/history/17.9/crawler/persist"
	"github.com/Kirk-Wang/Hello-Gopher/history/17.9/crawler/scheduler"
	"github.com/Kirk-Wang/Hello-Gopher/history/17.9/crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}

	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})
}
