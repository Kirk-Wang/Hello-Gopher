package main

import (
	"fmt"
)

var lastOccured = make([]int, 0xffff) // 65535

func lengthOfNonRepeatingSubStr(s string) int {

	// 定义一个map，记录每个字母最后出现的位置
	// lastOccured := make(map[byte]int)
	// 既然有一半的时间花在了这个 map 上，那我们优化一下
	// lastOccured := make(map[rune]int) // 每个rune塞进去->算hash->判重->分配空间->so,slow
	// 所以能不能一口气把空间分配好，老话是：空间换时间
	// 这步改动，针对字符型的map,不用map,而是开一个较大的Slice
	/*
		假设每个中文字符最大就是Oxffff(这里开了16位)，但是rune是32位的，所以只能是大概
		解析：Yes我爱敲代码!
			lastOccured['e'] = 1 -> lastOccured[0x65] = 1
			lastOccured['代'] = 6 -> lastOccured[0x8BFE] = 6
	*/
	// 这个移出去，避免500次每次都 makeslice
	// lastOccured := make([]int, 0xffff) // 65535--> 65k
	for i := range lastOccured {
		lastOccured[i] = -1 // 初始化值
	}
	start := 0
	maxLength := 0

	// for i, ch := range []byte(s) {
	for i, ch := range []rune(s) {
		// lastOccured[ch] 有可能不存在，不存在的话会出来一个零
		lastI := lastOccured[ch]
		if lastI != -1 && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccured[ch] = i
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
	fmt.Println(
		lengthOfNonRepeatingSubStr("我是中国人"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("一二三三二一"))
}
