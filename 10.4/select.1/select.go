package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(
				time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(5 * time.Second) // sleep 个 5 秒
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	// var c1, c2 chan int // c1 and c2 = nil
	var c1, c2 = generator(), generator()
	worker := createWorker(0)
	// n := <- c1
	// n := <- c2 同时收，谁来的快要谁

	n := 0
	hasValue := false
	for {
		// nil channel 在 select 里面可以正确运行
		// 但肯定不会被 select 到，就是说它永远是阻塞的
		var activeWorker chan<- int
		if hasValue {
			activeWorker = worker
		}
		select {
		// 好像是做了一个非阻塞式的获取(select + default)
		// channel 里面不管是发数据还收数据都是阻塞的
		case n = <-c1: // 如果这里连续的收? 要排队，不然后面的数据会冲掉前面的
			hasValue = true
			// fmt.Println("Received from c1:", n)
			// w <- n // 当然是可以，但又会被阻塞掉
		case n = <-c2:
			hasValue = true
			// fmt.Println("Received from c2:", n)
			// default: 注释掉了 --> deadlock 我在收数据，但没有人发数据
			// 	fmt.Println("No value received")
			// w <- n // 当然是可以，但又会被阻塞掉
		case activeWorker <- n:
			// 生成数据的 channel 和消耗数据的 channel ，速度是不一样的~
			hasValue = false
		}
	}
}
