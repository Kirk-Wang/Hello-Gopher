package pipleline

import (
	"io"
	"sort"
)

func ArraySource(a ...int) <-chan int {
	// <-chan int: 表示用它的人只能从里面拿东西
	// 在这里我们只能放东西
	out := make(chan int)
	go func() {
		for _, v := range a {
			out <- v
		}
		// 送完数据，直接关掉通道
		close(out)
	}()
	return out
}

// in 是节点的一个输入
// return 的 channel 是一个输出
// return 是相对于使用者的来说
// 使用的人只能用它去收东西
// 对于 InMemSort 只能发东西
// in <-chan int::只进不出的 channel
// 返回的是  <-chan int：只出不进的 channel
func InMemSort(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		// Read into Memory
		a := []int{}
		for v := range in {
			a = append(a, v)
		}
		// Sort
		sort.Ints(a)

		// Output
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		// 同时从两个 channel 获得数据，然后比较他们的大小
		// 而且我们获得的数据不一定有
		// 这两个 channel 的数据量不一定一样
		// 有的有数据，有的可能已经没了
		// 因此要处理没数据的情况
		v1, ok1 := <-in1 // 会等排序完成，才会送第一个数来
		v2, ok2 := <-in2 // 会等排序完成，才会送第一个数来
		for ok1 || ok2 {
			// 考虑啥时候会输出 v1 ?
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
	}()
	return out
}

func ReaderSource(reader io.Reader) <-chan int {
	// go 语言的 int 有多大呢？
	// 它是根据系统来的
	out := make(chan int)
	return out
}
