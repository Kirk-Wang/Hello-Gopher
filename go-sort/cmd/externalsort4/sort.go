package main

import (
	"bufio"
	"fmt"
	"github.com/Kirk-Wang/Hello-Gopher/go-sort/pipeline"
	"os"
)

func main() {
	p := createPipeline("large.in", 800000000, 8)
	writeToFile(p, "large.out")
	printFile("large.out")
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.ReaderSource(file, -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
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

	pipeline.WirterSink(writer, p)
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
		sortResults = append(sortResults, pipeline.InMemSort(source))
	}
	return pipeline.MergeN(sortResults...)
}
