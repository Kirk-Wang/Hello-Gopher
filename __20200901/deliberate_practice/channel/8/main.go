package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

// 自己建 channel
// 返回的 chanel 是干嘛用的呢，没错-->>是用来发数据的，送数据的
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	// 开 10 个 worker
	// 每个人都有一个 channel
	// 然后分别向它们分发
	var channels [10]chan<- int
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

	time.Sleep(time.Millisecond)
}

func bufferedChannel() {

	c := make(chan int, 3) // 加个缓冲区，对提升性能是有帮助的
	go worker(0, c)

	c <- 1
	c <- 2
	c <- 3
	c <- 4

	time.Sleep(time.Millisecond)
}

func channelClose() {
	// close 一定是发送方 close

	c := make(chan int, 3) // 加个缓冲区，对提升性能是有帮助的
	go worker(0, c)

	c <- 1
	c <- 2
	c <- 3
	c <- 4

	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()
	fmt.Println("Buffered channel")
	bufferedChannel()
	fmt.Println("Channel close and range")
	channelClose()
}

/*
go run 8/main.go
*/
