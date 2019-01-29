# Gopher

![Gopher](https://golang.org/doc/gopher/frontpage.png)

Go Programming Language 扫盲(go1.11.2)

### 读 Pipe（一本好书🤦‍♀️，可以学到生产级别的CRUD🤣）

[Pipe](https://github.com/b3log/pipe) 是一款小而美的 Go 博客平台。作者写的代码十分优雅，值得借鉴。

所读版本：1.8.6

#### 利用 Docker，先让它快速跑起来

下载 [SQLite](https://www.sqlite.org/download.html)，我的是 Mac，大家对应着下载。

[SQLite Database Browser](https://github.com/sqlitebrowser/sqlitebrowser)

我的是 Mac, 已经自带了。

1. 在项目根目录，利用 DB Browser for SQLite 创建一个 pipe.db。

2. 我这边更新下它的 Dockerfile，主要是升级到 1.8.6
```sh
FROM alpine:3.7
LABEL maintainer = "abcdsxg@gmail.com"

ENV PIPE_VERSION 1.8.6
ENV GLIBC_VERSION 2.28-r0

WORKDIR /opt/

RUN set -ex && \
    apk --no-cache add ca-certificates && \
    wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub && \
    wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/glibc-${GLIBC_VERSION}.apk && \
    apk add glibc-${GLIBC_VERSION}.apk && \
    wget -O pipe${PIPE_VERSION}.zip https://github.com/b3log/pipe/releases/download/v${PIPE_VERSION}/pipe-v${PIPE_VERSION}-linux.zip && \
    unzip pipe${PIPE_VERSION}.zip && \
    chmod +x pipe && \
    rm -f pipe${PIPE_VERSION}.zip glibc-${GLIBC_VERSION}.apk

CMD ["/opt/pipe"]
```

3. 安装 docker & docker-compose 很简单，大家官网 step by step 跟着走就 OK 了。
```sh
docker-compose --help # 看下帮助
# up                 Create and start containers
# down               Stop and remove containers, networks, images, and volumes

docker-compose up --help
#     -d, --detach               Detached mode: Run containers in the background,
#                                print new container names. Incompatible with
#                                --abort-on-container-exit.

docker-compose up -d # 起来，然后背后运行
docker-compose down # 销毁
```

4. 进入 [http://localhost:5897](http://localhost:5897)

5. 可以先使用本地账号初始化


### go module(go1.11)

统一包管理工具，开发时不需要关心 GoPath，任何一个目录都可以进行开发。

这里有两篇文章，大家可以扫盲一下：
* [golang包管理解决之道——go modules初探](https://www.cnblogs.com/apocelipes/p/9534885.html)
* [再探go modules：使用与细节](https://www.cnblogs.com/apocelipes/p/10295096.html)

当前 pipe 项目，采用是 go1.10.x 的版本进行迭代，所以用了 vendor folder 来解决项目用到的第三方的package。vendor 用来解决不同项目用到不同版本的同一个 package。[govendor](https://github.com/kardianos/govendor) 工具。

在 go1.11 中 `GO111MODULE` 默认是 `auto`(disabled)，需要`export GO111MODULE=on`。

*现在我们需要将 vender.json 转移到 go mod*
```sh
export GO111MODULE=on # 启用 go module

# creating new go.mod: module github.com/b3log/pipe
# copying requirements from vendor/vendor.json
go mod init github.com/b3log/pipe

go clean -i -x -modcache # 先清掉所有的东西

rm -rf ./vendor # 砍掉 vendor

go mod download

go build -i -v # build 一下

echo $GOPATH
# /Users/zoot/.gvm/pkgsets/go1.11.2/global

ls /Users/zoot/.gvm/pkgsets/go1.11.2/global/pkg/mod/github.com/
# 相关的包都下载到了这里

```

* 使用 go module
  * 将项目 vendor 转移到 go module
  ```sh
  # 启用 go module
  export GO111MODULE=on
  # copying requirements from vendor/vendor.json
  go mod init github.com/gin-contrib/expvar
  # 砍掉
  rm -rf vendor
  # download & test
  go test -v . 
  # /Users/zoot/.gvm/pkgsets/go1.11.2/global
  echo $GOPATH
  # 相关的包都下载到了这里
  ls /Users/zoot/.gvm/pkgsets/go1.11.2/global/pkg/mod/github.com/ 
  # 用指定版本
  go get -u github.com/gin-gonic/gin@v1.3.0
  ```
  * 新项目使用 go module
  ```sh
  go mod tidy
  ```
* 使用 Travis 整合 go module
  * [govendor 和 go mod 同时支持--Gin 的 travis ](https://github.com/gin-gonic/gin/blob/master/.travis.yml)


#### 本地开发，调试

1. 我本机开发环境
```sh
# vscode
# go version go1.11.2 darwin/amd64
```



----

### (Go 相关基础) 脑图

* [Go 网络排序](./images/sort.png)
* [Go 基础语法](./images/Go1.png)
* [Go 内建容器](./images/Go2.png)
* [Go 面向“对象”](./images/Go3.png)
* [Go 面向接口](./images/Go4.png)
* [Go 函数式编程](./images/Go5.png)
* [Go 错误处理和资源管理](./images/Go6.png)
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