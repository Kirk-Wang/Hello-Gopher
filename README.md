# Gopher

![Gopher](https://golang.org/doc/gopher/frontpage.png)

Go Programming Language 扫盲(go1.11.2)

### 基础知识 & 编程能力
### 读 Pipe（一本好书🤦‍♀️，可以学到生产级别的CRUD🤣）

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

### 红包 ❤️ 高并发

需求分析->概要设计->详细设计->开发->测试->部署->运维 

* 需求分析方法和用例定义方法
* 四色建模法和核心模型设计
* 系统架构设计
* 数据库物理模型设计
* Golang编程实践
* 资金交易安全问题探讨
* 超卖问题的解决方案
* 表结构设计中的性能优化考虑
* 红包拆解和红包算法
* 系统设计和算法的选择

基本架构概览

* layer 1. 前台应用：web mobile
* layer 2. 红包网关
* layer 3. 红包系统：红包商品，用户，账户，红包订单，秒杀活动，支付
* layer 4. 基础设施：缓存，数据库，消息队列
* 技术保障：监控告警，配置管理中心，日志平台，服务注册中心

红包系统演进之路

* 满足红包业务需求，快速迭代上线
* 出现超卖现象，事务锁来帮忙
* 流量增加，收红包出现性能瓶颈，改为乐观锁，性能提高3倍
* 流量继续增加，乐观锁也扛不住了，那就上缓存吧
* 缓存还没跑溜，服务器挂了，数据丢失了？
* 分布式消息队列来解决异步写
* 数据分片来解决数据库横向扩展？
* Golang 编码实战从 0 到 1 构建红包秒杀系统









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