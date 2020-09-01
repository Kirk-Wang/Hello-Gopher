package main

import (
	"fmt"
	"github.com/Kirk-Wang/Hello-Gopher/5.5/mock"
	"github.com/Kirk-Wang/Hello-Gopher/5.5/real"
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

const url = "https://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) string {
	return poster.Post(url, map[string]string{
		"name":   "mouse",
		"course": "golang",
	})
}

// 组合接口
type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked lotteryjs.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{Contents: "this is a fake lotteryjs.com"}
	r = &retriever
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

	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not mock retriever")
	}

	// fmt.Println(download(r))

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Println(" > Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
