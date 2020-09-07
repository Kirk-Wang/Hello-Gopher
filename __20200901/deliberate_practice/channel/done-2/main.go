package main

import (
	"fmt"
)

func doWorker(id int, in chan int, done chan bool) {
	for n := range in {
		fmt.Printf("Worker %d received %c\n", id, n)
		done <- true
	}
}

type worker struct {
	in   chan int
	done chan bool
}

// 自己建 channel
// 返回的 chanel 是干嘛用的呢，没错-->>是用来发数据的，送数据的
func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker(id, w.in, w.done)
	return w
}

func chanDemo() {
	// 开 10 个 worker
	// 每个人都有一个 channel
	// 然后分别向它们分发
	var workers [10]worker
	for i := 0; i < 10; i++ {
		// 建的 channel 把它存起来
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	// 第二组发现 deadlock
	// 所有 channel 的发都是 block 的(worker.in <- 'a' + i)
	// 因为先前的 done<-true 还没收（收的在后面）
	// 所以你再这个 worker 发，不就 deadlock 了吗~
	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	// wait for all of them
	// 不要发一个等一个，而是一口气发完这20个，然后等
	for _, worker := range workers {
		// 每个 worker 有两个任务，所以这里我等两次
		<-worker.done
		<-worker.done
	}
}

func main() {
	chanDemo()
}
