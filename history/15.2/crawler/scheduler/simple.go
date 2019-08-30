package scheduler

import (
	"github.com/Kirk-Wang/Hello-Gopher/history/15.2/crawler/engine"
)

type SimpleScheduler struct {
	workerChan chan engine.Request
}

// 会改变 struct 里面的内容，因此得用个指针
func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	// s.workerChan <- r 因为是同步的，这样会要卡死，会产生循环等待
	// 开个 goroutine，这样就会总有人收
	// 这样就为每个 request 创建了一个 goroutine
	// 只做一件事，就是往 worker 统一的 channel 去分发 Request
	go func() { s.workerChan <- r }()
}
