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
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()

	return out
}

func worker(id int, c chan int) {
	for {
		fmt.Printf("Worker %d received %d\n", id, <-c)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 chan int = generator(), generator()

	w := createWorker(0)

	for {
		select {
		case n := <-c1:
			w <- n // 问题是，这里也会 block
		case n := <-c2:
			w <- n
		}
	}
}
