package main

import (
	"github.com/Kirk-Wang/Hello-Gopher/history/15.1/crawler/engine"
	"github.com/Kirk-Wang/Hello-Gopher/history/15.1/crawler/scheduler"
	"github.com/Kirk-Wang/Hello-Gopher/history/15.1/crawler/zhenai/parser"
)

func main() {

	// engine.SimpleEngine{}.Run(engine.Request{
	// 	Url:        "http://www.zhenai.com/zhenghun",
	// 	ParserFunc: parser.ParseCityList,
	// })
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}
