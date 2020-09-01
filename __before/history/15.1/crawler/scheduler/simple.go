package scheduler

import (
	"github.com/Kirk-Wang/Hello-Gopher/history/15.1/crawler/engine"
)

type SimpleScheduler struct {
	workerChan chan engine.Request
}

// 会改变 struct 里面的内容，因此得用个指针
func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	// send request down to worker chan
	s.workerChan <- r
}
