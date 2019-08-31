package pipleline

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

}
