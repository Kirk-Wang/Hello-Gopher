package worker

import (
	"github.com/Kirk-Wang/Hello-Gopher/history/17.7/crawler/engine"
)

type CrawlService struct{}

func (CrawlService) Process(req Request, result *ParseResult) error {
	engineReq, err := DeserializedRequest(req)
	if err != nil {
		return err
	}
	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}

	*result = SerializedResult(engineResult)
	return nil
}
