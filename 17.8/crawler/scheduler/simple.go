package scheduler

import (
	"github.com/Kirk-Wang/Hello-Gopher/17.8/crawler/engine"
)

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerReady(w chan engine.Request) {}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { s.workerChan <- r }()
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}
