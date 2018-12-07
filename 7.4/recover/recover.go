package main

import (
	// "errors"
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()                // 它是一个 interface{}
		if err, ok := r.(error); ok { // 一定要是一个 error 类型
			fmt.Println("Error occurred:", err)
		} else {
			// panic(r)
			panic(fmt.Sprintf(
				"I don't know what to do: %v", r))
		}
	}()
	// panic(errors.New("this is an error"))
	// b := 0
	// a := 5 / b
	// fmt.Println(a)
	panic(123)
}

func main() {
	tryRecover()
}
