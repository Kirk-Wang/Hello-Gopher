package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	m := len(s) / 2
	left := s[:m]
	right := s[m:]
	fmt.Printf("%d\n", m/2)
	fmt.Printf("%v\n", left)
	fmt.Printf("%v\n", right)
}
