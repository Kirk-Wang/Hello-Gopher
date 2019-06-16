package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("https://www.zhenai.com/zhenghun")

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	printCityList(all)
}

func printCityList(contents []byte) {
	// {linkContent:"阿坝",linkURL:"http://m.zhenai.com/zhenghun/aba"}
	re := regexp.MustCompile(`{linkContent:"([^"征婚]+)",linkURL:"([^"]+)"}`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("City: %s, URL: %s\n", m[1], m[2])
		fmt.Println()
	}
	fmt.Printf("Matches found: %d\n", len(matches))
}
