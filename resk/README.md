### 红包业务和业务模型概述

#### 红包业务场景

通过移动互联网应用来发红包成为了常态

* 发红包场景：
  * 表白，祝福等
  * 庆祝，营销，显摆，曝光，求问等

#### 红包业务的定义

红包：一定数量和金额的红包序列

* 红包是具备虚拟资金特征的特殊商品
* 发红包和收红包实际上是资金交易过程
* 资金交易是资金从一个账户流向另一个账户的过程

#### 红包业务模型

* 资金账户
* 红包🧧

### 数据库物理模型设计

从业务领域模型来推到数据库物理模型

* 物理模型和逻辑模型保持一致，可以保持不一致
* 不能违背业务模型逻辑的前提下可以做一些优化
* 冗余、合并、拆分、异构


### 代码架构概述

- **Why:为什么要做代码架构？**

- **What:代码架构要做些什么事情？**

- **How:怎样设计代码架构？**

#### Why:为什么要做代码架构？

需求分析-》概要设计-》详细设计（**代码架构**）-》开发-》测试-》部署-》运维

**代码架构承上启下，决定软件质量**

- 承上：说明业务逻辑和业务领域模型
- 本身：保证代码有更好的可读性和可维护性、可扩展性
- 启下：承载了代码运行的硬件部署架构

#### What:代码架构要做些什么事情？

业务逻辑表达的职责

- 向上沟通的职责，提供交互入口
- 自身业务逻辑和技术实现的职责
- 向下沟通的职责，保存运行状态

#### How:怎样设计代码架构？

分层架构

- 单一职责
- 高内聚低耦合
- 提高复用性

三层架构

- 表现层
- 业务逻辑层
- 数据持久层

逻辑分层结构

- 用户接口：人机交互 (User Interface)
- 应用层 (Application)
- 领域层 (Domain)
- 基础设施层 (Infrastructure)

物理分层

- 用户接口 (User Interface)
- 应用服务层接口（Service Interface）
- 核心层 (Core)
  * 应用服务实现层 (Sevice Impl)
  * 领域层 (Domain)
  * 数据访问层 (Repository/Dao)
- 基础设施层 (Infrastructure)

红包资金账户模块-代码架构设计

- User Interface
  - AccountApi
- Application
  - AccountService(interface)
  - accountService
- Domain
  - accountDomain
- Infrastructure
  - AccountDao
  - AccountLogDao
  - IdGenerator

红包模块代码架构设计

- User Interface
  - EnvelopeApi
- Application
  - EnvelopeService(interface)
  - envelopeService
- Domain
  - envelopeDomain
- Infrastructure
  - GoodsDao
  - ItemDao
  - IdGenerator

### Go 编程中的一些规范

#### Golang 包名命名规则

* 完整包名的组成：引入路径 + 包名
* 源代码中 package 名称
  * 可以和文件夹名称不一致，建议尽量一致
  * 同一文件夹中所有源文件的包名必须一致
* 代码引用时使用包名，非文件夹名称
* 源代码 import 的是文件夹路径名称
  * 非包名
  * 非文件名

#### Golang 源代码文件名命名规则

**Go 语言源码文件名称**
* 文件名称只是描述性的，无编程含义

#### 红包系统 Golang 代码结构设计

项目空间主目录：（存放项目的位置
  - src (必选，项目开始的地方)
    - 导入路径1+项目名称1 (项目1)
    - 导入路径2+项目名称2 (项目2)
  - bin (可选，编译后的可执行命令)
  - pkg (可选，编译后的包文件)

#### 红包系统-包结构设计

resk 项目主目录
  - core 核心：应用层、领域、持久层
  - apis 用户接口层
  - brun 应用程序
  - doc  文档
  - infra 基础设施
  - services 应用层接口：应用服务

apis包：存放用户接口层

- 文件名称可以描述其业务含义的单词
- 定义外部交互逻辑和交互形式：UI、RESTful 接口
- 不涉及任何业务，随时可以替换为其他形式的交互方式
- services 构造和初始化

services包：存放应用层接口

- 文件名称使用可以描述性其业务含义
- 需要对外暴露
  - DTO、service interface
  - 枚举、常数等

core包：应用层实现，领域层、数据访问层所有代码

- 文件名称使用可以描述性其`业务含义+分层名称`
- Service实现，Domain、Dao、PO

### 使用 Go Modules 来管理依赖

go get -> vendor -> go modules

