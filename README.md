# Gopher

![Gopher](https://golang.org/doc/gopher/frontpage.png)

Go Programming Language 扫盲

go1.11.2

## 软件开发人员硬基础

目的：注重**基础知识**和**思维方式**

### linux 机器上使用 top 命令

* p 键 - 按 cpu 使用率排序
* m 键 - 按内存使用量排序

### Mac 下 top 命令

* 先输入 o，然后输入 cpu 则按 cpu 使用量排序，输入 rsize 则按内存使用量排序。

* 记不清了，可以在 top 的界面上按 ?，在弹出的帮助界面中即可看到。

### 脑图备案

* [Go 网络排序](./images/sort.png)
* [Go 基础语法脑图回顾](./images/Go1.png)

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

### 安装 MySQL 并初始化表

![安装 MySQL 并初始化表](./images/RESTful/api-7.png)

运行MySQL

```sh
docker run --name mysql -v ~/dockerdata:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -d -i -p 3306:3306 --restart=always  mysql:5.6

docker cp api-03/db.sql e60fcab227d4:/tmp/

docker exec -ti mysql bash

mysql -uroot -p123456

source  /tmp/db.sql
```

### 初始化 MySQL 数据库并建立连接

![初始化 MySQL 数据库并建立连接](./images/RESTful/api-8.png)

### 自定义业务错误信息

![自定义业务错误信息](./images/RESTful/api-9.png)

### 读取和返回 HTTP 请求

![读取和返回 HTTP 请求](./images/RESTful/api-10.png)

### 用户业务逻辑处理

![用户业务逻辑处理](./images/RESTful/api-11.png)

### HTTP 调用添加自定义处理逻辑

![HTTP 调用添加自定义处理逻辑](./images/RESTful/api-12.png)

### API 身份验证

![API 身份验证](./images/RESTful/api-13.png)

### 用 HTTPS 加密 API 请求

![用 HTTPS 加密 API 请求](./images/RESTful/api-14.png)

API Server 添加 HTTPS 支持

生成私钥文件（server.key）和自签发的数字证书（server.crt）

```sh
openssl req -new -nodes -x509 -out conf/server.crt -keyout conf/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=127.0.0.1/emailAddress=xxxxx@qq.com"
```

### 用 Makefile 管理 API 项目

![用 Makefile 管理 API 项目](./images/RESTful/api-15.png)

### 给 API 命令增加版本功能

![给 API 命令增加版本功能](./images/RESTful/api-16.png)

### 给 API 增加启动脚本

![给 API 增加启动脚本](./images/RESTful/api-17.png)

### 基于 Nginx 的 API 部署方案

![基于 Nginx 的 API 部署方案](./images/RESTful/api-18.png)

### API 高可用方案

![API 高可用方案](./images/RESTful/api-19.png)

### go test 测试你的代码

![go test 测试你的代码](./images/RESTful/api-20.png)

### API 性能分析

![API 性能分析](./images/RESTful/api-21.png)

### 生成 Swagger 在线文档

![生成 Swagger 在线文档](./images/RESTful/api-22.png)

### API 性能测试和调优

![API 性能测试和调优](./images/RESTful/api-23.png)


## IRIS+XORM 实战

![IRIS XORM 实战](./images/iris+xorm1.png)

## 高并发，高性能实战

### 抽奖系统

![抽奖系统](./images/Lottery/lottery-1.png)

### 常见抽奖活动

![常见抽奖活动](./images/Lottery/lottery-2.png)

### 系统设计和架构设计

![系统设计和架构设计](./images/Lottery/lottery-3.png)

###  项目框架与核心代码

![项目框架与核心代码](./images/Lottery/lottery-4.png)

###  后台功能开发

![后台功能开发](./images/Lottery/lottery-5.png)

### 基于mysql的抽奖功能开发

![基于mysql的抽奖功能开发](./images/Lottery/lottery-6.png)

### GO实现千万级WebSocket消息推送服务

![GO实现千万级WebSocket消息推送服务](./images/go-websocket.png)

## 使用 GoLang 构建高性能网络游戏服务器

![使用 GoLang 构建高性能网络游戏服务器](./images/nano/nano-1.png)

