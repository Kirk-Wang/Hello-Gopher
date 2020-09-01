package main

import (
	"github.com/Kirk-Wang/Hello-Gopher/re/14.5/crawler/engine"
	"github.com/Kirk-Wang/Hello-Gopher/re/14.5/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
