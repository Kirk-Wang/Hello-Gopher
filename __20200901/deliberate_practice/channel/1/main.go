package main

import "fmt"

func chanDemo() {
	c := make(chan int)
	// 发了数据没人收，就会 deadlock
	c <- 1
	c <- 2

	n := <-c
	fmt.Println(n)
}

func main() {
	chanDemo()
}

/*
go run 1/main.go

goroutine 1 [chan send]:
main.chanDemo()
        /Users/wangzuowei/github/Hello-Gopher/__20200901/deliberate_practice/channel/1/main.go:7 +0x59
main.main()
        /Users/wangzuowei/github/Hello-Gopher/__20200901/deliberate_practice/channel/1/main.go:15 +0x20
exit status 2
*/
