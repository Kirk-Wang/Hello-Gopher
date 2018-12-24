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

### 数据存储->Elasticsearch

![数据存储](./images/Go15.png)

### 分布式初探

![分布式初探](./images/Go16.png)

### 总结

![总结](./images/Go17.png)

---

## 基于 Go 语言构建企业级的 RESTful API 服务


### 高可用 API 构建

![高可用 API 构建](./images/RESTful/api-1.png)

### RESTful API

![RESTful API](./images/RESTful/api-2.png)

### API流程和代码结构

![API流程和代码结构](./images/RESTful/api-3.png)

### 启动一个最简单的 RESTful API 服务器

![启动一个最简单的RESTful API 服务器](./images/RESTful/api-4.png)

### 配置文件读取

![配置文件读取](./images/RESTful/api-5.png)

### 记录和管理 API 日志

![记录和管理 API 日志](./images/RESTful/api-6.png)



