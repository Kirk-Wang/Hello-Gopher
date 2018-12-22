package client

import (
	"fmt"
	"github.com/Kirk-Wang/Hello-Gopher/17.8/crawler/engine"
	"github.com/Kirk-Wang/Hello-Gopher/17.8/crawler_distributed/config"
	"github.com/Kirk-Wang/Hello-Gopher/17.8/crawler_distributed/rpcsupport"
	"github.com/Kirk-Wang/Hello-Gopher/17.8/crawler_distributed/worker"
)

func CreateProcessor() (engine.Processor, error) {
	client, err := rpcsupport.NewClient(
		fmt.Sprintf(":%d", config.WorkerPort0))
	if err != nil {
		return nil, err
	}

	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializedRequest(req)

		var sResult worker.ParseResult
		err := client.Call(config.CrawlServiceRpc, sReq, &sResult)

		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializedResult(sResult), nil
	}, nil
}
