package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/Kirk-Wang/Hello-Gopher/7.2/fib"
	"os"
)

func tryDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	// 先产生一个错误
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)

	cc := errors.New("custom Error")
	fmt.Println(fmt.Errorf("%s", cc))
	if err != nil {
		// fmt.Println("Error：", err.Error())
		// err 肚子里有啥呢？
		if pathError, ok := err.(*os.PathError); ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
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
	// tryDefer()
	writeFile("fib.txt")
}
