package main

import (
	"bufio"
	"fmt"
	"os"

	pipeline "github.com/Kirk-Wang/Hello-Gopher/history/go-sort/re/5-pipeline"
)

func main() {
	p := createPipeline("small.in", 512, 4)
	writeToFile(p, "small.out")
	printFile("small.out")
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.ReaderSource(file, -1)
	for v := range p {
		fmt.Println(v)
	}
}

func writeToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriterSink(writer, p)
}

func createPipeline(
	filename string,
	fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount

	sortResults := []<-chan int{}
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		// 从第几块开始读，0 表示从头开始
		file.Seek(int64(i*chunkSize), 0)

		source := pipeline.ReaderSource(
			bufio.NewReader(file),
			chunkSize,
		)
		// 分成了 4 块，每一块单独的进行 InMemSort
		sortResults = append(sortResults, pipeline.InMemSort(source))
	}
	return pipeline.MergeN(sortResults...)
}
