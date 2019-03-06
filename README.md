# Gopher

![Gopher](https://golang.org/doc/gopher/frontpage.png)

Go Programming Language 扫盲(go1.11.2)

### 基础知识 & 编程能力

[CodeJam](https://codingcompetitions.withgoogle.com/codejam)

### 一道题：
* [Beautiful Numbers](https://code.google.com/codejam/contest/5264487/dashboard#s=p1)
    * 1, 11, 111 .... 是 beautiful 的
    * 3 -> 2进制 -> 11
    * 13 -> 3进制 -> 111 (√)
        * 到过来算（111）-> 1*3*3 + 1*3 + 1 = 13 (进位制)
        * 13 % 3 = 1(低位)，13 / 3 = 4
        * 4 % 3 = 1，4 / 3 = 1
        * 1 % 3 = 1，1 / 3 = 0
    * 13 -> 12进制 -> 11 (X)

#### 操作系统

* 进程(隔离)
    * 线程
    * 内存（逻辑内存）
    * 文件/网络句柄
* 线程(真正运行)
    * 栈（调用堆栈）,从主线程的入口 main 函数调用，每次调用把所有的参数和返回地址压到栈里面
    * PC (program counter)：下一条执行指令地址，放在内存
    * TLS (Thread Local Storage)：线程独立的内存
* 存储
    * 寄存器 > 缓存 > 内存 > 硬盘
* 寻址空间
    * 32位 -> 4G
    * 64位 -> ~10^19 Bytes
    * 64位JVM -> 可使用更大内存，需重新编译

寻址 int n = *p; -> MOV EAX,[EBX]

指针p --> 逻辑内存，进程独立 2^32 或 2^64 ---> 物理内存（分页，虚拟内存）---> 寄存器


#### 网络

数据链路层->网络层->传输层->应用层

* 网络传输
    * 不可靠
        * 丢包，重复包
        * 出错
        * 乱序
    * 不安全
        * IS THIS LINE SECURE？
        * 中间人攻击
        * 窃取
        * 篡改

* 例：滑动窗口（增加线路的吞吐量，按照工程化的思维理解）
    * TCP 协议中使用
    * 为维持发送方 / 接收方缓冲区
        * 发送 & Ack
    * 丢 Ack (一定是按照顺序 Ack )
        * 超时重传
* 题：
    [阿里巴巴有相距1500km的机房A和B，现有100GB数据需要通过...](https://www.nowcoder.com/questionTerminal/97a4bed9cf644832bbb8dec72afccfa8)
### Pipe

[Pipe](https://github.com/b3log/pipe) 是一款小而美的 Go 博客平台。作者写的代码十分优雅，值得借鉴。所读版本：1.8.6。

* [利用 Docker，先让它快速跑起来](./stories/pipe-docker.md)
* [go module](./stories/go-module.md)
* [本地开发，调试](./stories/pipe-debug.md)

#### Traefik & Docker

[Traefik](https://docs.traefik.io/)

[Traefik CN](http://traefik.cn/)

[用 Traefik 搭配 Docker 快速架設服務](https://blog.wu-boy.com/2019/01/deploy-service-using-traefik-and-docker/#more-7193)

项目：[drone-golang-example](https://github.com/go-training/drone-golang-example)
* Traefik VS. HAProxy,Nginx,Caddy
* 如何快速在同一台机器上假设两个服务
* 如何搭配 Drone 自动化更新服务
* 自动化部署
  * 编译 Binary
  * 上传 Docker Image
  * ssh 远端机器
    * 下载最新 Image
    * 重新启动服务


### Traefik Let's Encrypt & Docker

[Traefik Let's Encrypt & Docker](https://github.com/go-training/training/tree/master/example25-traefik-golang-app-lets-encrypt)

### (Go 相关基础) 脑图

* [Go 网络排序](./images/sort.png)
* [Go 基础语法](./images/Go1.png)
* [Go 内建容器](./images/Go2.png)
* [Go 面向“对象”](./images/Go3.png)
* [Go 面向接口](./images/Go4.png)
* [Go 函数式编程](./images/Go5.png)
* [Go 错误处理和资源管理](./images/Go6.png)
    * CACHT ALL THE ERRORS
        * defer 调用
            * 确保调用在函数结束时发生
            * 参数在 defer 语句时计算
            * defer 列表为后进先出
        * 何时使用 defer 调用
            * Open/Close
            * Lock/Unlock
            * PrintHeader/PrintFooter
        * 错误处理的概念
            * `error` 就是一个 `interface`，我们可以自己去实现它的接口
        * 如何实现统一的错误处理逻辑
            * errWrapper
        * panic
            * 停止当前函数执行
            * 一直向上返回，执行每一层的 defer
            * 如果没有遇见 recover，程序退出
        * recover
            * 仅在 defer 调用中使用
            * 获取 painc 的值
            * 如果无法处理，可重新 panic
* [Go 测试与性能调优](./images/Go7.png)
* [Goroutine](./images/Go8.png)
* [Channel](./images/Go9.png)
* [http及其他标准库](./images/Go10.png)
* [迷宫的广度优先搜索](./images/Go11.png)
* [实战项目Spider](./images/Go12.png)
* [单任务版Spider](./images/Go13.png)
* [单任务版网络利用率](./images/1.0spider.png)
* [并发版Spider](./images/Go14.png)
* [并发版网络利用率](./images/2.0spider.png)
* [Spider & top](./images/2.0top.png)
* [数据存储->Elasticsearch](./images/Go15.png)
* [分布式初探](./images/Go16.png)
* [总结](./images/Go17.png)
---

### 基于 Go 语言构建企业级的 RESTful API 服务

* [高可用 API 构建](./images/RESTful/api-1.png)
* [RESTful API](./images/RESTful/api-2.png)
* [API流程和代码结构](./images/RESTful/api-3.png)
* [启动一个最简单的RESTful API 服务器](./images/RESTful/api-4.png)
* [配置文件读取](./images/RESTful/api-5.png)
* [记录和管理 API 日志](./images/RESTful/api-6.png)
* [安装 MySQL 并初始化表](./images/RESTful/api-7.png)
* [初始化 MySQL 数据库并建立连接](./images/RESTful/api-8.png)
* [自定义业务错误信息](./images/RESTful/api-9.png)
* [读取和返回 HTTP 请求](./images/RESTful/api-10.png)
* [用户业务逻辑处理](./images/RESTful/api-11.png)
* [HTTP 调用添加自定义处理逻辑](./images/RESTful/api-12.png)
* [API 身份验证](./images/RESTful/api-13.png)
* [用 HTTPS 加密 API 请求](./images/RESTful/api-14.png)
* [用 Makefile 管理 API 项目](./images/RESTful/api-15.png)
* [给 API 命令增加版本功能](./images/RESTful/api-16.png)
* [给 API 增加启动脚本](./images/RESTful/api-17.png)
* [基于 Nginx 的 API 部署方案](./images/RESTful/api-18.png)
* [API 高可用方案](./images/RESTful/api-19.png)
* [go test 测试你的代码](./images/RESTful/api-20.png)
* [API 性能分析](./images/RESTful/api-21.png)
* [生成 Swagger 在线文档](./images/RESTful/api-22.png)
* [API 性能测试和调优](./images/RESTful/api-23.png)
* [IRIS XORM 实战](./images/iris+xorm1.png)
* [抽奖系统](./images/Lottery/lottery-1.png)
* [常见抽奖活动](./images/Lottery/lottery-2.png)
* [系统设计和架构设计](./images/Lottery/lottery-3.png)
* [项目框架与核心代码](./images/Lottery/lottery-4.png)
* [后台功能开发](./images/Lottery/lottery-5.png)
* [基于mysql的抽奖功能开发](./images/Lottery/lottery-6.png)
* [GO实现千万级WebSocket消息推送服务](./images/go-websocket.png)
* [使用 GoLang 构建高性能网络游戏服务器](./images/nano/nano-1.png)

* [红包高并发业务](./RED.md)

## Other

### 软件开发人员硬基础

目的：注重**基础知识**和**思维方式**

### linux 机器上使用 top 命令

* p 键 - 按 cpu 使用率排序
* m 键 - 按内存使用量排序

### Mac 下 top 命令

* 先输入 o，然后输入 cpu 则按 cpu 使用量排序，输入 rsize 则按内存使用量排序。

* 记不清了，可以在 top 的界面上按 ?，在弹出的帮助界面中即可看到。

### Docker 启一个 MySQL

```sh
docker run --name mysql -v ~/dockerdata:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -d -i -p 3306:3306 --restart=always  mysql:5.6
docker cp api-03/db.sql e60fcab227d4:/tmp/
docker exec -ti mysql bash
mysql -uroot -p123456
source  /tmp/db.sql
```

### API Server 添加 HTTPS 支持

生成私钥文件（server.key）和自签发的数字证书（server.crt）

```sh
openssl req -new -nodes -x509 -out conf/server.crt -keyout conf/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=127.0.0.1/emailAddress=xxxxx@qq.com"
```