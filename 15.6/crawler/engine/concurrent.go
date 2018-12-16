package engine

import (
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
		e.Scheduler.Submit(r)
	}

	// 收 <-
	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item #%d: %v", itemCount, item)
			itemCount++
		}

		for _, request := range result.Requests {
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
