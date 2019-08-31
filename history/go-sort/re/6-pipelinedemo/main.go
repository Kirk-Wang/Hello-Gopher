package main

import (
	"bufio"
	"fmt"
	"os"

	pipeline "github.com/Kirk-Wang/Hello-Gopher/history/go-sort/re/5-pipeline"
)

func main() {
	const filename = "large.in"
	// 我的机器：1 个 int = 64 bit = 8 byte
	// 100000000 * 8 = 800000000 byte = 800000 KB = 800 MB 数据文件
	const n = 100000000
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.RandomSource(n)
	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer, p)
	writer.Flush() // 最后全部倒出去

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p = pipeline.ReaderSource(bufio.NewReader(file))
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
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
