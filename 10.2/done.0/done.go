package main

import (
	"fmt"
)

// 不要通过共享内存来通信；通过通信来共享内存
// 这里加一个 channel 就好了
func doWork(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
		// 注意，所有 channel 的发都是 block 的
		// 外面必须要有人收这个 done
		// done <- true
		// 让他不阻塞，简单的方法是-->我们又开个 goroutine
		go func() { done <- true }()
	}
}

type worker struct {
	in   chan int
	done chan bool
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w.in, w.done)
	return w
}

func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		// 注意，所有 channel 的发都是 block 的
		// 发一个任务给它，另一端必须要有人收
		worker.in <- 'a' + i
		// <-workers[i].done-->> 这样没卵用^_^
	}

	for i, worker := range workers {
		// 没有人收的情况下，又给它发
		// 这就 deadlock，循环等待了
		worker.in <- 'A' + i
		// <-workers[i].done
	}

	// wait for all of them
	for _, worker := range workers {
		// 每个个 worker 发了两遍，所以收两遍
		<-worker.done
		<-worker.done
	}
}

func main() {
	chanDemo()
}
