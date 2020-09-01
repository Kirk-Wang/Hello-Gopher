package engine

import (
	"github.com/Kirk-Wang/Hello-Gopher/history/16.8/crawler/fetcher"
	"log"
)

func worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body, r.Url), nil
}
