package main

import (
	"fmt"
)

func sum(numbers ...int) int {
	s := 0
	for _, val := range numbers {
		s += val
	}

	return s
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
}
