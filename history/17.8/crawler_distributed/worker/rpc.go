package worker

import (
	"github.com/Kirk-Wang/Hello-Gopher/17.8/crawler/engine"
	"log"
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
	log.Printf("CrawlService %s", engineReq.Url)
	return nil
}
