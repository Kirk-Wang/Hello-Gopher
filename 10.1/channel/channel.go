package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	// for {
	// 	// n := <-c
	// 	// fmt.Println(n)
	// 	// fmt.Printf("Worker %d received %d\n", id, <-c)
	// 	// 为啥会啥乱序？因为 IO 操作，goroutine 会进行调度
	// 	// fmt.Printf("Worker %d received %c\n", id, <-c)
	// 	n, ok := <-c
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Printf("Worker %d received %d\n", id, n) // channel close 还是会收到数据的(具体类型的0,int->0 ,string->"")
	// }

	for n := range c { // 等到 c close，它就跳出来
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

// 返回一个 channel
// chan<- <=====> 告诉外面这个 channel 是用来发送数据的
// <-chan(告诉外面只能收)
// chan<-(告诉外面只能发)
func createWorker(id int) chan<- int {
	c := make(chan int)
	// go func() {
	// 	for {
	// 		// 里面就只能收
	// 		fmt.Printf("Worker %d received %c\n", id, <-c)
	// 	}
	// }()
	go worker(id, c)
	return c
}

func chanDemo() {
	// 表示 c 是一个 channel, 里面内容是一个 int
	// 这里只是定义了，所以此时 c == nil，目前是没有办法用的，所以暂时不用
	// var c chan int

	// 创建一个 channel
	// c := make(chan int)
	// fatal error: all goroutines are asleep - deadlock!
	// channel 是 goroutine 与 goroutine 之间的交互
	// 得要人收
	// go func() {
	// 	for {
	// 		n := <-c // 不断的从 channel 里面收数据
	// 		fmt.Println(n)
	// 	}
	// }() 放出去，不用闭包
	// go worker(0, c)              // 传个 channel 过去

	// c <- 1                       // 发一个 1 给 channel c
	// c <- 2                       // 发一个 2 给 channel c

	// 开多个worker，为每个 worker 建立单独的 channel
	// var channels [10]chan int
	// for i := 0; i < 10; i++ {
	// 	channels[i] = make(chan int)
	// 	go worker(i, channels[i])
	// }

	var channels [10]chan<- int // 专用来送数据的，不能收
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	// 分别向每个 channel 分发数据
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
		// n := <-channels[i] invalid operation: <-channels[i] (receive from send-only type chan<- int)
	}
	// 再来一次
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond) // 避免来不及打印
}

func bufferedChannel() {
	// c := make(chan int)
	// c <- 1 // 这样会 deadlock, 发了之后必须要有人收

	// 我们可以弄个缓冲区，提升性能有一定的优势
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	time.Sleep(time.Millisecond) // 避免来不及打印
}

func channelClose() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
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
