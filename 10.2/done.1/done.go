package main

import (
	"fmt"
	"sync"
)

// 不要通过共享内存来通信；通过通信来共享内存
// 这里加一个 channel 就好了
func doWork(id int, c chan int, wg *sync.WaitGroup) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
		// 注意，所有 channel 的发都是 block 的
		// 外面必须要有人收这个 done
		// done <- true
		// 让他不阻塞，简单的方法是-->我们又开个 goroutine
		wg.Done()
	}
}

type worker struct {
	in chan int
	wg *sync.WaitGroup // 我们必须使用外面wg, 不能说拷贝一份
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	go doWork(id, w.in, wg)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20) // 添加 20 个 task
	for i, worker := range workers {
		worker.in <- 'a' + i
		// wg.Add(1) 也可以
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	wg.Wait()
}

func main() {
	chanDemo()
}
