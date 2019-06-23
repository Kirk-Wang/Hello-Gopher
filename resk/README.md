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



