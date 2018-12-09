package main

import (
	"fmt"
	"time"
)

func main() {
	// goroutine开 1000 个，看看我们系统启用了多少个线程
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from goroutine %d\n", i)
			}
		}(i)
	}
	// 长一点，用top看一下
	// 发现总线程是8个，活动的是4个，我这个机器是个4核的机器
	// 调度器会根据我们的核来，没必要让系统来
	time.Sleep(time.Minute)
}
