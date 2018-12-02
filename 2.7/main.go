package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b) // _ 表示第二个我不想要
		return q, nil
	default:
		// panic("unsupported operation: " + op) // 直接中断，很难看
		return 0, fmt.Errorf(
			"unsupported operation: %s", op)
	}
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

func div2(a, b int) (q, r int) {
	// q = a / b
	// r = a % b
	return a / b, a % b // 建议
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling functions %s with args "+"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

// func swap(a, b *int) {
// 	*b, *a = *a, *b
// }

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	// fmt.Println("main")
	// fmt.Println(eval(3, 4, "/"))
	// fmt.Println(div(13, 3))
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	q, r := div2(14, 3)
	fmt.Println(q, r)

	fmt.Println(apply(pow, 3, 4))

	fmt.Println(apply(
		func(a int, b int) int {
			return int(math.Pow(
				float64(a), float64(b)))
		}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))

	a, b := 3, 4
	// swap(&a, &b)
	a, b = swap(a, b)
	fmt.Println(a, b)
}
