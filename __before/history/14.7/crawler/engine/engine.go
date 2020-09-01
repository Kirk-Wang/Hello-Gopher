package engine

import (
	"github.com/Kirk-Wang/Hello-Gopher/history/14.7/crawler/fetcher"
	"log"
)

// 可以送很多种子
func Run(seeds ...Request) {
	// 维护一个queue
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s", r.Url)
		// 对于每一个 request 进行 fetch
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
			continue
		}

		// parser
		parseResult := r.ParserFunc(body)
		// 展开一个个入队
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			// %v 啥都不转义
			log.Printf("Got item %v", item)
		}
	}
}
