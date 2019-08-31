package main

import (
	"fmt"
	"os"

	pipeline "github.com/Kirk-Wang/Hello-Gopher/history/go-sort/re/5-pipeline"
)

func main() {
	file, err := os.Create("small.in")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.RandomSource(50)
	pipeline.WriterSink(file, p)
}

func mergeDemo() {
	p := pipeline.Merge(
		pipeline.InMemSort(
			pipeline.ArraySource(3, 2, 6, 7, 4)),
		pipeline.InMemSort(
			pipeline.ArraySource(7, 4, 0, 3, 2, 13, 8)))
	// for {
	// 	if num, ok := <-p; ok {
	// 		fmt.Println(num)
	// 	} else {
	// 		// channel 关闭了
	// 		break
	// 	}
	// }
	for v := range p {
		// 注意：
		// fatal error: all goroutines are asleep - deadlock!
		// goroutine 1 [chan receive]:
		// 记得发送方要 close 掉
		fmt.Println(v)
	}
}
