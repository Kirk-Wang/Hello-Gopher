package main

import (
	"github.com/Kirk-Wang/Hello-Gopher/history/14.7/crawler/engine"
	"github.com/Kirk-Wang/Hello-Gopher/history/14.7/crawler/zhenai/parser"
)

func main() {

	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}
