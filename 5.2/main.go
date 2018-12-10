package main

import (
	"fmt"
	"github.com/Kirk-Wang/Hello-Gopher/5.2/mock"
	"github.com/Kirk-Wang/Hello-Gopher/5.2/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.lotteryjs.com")
}

func main() {
	var r Retriever
	r = &mock.Retriever{"this is a fake lotteryjs.com"}
	r = &real.Retriever{}
	fmt.Println(download(r))
}
