package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for {
		fmt.Printf("Worker %d received %d\n", id, <-c)
	}
}

func chanDemo() {
	c := make(chan int)

	go worker(0, c)
	// 发了数据没人收，就会 deadlock
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
}

/*
go run 3/main.go
*/