- go get 无版本概念
- vendor 曲线救国，但任然没有版本化
- go1.11 modules 开启了版本依赖新历程

### Go modules 简介

通过 GO111MODULE 环境来开启或者关闭，默认是 auto

- off/on/auto：关闭、开启、自动识别
- 使用 module 后：GOPATH 失去了部分意义
- 要用 module，第一步将项目从 GOPATH 中移出去

#### go.mod 文件

go.mod 文件来管理依赖，定义模块依赖

- go.mod 文件放在项目根目录
- go.mod 文件面向行，由指令+参数组成
- 注释使用 `//`
- module:定义当前模块和包路径
- require:定义依赖的模块和版本

#### go.mod 文件主要指令

- module:定义当前模块和包路径
- require:定义依赖的模块和版本
- exclude:排除特定模块和版本的使用
- replace:模块源的替换

#### go mod 命令

go.mod 文件用 go mod 命令来创建和维护

- 命令格式：go mod <命令> [可选参数]
- 8个子命令
  - init, tidy, vendor, verify
  - download, edit, graph, why

#### go mod 命令使用

- 使用 `go mod init` 来创建和初始化 go.mod 文件
- 使用 `go mod tidy` 来更新依赖模块
- 使用 `go get` 命令来下载和更新依赖包

#### replace 子指令

解决网络访问不了 golang.org/x 等谷歌的扩展包

- 用新的包去替换老的包
- 格式：replace 包路径[版本] => 包路径 版本
- golang.org/x/sys => github.com/golang/sys

```sh
replace(
  golang.org/x/sys => github.com/golang/sys latest
)
```

#### resk 项目为 module 化流程

1. 移动项目到 GOPATH 工作空间之外
2. go mod init 创建和初始化 go.mod
3. go mod tidy 来整理更新已有依赖
4. 如果存在不能下载的谷歌库，replace 指令替换
5. go verify 来验证 module
6. go mod vendor
  - 复制依赖到 `vendor` 目录，方便源代码定位和查看


#### 在 Golang 中如何设计枚举？

Golang 中枚举通过定义常数来实现

- 类型别名的形式来声明类型
- iota来自增和自动赋值

什么时候可以使用 iota, 什么时候不能使用 iota 呢？

- 无状态且非持久化，可以使用 iota
- 有状态或者需要持久化，不能使用 iota

#### 在 Golang 中如何使用 JSON?

标准库内建的 JSON 包

- json.func Marshal(v interface{})([]byte, error)
- json.Unmarshal(data []byte, v interface{}) error
- 默认使用的 JSON 字段名称是它的 `Field` 名称

不是所有的类型都能序列化

- 支持 string、bool、数字类型、数组和切片、结构体、map
- Channel、complex、function 类型无法进行 json 序列化
- 结构体重的循环数据结构，序列化时不会被处理

结构体中自定义字段名称

- JSON tag key: name,[omitempty,-]
  * omitempty 忽略空值
- [string]标记，定义 bool、浮点、整型类型使用字符串编码
- 临时添加字段：内嵌结构体 合并多个结构体

高性能 JSON 库推荐：jsoniter [json-iterator]

- 快，并且更快
- 支持 java 和 go
- 100 兼容 JSON 标准库，一行代码迁移到 `jsoniter`
- 安装：`go get github.com/json-iterator/go`

### 基础设施资源配置设计

#### 不同维度的分类

- 按内容：静态和动态
- 按环境：开发、测试、灰度、生产
- 按形式：本地文件和分布式服务
- 按格式：properties、json、xml、ini、yaml、toml 等
- 按用途：程序和应用级别
  - 配置文件命名规范：`前缀+[-_]+功能`
  - 程序级别，可以不分组，比如：boot.ini, config.ini
  - 应用级别分组：比如 app-mysql.properties
- 按环境来变量化配置项
  - 不同环境使用不同的配置文件
  - 变量化因环境不同的可变配置项

**分布式配置管理中心**


#### 配置设计

**props 是统一的配置工具库**

- 各种配置源抽象或转换为 key/value 结构
- 支持 prop, ini, zk, consul, etcd, nacos
- 支持 unmarshal
- props 配置客户端工具库加持+INI格式
- INI 文件由节、键、值组成
- INI 格式三要素：节、参数(key/value)、注释
- section:方括号包围，比如：[mysql]
- 参数：key=value, 有些工具支持冒号分割
- 注释，使用`;`,有些编辑器和工具支持 `#`
- Section 作为分组
- props 配置客户端使用完整的 key 来解释
- 完整的 key 由 Section 和参数 key 组成，`.` 分割

