package main

import (
	"fmt"
	"time"
)

// Mutex（互斥量） 的使用
type atomicInt int

func (a *atomicInt) increment() {
	*a++
}

func (a *atomicInt) get() int {
	return int(*a)
}

// atomeic 指的是线程安全，当然在 go 语言中没有线程的事
func main() {
	// "sync/atomic"
	// atomic.AddInt32 多个 goroutine 并发的执行的话，它也是安全的，具体的应用中用系统的
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	// go run -race ./10.5/atomic/atomic.go
	// 读的时候，有可能正在写
	fmt.Println(a)
}
