package main

import (
	"github.com/Kirk-Wang/Hello-Gopher/17.4/crawler/engine"
	"github.com/Kirk-Wang/Hello-Gopher/17.4/crawler/persist"
	"github.com/Kirk-Wang/Hello-Gopher/17.4/crawler/scheduler"
	"github.com/Kirk-Wang/Hello-Gopher/17.4/crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
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
