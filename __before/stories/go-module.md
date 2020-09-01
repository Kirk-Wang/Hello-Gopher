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