```go
import "https://github.com/tietang/props"
```

### 基础设施资源 mysql starter 编码实践

#### 数据库连接池概述

**为什么需要数据库连接池**

- 一种有限且昂贵的资源
- 连接过程：socket连接 > 认证 > 协议协商 > ... > 命令和结果 > 关闭
- 资源频繁分配、释放造成性能问题

**数据库连接池解决了哪些问题？** `连接复用`

- 重用连接，节省连接时间，提高性能
- 减少了内存的分配和消耗，提高资源利用
- 高并发系统下，性能提升和资源利用更为明显

**连接池原理**

- 池中维护一定数量的数据库连接
- 数据库存取时获取连接，结束时返回到连接池
- 维护连接资源的数量和质量

#### database/sql包使用

- sql.Open 函数创建 sql.DB
- sql.DB 是数据库操作的高级抽象，维护了一个连接池
- sql.DB 会自动创建和释放连接

**sql.DB连接池配置方法**

- db.SetMaxIdleConns 保持连接的最大空闲连接数
  * 默认0，返回到连接池后，会频繁的关闭和创建
- db.SetMaxOpenConns 最大连接数
  * 默认0， 无限制
- SetConnMaxLifetime 闲置连接的最大存活时间
  * 小于等于0，永远存活

sql.Open 的使用方法

- db, err := sql.Open("dirverName", "dataSourceName")
- driverName: 数据库驱动注册的名称
- [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
- 例：mysql
  - username:password@tcp(ip:3306)/dbname[?param1=value1&...&paramN=valueN]
  - username:password@tcp(hostName:3306)/dbname?charset=utf8&parseTime=true&loc=Local
    - charset: 设置字符串字符集，这个要和数据库中的 schema 保持一致
    - parseTime：程序是否自动解析时间字段
    - loc: `local` 表示使用操作系统时区

sql.DB 主要功能

- db.Exec/ExecContext() 执行查询而不返回数据
- db.Query/QueryContext() 执行查询并返回结果
- db.QueryRow/QueryRowContext()
  - 执行查询并最多只返回1行数据
- db.Prepare/PrepareContext()
  - 创建一个预编译 Statement 对象
- db.Begin/BeginTx()
  - 开启一个事务，返回的Tx事务对象会被绑定到单个连接
- db.Stats() 返回数据库统计信息

#### dbx数据库扩展工具库

* 基于 `sql.DB` 基础之上构建的高性能数据库工具
* 安装
    ```sh
    go get -u github.com/tietang/dbx
    ```

**主要的特性和目标是：**

- 简单高效，最大限度的保留原生特性的基础
- 高性能，支持 orm 的基础上，最大限度的减少性能损耗
- 自动表名和字段名称映射，驼峰命名转换成下划线命名
- 支持自定义表名和字段名称映射

### 基础设施资源 log starter 编码实践

#### 日志的重要性

- Bug 无处不在，存在的 bug 就一定会出现的
- 系统运行日志作为系统问题和业务问题跟踪的依据
- 数据价值：监控告警、安全审计、业务数据分析

#### 日志输出的几个重要点

- 日志格式：简单，规范化和结构化
- 日志内容：关键性节点
- 日志级别：trace,`debug`,`info`,`warn`,`error`,fatal,panic

#### 日志框架

**Logrus 是一个可插拔的、结构化的日志框架**

- 完全兼容 golang 标准库日志模块
- 六种级别：debug、info、warn、error、fatal 和 panic
- 支持自定义扩展 hook 和自定义输出格式

#### LogrusStarter 编码要点

**按照本身的配置和日志输出的编写要点**

- 定义日志输出格式、日志级别
- 控制台高亮输出、功能 hook 配置
- 日志文件和滚动配置

获取系统环境变量：

```go
level := os.Getenv("log.debug")
```

### 基础设施资源验证器 starter 编码实践

#### 效验框架的重要性

- 保证数据的正确性和完整性
- 避免数据校验与业务逻辑耦合太紧
- 有助于良好的代码设计

#### validator 验证框架

- go get gopkg.in/go-playground/validator.v9
- 自称100分的验证框架
- struct 和 field 验证
  - 嵌套结构体跨 field 和 struct,
  - 支持深入 Map, Slice and Array
- 支持自定义：字段类型和错误

### 基础设施资源 web 框架 starter 编码实践



