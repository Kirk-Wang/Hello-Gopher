package engine

import (
	"github.com/Kirk-Wang/Hello-Gopher/15.7/crawler/fetcher"
	"log"
)

type SimpleEngine struct{}

// 可以送很多种子
func (e SimpleEngine) Run(seeds ...Request) {
	// 维护一个queue
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}

		// 展开一个个入队
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			// %v 啥都不转义
			log.Printf("Got item %v", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	// 对于每一个 request 进行 fetch
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}
