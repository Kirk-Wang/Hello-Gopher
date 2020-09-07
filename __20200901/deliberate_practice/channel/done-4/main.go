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

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	// 十个管子里面放东西
	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	// 十个管子里面收东西
	// 大小写，分开等是没有问题的
	for _, worker := range workers {
		<-worker.done
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	// 分开等是没有问题的
	for _, worker := range workers {
		<-worker.done
	}
}

func main() {
	chanDemo()
}
