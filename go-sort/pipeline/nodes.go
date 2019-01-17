package pipeline

// 可变长参数
func ArraySource(a ...int) chan int {
	out := make(chan int)
	go func() {
		for _, v := range a {
			out <- v
		}
		close(out) // 数据送完了，结束它
	}()
	return out
}
