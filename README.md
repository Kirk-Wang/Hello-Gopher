# Gopher

![Gopher](https://golang.org/doc/gopher/frontpage.png)

Go Programming Language 扫盲

go1.11.2

### Go 基础语法脑图回顾

![Go 基础语法脑图回顾](./images/Go1.png)

### Go 内建容器脑图回顾

![Go 内建容器脑图回顾](./images/Go2.png)

### Go 面向“对象”

![Go 面向“对象”](./images/Go3.png)

### Go 面向接口

![Go 面向接口](./images/Go4.png)

### Go 函数式编程

![Go 函数式编程](./images/Go5.png)

### Go 错误处理和资源管理

![Go 错误处理和资源管理](./images/Go6.png)

### Go 测试与性能调优

![Go 测试与性能调优](./images/Go7.png)

### Goroutine

![Goroutine](./images/Go8.png)

### Channel

![Channel](./images/Go9.png)

### http及其他标准库

![http及其他标准库](./images/Go10.png)

### 热身：迷宫的广度优先搜索
              6 5

      (start) 0 1 0 0 0

              0 0 0 1 0

              0 1 0 1 0

              1 1 1 0 0

              0 1 0 0 1

              0 1 0 0 0(end)

![迷宫的广度优先搜索](./images/Go11.png)

### 实战项目

如何只利用Go语言，来实现一个分布式应用程序(麻雀虽小，五脏俱全^_^)？

![实战项目](./images/Go12.png)

### 单任务版

![单任务版](./images/Go13.png)

单任务版网络利用率（每秒才80多K,因为是一个一个fetch）

![1.0](./images/1.0spider.png)

### 并发版

![单任务版](./images/Go14.png)

并发版网络利用率（瞬间飙到1M多每秒，你网速稍微快一点的话3，4M每秒是很正常的）

![1.0](./images/2.0spider.png)

我这里开了100个 goroutine 去做任务，看下面 Go 给我们开了 13 线程跑这 100 个协程，但活跃的只有3，4个

因为我的机器就只有4核（🤣），当然还可以开大一点来榨干 CPU 🤣

![1.0](./images/2.0top.png)

