package main

import (
	"fmt"

	pipleline "github.com/Kirk-Wang/Hello-Gopher/history/go-sort/re/5-pipleline"
)

func main() {
	p := pipleline.ArraySource(3, 2, 6, 7, 4)
	// for {
	// 	if num, ok := <-p; ok {
	// 		fmt.Println(num)
	// 	} else {
	// 		// channel 关闭了
	// 		break
	// 	}
	// }
	for v := range p {
		// 发送方一定要 close
		fmt.Println(v)
	}
}
