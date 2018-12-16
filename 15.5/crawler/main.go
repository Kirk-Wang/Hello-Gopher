package main

import (
	"github.com/Kirk-Wang/Hello-Gopher/15.5/crawler/engine"
	"github.com/Kirk-Wang/Hello-Gopher/15.5/crawler/scheduler"
	"github.com/Kirk-Wang/Hello-Gopher/15.5/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		// Scheduler: &scheduler.SimpleScheduler{},
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
