package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		result := <-out

		for _, item := range result.Items {
			itemCount++
			log.Printf("Got item #%d: %v", itemCount, item)
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(
	in chan Request,
	out chan ParseResult, s Scheduler) {
	go func() {
		for {
			// tell scheduler i'm ready
			s.WorkerReady(in)
			request := <-in
			parseResult, err := worker(request)
			if err != nil {
				continue
			}
			out <- parseResult
		}
	}()
}
