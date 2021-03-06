package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}

	printCityList(all)
}

// 确认页面编码
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	// 拿到前 1024 个 bytes
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	// matches := re.FindAll(contents, -1) // [][]byte -> []byte 相当于 string
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		// for _, subMatch := range m {
		// 	// fmt.Printf("%s ", subMatch)
		// }
		fmt.Printf("City: %s, URL: %s", m[2], m[1])
		fmt.Println()
	}
	fmt.Printf("Matches found: %d\n", len(matches))
}
