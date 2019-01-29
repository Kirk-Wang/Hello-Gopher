package main

import (
	"github.com/Kirk-Wang/Hello-Gopher/15.2/crawler/engine"
	"github.com/Kirk-Wang/Hello-Gopher/15.2/crawler/scheduler"
	"github.com/Kirk-Wang/Hello-Gopher/15.2/crawler/zhenai/parser"
)

func main() {

	// engine.SimpleEngine{}.Run(engine.Request{
	// 	Url:        "http://www.zhenai.com/zhenghun",
	// 	ParserFunc: parser.ParseCityList,
	// })
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}
