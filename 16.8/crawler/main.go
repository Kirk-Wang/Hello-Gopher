package main

import (
	"github.com/Kirk-Wang/Hello-Gopher/16.8/crawler/engine"
	"github.com/Kirk-Wang/Hello-Gopher/16.8/crawler/persist"
	"github.com/Kirk-Wang/Hello-Gopher/16.8/crawler/scheduler"
	"github.com/Kirk-Wang/Hello-Gopher/16.8/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
