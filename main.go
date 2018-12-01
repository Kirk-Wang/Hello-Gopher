package main

import (
	"fmt"
	"strings"
	// _ "github.com/Kirk-Wang/Hello-Go/bar"
	// _ "github.com/Kirk-Wang/Hello-Go/say"
)

// var foo string
// var bar int

// var (
// 	foo string = "Hello"
// 	bar int    = 100
// )

// const (
// 	Monday    = 1
// 	Tuesday   = 2
// 	Wednesday = 3
// )

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

func add(i, j int) int {
	return i + j
}

func swap(i, j int) (int, int) {
	return j, i
}

func foo() func() int {
	return func() int {
		return 100
	}
}

func getUserListSQL(username, email string, sexy int) string {
	sql := "select * from user"
	where := []string{}

	if username != "" {
		where = append(where, fmt.Sprintf("username = '%s'", username))
	}

	if email != "" {
		where = append(where, fmt.Sprintf("email = '%s'", email))
	}

	if sexy != 0 {
		where = append(where, fmt.Sprintf("sexy = '%d'", sexy))
	}

	return sql + " where " + strings.Join(where, " or ")
}

type searchOpts struct {
	username string
	email    string
	sexy     int
}

func getUserListOptsSQL(opts searchOpts) string {
	sql := "select * from user"
	where := []string{}

	if opts.username != "" {
		where = append(where, fmt.Sprintf("username = '%s'", opts.username))
	}

	if opts.email != "" {
		where = append(where, fmt.Sprintf("email = '%s'", opts.email))
	}

	if opts.sexy != 0 {
		where = append(where, fmt.Sprintf("sexy = '%d'", opts.sexy))
	}

	return sql + " where " + strings.Join(where, " or ")
}

func main() {
	// foo = "Hello"
	// bar = 100
	// foo := "Hello"
	// bar := 100
	// fmt.Println(foo)
	// fmt.Println(bar)
	// fmt.Println(Monday)
	// fmt.Println(Tuesday)
	// fmt.Println(add(1, 2))
	// a := 1
	// b := 2
	// a, b = swap(a, b)
	// fmt.Println("a:", a)
	// fmt.Println("b:", b)
	// a, b = b, a
	// fmt.Println("a:", a)
	// fmt.Println("b:", b)

	// bar := foo()
	// fmt.Printf("%T\n", bar)
	// fmt.Println(bar())

	// bar2 := func(i, j float32) float32 {
	// 	return i + j
	// }
	// fmt.Printf("%T\n", bar2)
	// fmt.Println(bar2(1.45, 2.3))

	// foo := func() string {
	// 	return "Hello World"
	// }
	// fmt.Println(foo())

	// bar := func() {
	// 	fmt.Println("Hello World 2")
	// }

	// bar()

	// func() {
	// 	fmt.Println("Hello World 3")
	// }()

	// go func(i, j int) {
	// 	fmt.Println(i + j)
	// }(1, 2)

	fmt.Println(getUserListSQL("Kirk", "", 0))
	fmt.Println(getUserListSQL("Kirk", "Kirk@gmail.com", 1))

	fmt.Println(getUserListOptsSQL(searchOpts{
		username: "Kirk",
	}))

	fmt.Println(getUserListOptsSQL(searchOpts{
		username: "Kirk",
		email:    "Kirk@gmail.com",
	}))

	fmt.Println(getUserListOptsSQL(searchOpts{
		username: "Kirk",
		email:    "Kirk@gmail.com",
		sexy:     1,
	}))
}
