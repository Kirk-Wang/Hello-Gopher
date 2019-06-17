package main

import (
	"github.com/Kirk-Wang/Hello-Gopher/re/14.6/crawler/engine"
	"github.com/Kirk-Wang/Hello-Gopher/re/14.6/crawler/zhenai/parser"
)

func main() {
	engine.SimpleEngine{}.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
