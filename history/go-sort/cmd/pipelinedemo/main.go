package main

import (
	"bufio"
	"fmt"
	"github.com/Kirk-Wang/Hello-Gopher/go-sort/pipeline"
	"os"
)

func arraySource() {
	p := pipeline.ArraySource(3, 2, 6, 7, 4)
	for {
		if num, ok := <-p; ok { // 管道还在就打印
			fmt.Println(num)
		} else {
			break
		}
	}
}

func inMemSort() {
	p := pipeline.InMemSort(
		pipeline.ArraySource(3, 2, 6, 7, 4),
	)
	for v := range p {
		fmt.Println(v)
	}
}

func merge() {
	p := pipeline.Merge(
		pipeline.InMemSort(pipeline.ArraySource(3, 2, 6, 7, 4)),
		pipeline.InMemSort(pipeline.ArraySource(7, 4, 0, 3, 2, 13, 8)),
	)
	// 使用 range ，发送方一定要close
	for v := range p {
		fmt.Println(v)
	}
}

func main() {
	// const filename = "large.in"
	// const n = 100000000 // 100M * 8
	const filename = "small.in"
	const n = 64
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.RandomSource(n)
	// pipeline.WirterSink(file, p)
	writer := bufio.NewWriter(file)
	pipeline.WirterSink(writer, p)
	writer.Flush()

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// p = pipeline.ReaderSource(file)
	p = pipeline.ReaderSource(bufio.NewReader(file), -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
	// ls -lrt
}
