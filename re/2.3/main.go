package main

import (
	"fmt"
)

func main() {
	const (
		a = 1 << (10 * iota)
		b
		c
	)

	fmt.Println(a, b, c)
}
