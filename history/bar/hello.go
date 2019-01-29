package bar

import (
	"fmt"
)

func init() {
	fmt.Println("bar init")
}

func Hello() string {
	return "Hello bar"
}
