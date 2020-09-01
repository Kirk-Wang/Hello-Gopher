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
		time.Sleep(time.Second) // sleep 个 5 秒
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

// 真实产品中，输入数据和输出数据很可能是不一样的速度
func main() {
	var c1, c2 = generator(), generator()
	worker := createWorker(0)
	var values []int // 用来排队
	// 想要 10 秒钟之后它退出
	tm := time.After(10 * time.Second)
	// 每秒钟，我想看看队列中积压了多少数据
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		// 队列里面有值，就去消耗
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:] // 出队
		// 每次 select 花的时间，800 毫秒内没有生出数据，就timeout
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick: // 这是 timeout 就很难达到了
			fmt.Println("queue len =", len(values))
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}
