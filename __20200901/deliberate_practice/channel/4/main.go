package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for {
		// 为啥乱序
		// 收是按顺序收的
		// 但这两个 Print 是 IO 操作，goroutine 会进行调度
		fmt.Printf("Worker %d received %c\n", id, <-c)
	}
}

func chanDemo() {
	// 开 10 个 worker
	// 每个人都有一个 channel
	// 然后分别向它们分发
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Microsecond)
}

func main() {
	chanDemo()
}

/*
go run 4/main.go
*/
