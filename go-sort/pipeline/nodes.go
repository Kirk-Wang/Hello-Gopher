package pipeline

import (
	"sort"
)

// 可变长参数
// 对于用的人来说，只能拿东西
func ArraySource(a ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range a {
			out <- v
		}
		close(out) // 数据送完了，结束它
	}()
	return out
}

// in 是节点的一个输入，方向是从 channel 进到我们里面来（只进不出）
// return的 <-chan int 是一个输出，相对于使用它的人来说，只能从它收东西（只出不进)
// InMemSort只能发东西
func InMemSort(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		// Read into memory
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
		// 1.同时从两个 chanel 去获得数据
		// 2.获得的数据不一定有
		// 3.两个 chanel 的数据量不一定一样
		// 4.有的人还有数据，有的人已经没有数据了
		// 所以要处理没数据的情况
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		// 只要一个有数据，就送出数据
		for ok1 || ok2 {
			// 每次只送出一个数据，要么是v1，要么是v2，
			// 这里的条件比较复杂

			// 1. 送chanel1 的数据
			// chanel 2 不能有数据 或者
			// chanel 1必须有数据，而且数据要小于等于chanel 2的数据
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1       // 送出数据
				v1, ok1 = <-in1 // 更新 v1 的结果
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
	}()
	return out
}
