package main

import (
	"fmt"
	"time"
)

func chanDemo() {
	c := make(chan int)

	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()
	// 发了数据没人收，就会 deadlock
	c <- 1
	c <- 2

	time.Sleep(time.Microsecond)
}

func main() {
	chanDemo()
}

/*
go run 2/main.go
*/
