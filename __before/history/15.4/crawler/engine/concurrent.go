package engine

import (
	"log"
)

type ConcurrentEngine struct {
	// scheduler 哪里来？肚子里放一个就可以了
	Scheduler   Scheduler
	WorkerCount int
}

// 使用者来定义，你去实现
type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

// 自己也最好变成指针类型的接收者
func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	// 送进去
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(out, e.Scheduler)
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

func createWorker(out chan ParseResult, s Scheduler) {
	// 每个 channel 是自己的
	in := make(chan Request)
	go func() {
		for {
			// tell scheduler i'm ready
			// 可以通过这个 channel 做任务了
			s.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
