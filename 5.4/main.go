package main

import (
	"fmt"
	"github.com/Kirk-Wang/Hello-Go/5.2/mock"
	"github.com/Kirk-Wang/Hello-Go/5.2/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

// 使用者定义接口
// 一个文件

func download(r Retriever) string {
	return r.Get("http://www.lotteryjs.com")
}

func post(poster Poster) string {
	return poster.Post("http://www.lotteryjs.com", map[string]string{
		"name":   "mouse",
		"course": "golang",
	})
}

func main() {
	var r Retriever
	r = &mock.Retriever{Contents: "this is a fake lotteryjs.com"}
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)
	// fmt.Println(r.(*real.Retriever).UserAgent)

	// Type assertion
	// realRetriever := r.(*real.Retriever)
	// fmt.Println(realRetriever.TimeOut)

	if mockRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not mock retriever")
	}

	// fmt.Println(download(r))
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	fmt.Println("Type switch:")
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
