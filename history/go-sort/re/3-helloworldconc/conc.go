package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	// go starts a goroutine
	for i := 0; i < 5000; i++ {
		// go starts a goroutine
		go printHelloWorld(i, ch)
	}
	for {
		msg := <-ch
		fmt.Println(msg)
	}
	// time.Sleep(time.Millisecond)
}

func printHelloWorld(i int, ch chan string) {
	for {
		ch <- fmt.Sprintf("Hello world from goroutine %d!\n", i)
	}
}
