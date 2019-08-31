## 搭建并行处理管道，感受 Go 语言的魅力

慕课网视频：[搭建并行处理管道，感受 Go 语言的魅力](https://www.imooc.com/learn/927)
* CCMOUSE(Google资深工程师) 大佬讲解，个人感觉非常经典
* 随手撸**网络版外部排序**，培养自己的基本编码素质
* PPT 文字脱敏⬇️

大部分 Go 语言的重心都放在语法的特点与小细节上面

`ccmouse 大佬` 学习新语言的特点后，找一个不那么简单的项目去做
* 边做边看文档
* 边做边查 StackOverflow

### Go 语言的项目
**完全使用 Go 语言**
* Docker
* Kubernetes
* Caddy
* CockroachDB

**部分使用Go语言**
* MongoDB/Couchbase
* Dropbox
* Uber
* Google

### Go 语言的发展趋势
* [https://octoverse.github.com/](https://octoverse.github.com/)
* [https://www.indeed.com/jobtrends/q-golang.html](https://www.indeed.com/jobtrends/q-golang.html)

### Go 语言在中国
* [https://trends.google.com/trends/explore?q=%2Fm%2F09gbxjr](https://trends.google.com/trends/explore?q=%2Fm%2F09gbxjr)

### Google 内部的 “标准” 编程语言
* C++：必须有性能保障的部分，如搜索引擎
* Java：复杂业务逻辑，如 adwords, google docs
* Python：大量内部工具
* Go：新的内部工具，及其他业务模块，如dl.google.com

### Go语言的设计初衷
* 如果有一门语言，有像C/C++那样的性能，可以做系统开发
* 但是没有繁琐的类型系统，有简单统一化的模块依赖管理，编译速度飞快
* 如果有一门语言，像Java那样拥有垃圾回收
* 但是更快，对业务的影响更小
* 如果有一门语言，像 Python 那样简单易学，拥有灵活的类型，支持函数式编程，异步IO
* 但是有编译器进行静态类型检查
* 如果有一门语言，针对上述痛点进行设计，并加入并发编程
* 这就是 Go 语言

