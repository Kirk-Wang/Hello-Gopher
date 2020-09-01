package main

import (
	"fmt"
)

func doWorker(id int, in chan int, done chan bool) {
	for n := range in {
		fmt.Printf("worker %d recieved %c\n", id, n)
		done <- true
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
	go doWorker(id, w.in, w.done)
	return w
}

func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for _, worker := range workers {
		<-worker.done
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	for _, worker := range workers {
		<-worker.done
	}
}

func main() {
	chanDemo()
}
