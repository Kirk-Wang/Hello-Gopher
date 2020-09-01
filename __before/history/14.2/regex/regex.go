package main

import (
	"fmt"
	"regexp"
)

const text = `
my email is ccmouse@gmail.com@abc.com
email1 is abc@def.org
email2 is      kkk@qq.com
email3 is ddd@abc.com.cn
`

func main() {
	// re, err := regexp.Compile("ccmouse@gmail.com")
	// 其实不用返回 err, 我们写程序的认为我们写的就是对的
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	// match := re.FindString(text)// 指参数是一个 string
	// match := re.FindAllString(text, -1) // -1 全局匹配
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
}
