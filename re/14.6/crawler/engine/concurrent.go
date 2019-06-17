package engine

type ConcurrentEngine struct{}

type Scheduler interface{
	Submit(Request)
}

func (ConcurrentEngine) Run(seeds ...Request) {
	for _, r := range seeds {
		scheduler.Submit(r)
	}
}
