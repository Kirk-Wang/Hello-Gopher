package main

import (
	"fmt"
	"github.com/Kirk-Wang/Hello-Gopher/go-sort/pipeline"
)

func main() {
	// p := pipeline.ArraySource(3, 2, 6, 7, 4)
	// for {
	// 	if num, ok := <-p; ok { // 管道还在就打印
	// 		fmt.Println(num)
	// 	} else {
	// 		break
	// 	}
	// }

	// p := pipeline.InMemSort(
	// 	pipeline.ArraySource(3, 2, 6, 7, 4),
	// )

	p := pipeline.Merge(
		pipeline.InMemSort(pipeline.ArraySource(3, 2, 6, 7, 4)),
		pipeline.InMemSort(pipeline.ArraySource(7, 4, 0, 3, 2, 13, 8)),
	)

	// 使用 range ，发送方一定要close
	for v := range p {
		fmt.Println(v)
	}
}
