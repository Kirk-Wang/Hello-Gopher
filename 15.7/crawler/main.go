package main

import (
	"github.com/Kirk-Wang/Hello-Gopher/15.7/crawler/engine"
	"github.com/Kirk-Wang/Hello-Gopher/15.7/crawler/scheduler"
	"github.com/Kirk-Wang/Hello-Gopher/15.7/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}
	// e.Run(engine.Request{
	// 	Url:        "http://www.zhenai.com/zhenghun",
	// 	ParserFunc: parser.ParseCityList,
	// })
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})

}
