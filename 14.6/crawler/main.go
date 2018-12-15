package main

import (
	"github.com/Kirk-Wang/Hello-Gopher/14.5/crawler/engine"
	"github.com/Kirk-Wang/Hello-Gopher/14.5/crawler/zhenai/parser"
)

func main() {

	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}
