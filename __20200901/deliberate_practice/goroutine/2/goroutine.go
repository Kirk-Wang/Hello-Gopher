package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Running in", runtime.Version())
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from goroutine %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Second * 5)
}
