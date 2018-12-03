package main

import (
	"fmt"
)

func lengthOfNonRepeatingSubStr(s string) int {
	// 定义一个map，记录每个字母最后出现的位置
	lastOccurd := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {
		// lastOccurd[ch] 有可能不存在，不存在的话会出来一个零
		lastI, ok := lastOccurd[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurd[ch] = i
	}

	return maxLength
}

func main() {
	fmt.Println(
		lengthOfNonRepeatingSubStr("ababcab"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("bbbb"))
	fmt.Println(
		lengthOfNonRepeatingSubStr(""))
	fmt.Println(
		lengthOfNonRepeatingSubStr("cdcdcxxxdcd"))
}
