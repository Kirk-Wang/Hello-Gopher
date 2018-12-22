package client

import (
	"github.com/Kirk-Wang/Hello-Gopher/17.9/crawler/engine"
	"github.com/Kirk-Wang/Hello-Gopher/17.9/crawler_distributed/config"
	"github.com/Kirk-Wang/Hello-Gopher/17.9/crawler_distributed/worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {

	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializedRequest(req)

		var sResult worker.ParseResult
		// 每次做事，就从管道里面拿一个
		c := <-clientChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)

		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializedResult(sResult), nil
	}
}
