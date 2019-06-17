package main

import (
	"github.com/Kirk-Wang/Hello-Gopher/re/14.6/crawler/engine"
	"github.com/Kirk-Wang/Hello-Gopher/re/14.6/crawler/scheduler"
	"github.com/Kirk-Wang/Hello-Gopher/re/14.6/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}

	e.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
