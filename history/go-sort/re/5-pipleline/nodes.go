package pipleline

func ArraySource(a ...int) chan int {
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
