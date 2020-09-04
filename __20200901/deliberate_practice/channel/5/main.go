package main

import (
	"fmt"
	"time"
)

// 自己建 channel
func createWorker(id int) chan int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received %c\n", id, <-c)
		}
	}()
	return c
}

func chanDemo() {
	// 开 10 个 worker
	// 每个人都有一个 channel
	// 然后分别向它们分发
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		// 建的 channel 把它存起来
		channels[i] = createWorker(i)
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
go run 5/main.go
*/
