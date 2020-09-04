package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	fmt.Println("Running in", runtime.Version())
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				// IO 的操作，会有协程之间的切换
				// fmt.Printf("Hello from goroutine %d\n", i)
				// race condition
				// go run -race goroutine.go
				a[i]++
				// runtime.Gosched() --> 让其它 goroutine 有机会调度。当然，go 1.14 已经把这个坑已经填了~~
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

/*
Running in go1.14.2
==================
WARNING: DATA RACE
Read at 0x00c0001ae050 by main goroutine:
  main.main()
      /Users/wangzuowei/github/Hello-Gopher/__20200901/deliberate_practice/goroutine/goroutine.go:25 +0x193

Previous write at 0x00c0001ae050 by goroutine 7:
  main.main.func1()
      /Users/wangzuowei/github/Hello-Gopher/__20200901/deliberate_practice/goroutine/goroutine.go:19 +0x68

Goroutine 7 (running) created at:
  main.main()
      /Users/wangzuowei/github/Hello-Gopher/__20200901/deliberate_practice/goroutine/goroutine.go:13 +0x15e
==================
[3408 259 3079 667 604 647 677 667 671 640]
Found 1 data race(s)
exit status 66

看到 25 行我们在读~
看到 19 行我们在写
说明，我们一边在 Println(a)，一边在并发在里面写 --> 它被 go race 检测出来了~这个可以用channel来解决~
*/
