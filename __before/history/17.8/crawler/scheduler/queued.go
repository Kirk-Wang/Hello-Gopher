package scheduler

import (
	"github.com/Kirk-Wang/Hello-Gopher/history/17.8/crawler/engine"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	// channel 套 channel, 100 个 worker, 每个都有自己的 channel
	// 把(chan engine.Request) 看成一个类型
	workerChan chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	// 每个 worker 有自己的 channel
	return make(chan engine.Request)
}

// 有一个 worker ready了，可以负责去接受 request 了
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

// 总控
func (s *QueuedScheduler) Run() {
	// 生成 channel
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request     //
			var activeWorker chan engine.Request // nil 是永远不会 select 到的
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			// 谁先来，先做谁
			select {
			case r := <-s.requestChan:
				// send r to a   ?worker---> 加入队列就好了
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				// send ?next_request to w---> 加入队列就好了
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}

	}()
}
