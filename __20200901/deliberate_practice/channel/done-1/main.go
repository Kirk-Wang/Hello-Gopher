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

	// 发现所有的东西都是顺序打印的，这就尴尬了😂
	// 那还要并行的干啥建10个 worker?

	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
		// 收什么先别管，先收个东西进来再说
		<-workers[i].done
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
		<-workers[i].done
	}
}

func main() {
	chanDemo()
}
