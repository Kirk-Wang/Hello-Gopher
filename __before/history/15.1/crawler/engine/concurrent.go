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
}

// 自己也最好变成指针类型的接收者
func (e *ConcurrentEngine) Run(seeds ...Request) {

	// 所有的 worker 共用一个输入
	in := make(chan Request)
	out := make(chan ParseResult)
	// 送进去
	e.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	// 将 Rquest 送入调度
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	// 死循环的收 <-
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item: %v", item)
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
