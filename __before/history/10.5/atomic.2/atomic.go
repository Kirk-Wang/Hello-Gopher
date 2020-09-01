package main

import (
	"fmt"
	"sync"
	"time"
)

// Mutex（互斥量） 的使用
type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increment() {
	// 锁来保护
	// a.lock.Lock()
	// a.value++
	// a.lock.Unlock()
	a.lock.Lock()
	defer a.lock.Unlock()

	a.value++
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()

	return a.value
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
	// 读的时候，有可能正在写，所以要加锁
	fmt.Println(a.get())
}
