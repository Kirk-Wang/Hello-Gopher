package main

import (
	"fmt"
)

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int // m3 == nil Go语言的 nil 是可以参与运算的

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	// for k := range m {
	// 	fmt.Println(k)
	// }

	// for _, v := range m {
	// 	fmt.Println(v)
	// }

	fmt.Println("Getting values")
	// courseName := m["course"]
	// fmt.Println(courseName)
	// causeName := m["cause"]
	// fmt.Println(causeName) // 实质是一个Zero Value

	courseName, ok := m["course"]
	fmt.Println(courseName, ok)

	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
