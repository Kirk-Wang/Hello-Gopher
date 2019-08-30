package engine

import (
	"github.com/Kirk-Wang/Hello-Gopher/history/15.7/crawler/model"
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	// 送进去
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		// 问调度器要 worker channel
		// 至于每人一个，还是共用一个, engine 不关心
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	// 将 Rquest 送入调度
	for _, r := range seeds {
		if isDuplicate(r.Url) {
			// log.Printf("Dupliacte request: %s", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}

	// 收 <-
	profileCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			if _, ok := item.(model.Profile); ok {
				log.Printf("Got Profile #%d: %v", profileCount, item)
				profileCount++
			}
		}
		// URL dedup 去重
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				// log.Printf("Dupliacte request: %s", request.Url)
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	// 没见过的，默认就是 false
	if visitedUrls[url] {
		return true
	}

	visitedUrls[url] = true
	return false
}
