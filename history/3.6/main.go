package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes学内建容器!"
	fmt.Println(len(s)) // 19 每个中文三字节 UTF-8采用可变长，英文一字节，中文三字节
	fmt.Println(s)

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	} // 59 65 73 E5 AD A6 E5 86 85 E5 BB BA E5 AE B9 E5 99 A8 21
	fmt.Println()

	for i, ch := range s { // ch is a rune，rune 和 int32 就一个4字节整数的别名
		fmt.Printf("(%d %X) ", i, ch)
	}
	/*
		(0 59) (1 65) (2 73) (3 5B66) (6 5185) (9 5EFA) (12 5BB9) (15 5668) (18 21)
		(3 5B66) “学” 他是一个4️字节的，在Unicode 编码中，这里只有后面两个字节是有字的
		E5 AD A6 : 这个是 utf-8 的编码，它转成了一个Unicode 的编码，叫 5B66

		这里是它把's'进行utf-8的解码，转出来的每一个字符又把它转成Unicode,然后又把它放在rune这个4字节类型里面，最后给了我们
	*/

	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s)) // 9

	bytes := []byte(s)

	for len(bytes) > 0 {
		// size: 英文字符是1，中文字符是 3
		ch, size := utf8.DecodeRune(bytes)
		// 拿完之后，继续切片
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	// 其实直接转 rune 切片就好，比较上层
	for i, ch := range []rune(s) { // 这里另外开了一块内存，转码(utf8-unicode-rune)后存下来
		fmt.Printf("(%d %c)", i, ch)
	}
	fmt.Println()
}
