package main

import (
	"fmt"
	"sync"
)

func doWorker(id int, in chan int, wg *sync.WaitGroup) {
	for n := range in {
		fmt.Printf("Worker %d received %c\n", id, n)
		wg.Done()
	}
}

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

// 自己建 channel
// 返回的 chanel 是干嘛用的呢，没错-->>是用来发数据的，送数据的
func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	go doWorker(id, w.in, wg)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	wg.Wait()
}

func main() {
	chanDemo()
}
