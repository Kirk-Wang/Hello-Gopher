package main

import (
	"bufio"
	"fmt"
	"github.com/Kirk-Wang/Hello-Go/7.1/fib"
	"os"
)

func tryDefer() {
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// fmt.Println(3)
	// // return
	// panic("error occured")
	// fmt.Println(4)

	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	// 建文件
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	// 建完，就要想着关掉这个io
	defer file.Close() // 想到啥就是啥，先进后出都是非常合理的

	// 这个写的比较快，先到内存，大到一定程度，再往硬盘里写
	// 建了buffer io 就要想到，我们之后要 flush
	writer := bufio.NewWriter(file)
	defer writer.Flush() // 导入文件

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	tryDefer()
	// writeFile("fib.txt")
}
