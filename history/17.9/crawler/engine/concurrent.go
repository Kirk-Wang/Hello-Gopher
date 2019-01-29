package engine

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan Item
	RequestProcessor Processor
}

type Processor func(Request) (ParseResult, error)

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
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
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
	for {
		result := <-out
		for _, item := range result.Items {
			// 脱手
			go func(item Item) { e.ItemChan <- item }(item)
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

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			// result, err := Worker(request)
			// Call RPC
			result, err := e.RequestProcessor(request)
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