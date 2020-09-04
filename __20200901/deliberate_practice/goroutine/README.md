# Goroutine

### 协程 Coroutine

* 轻量级“线程”
* `非抢占式`多任务处理，由协程主动交出控制权
* 编译器/解释器/虚拟机层面的多任务
* 多个协程可能在一个或多个线程上运行

1.4.2 Corotines

Subroutines are special cases of more general program
components，called coroutines. In contrast to the unsymmetric
* 子程序是协程的一个特例
* Donnald Knuth "The Art of Computer Programming. Vol1"

### goroutine 的定义

* 任何函数只需加上 go 就能送给调度器运行
* 不需要在定义时区分是否是异步函数
* 调度器在合适的点进行切换
* 使用 -race 来检测数据访问的冲突

### goroutine 可能的切换点

* I/O, select
* channel
* 等待锁
* 只是参考，不能保证切换，不能保证在其它地方不切换

### goroutine 开 1000 个，我们系统到低起了多少个线程

```sh
top
```

