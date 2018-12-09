package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(ii int) { // go func() -> race condition!数据访问冲突的意思
			for {
				// 好像这段程序和抢占式没有啥区别呀，每个人都在打，打到一半都会跳出来，都会换成别人
				// Printf是个IO的操作，里面会存在协程的切换，IO操作里面总会有等待的过程
				// fmt.Printf("Hello from goroutine %d\n", i) // IO 操作会交出控制权

				// 这里只是一句指令，没有机会进行协程之间的一个切换，这样就会被一个协程给抢掉
				// 如果不主动交出控制权，它就始终在这个协程里面
				// 运行之后，发现死机了，看下 top 发现 cpu 快 400% 了(本机是4核)
				// 在死循环中一直退不出来，原因是goroutine里面a[i]++交不出控制权，交不出就一直死在里面
				// 同样 main 函数自己也是一个goroutine，虽然其他goroutine是它开出来的
				a[ii]++
				// 手动交出控制权，让别人也有机会运行，这样大家才有机会一起并发的运行，但这句很少用，有其它方式
				runtime.Gosched()
			}
		}(i)
	}
	//因为没人交出控制权，所以永远Sleep不出来，虽然只sleep了一秒。
	time.Sleep(time.Millisecond)

	// IO 操作，会出现一边我们在 Print `a`，一边上面还在并发的写，这个要用后面的 channel 解决
	fmt.Println(a)
}
