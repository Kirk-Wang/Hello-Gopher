package worker

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Kirk-Wang/Hello-Gopher/17.9/crawler/engine"
	"github.com/Kirk-Wang/Hello-Gopher/17.9/crawler/zhenai/parser"
	"github.com/Kirk-Wang/Hello-Gopher/17.9/crawler_distributed/config"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializedRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializedResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializedRequest(req))
	}
	return result
}

func DeserializedRequest(r Request) (engine.Request, error) {
	parser, err := deserializedParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

func DeserializedResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		engineReq, err := DeserializedRequest(req)
		if err != nil {
			log.Printf("error deserializing request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result
}

func deserializedParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseProfile:
		if args, ok := p.Args.(string); ok {
			var pargs parser.ProfileParams
			json.Unmarshal([]byte(args), &pargs)
			return parser.NewProfileParser(pargs.Name, pargs.Gender), nil
		} else {
			return nil, fmt.Errorf("invalid arg: %v", p.Args)
		}
	default:
		log.Printf("%v", p.Name)
		panic("xxx")
		return nil, errors.New("unknown parser name")
	}
}
