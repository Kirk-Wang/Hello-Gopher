package main

import (
	"fmt"
	"regexp"
)

const text = `
	a is haha@sina.com
	bb is kkA1@qq.com.cn
	oo is	o-o@gmail.com
	xxx-x is lot-o_s@163.com  
`

func main() {
	re := regexp.MustCompile(`[a-zA-Z0-9\-_]+@([a-zA-Z0-9\-_\.]+)`)
	// match := re.FindAllString(text, -1)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
}
