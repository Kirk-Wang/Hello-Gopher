package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// 求mo &  / 2
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	printFileContents(file)
	// scanner := bufio.NewScanner(file)

	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("kirk")
	}
}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(72387885),
		convertToBin(0),
	)

	printFile("kirk.txt")

	s := `
	dasdf
	asdfdas
	asdfa
	123
	"hu"
	"lu"
	
	`

	printFileContents(strings.NewReader(s))
}
