### MongoDB 是什么？

存储 `文档` 的 `非关系型` 数据库

数据库--》集合--》文档

### 一行命令在 Docker 中运行 MongoDB
```sh
docker pull mongo:4 # 下载 MongoDB 的官方 docker 镜像
docker images # 查看下载的镜像

docker run --name mymongo -v /mymongo/data:/data/db -d mongo:4
# --name mymongo --> 容器名字
# -v /mymongo/data:/data/db --> 挂载数据目录
# -d --> 后台运行容器

docker ps # 查看 docker 容器状态

docker logs mymongo # 查看数据库服务器日志

# Mongo Express 是一个基于网络的 MongoDB 数据库管理界面
# 下载 mongo-express 镜像
docker pull mongo-express

# 运行 mongo-express
docker run --link mymongo:mongo -p 8081:8081 mongo-express
```

### Mongo Shell

mongo shell 是用来操作 MongoDB 的 javascript 客户端界面

运行 mongo shell

```sh
docker exec -it mymongo mongo
# docker exec -it b266b3552bcb mongo -u root -p 123456

help
```

### 基本操作

  * `C`reate 创建
  * `R`ead   读取
  * `U`pdate 更新
  * `D`elete 删除

#### 文档主键 _id

  * 文档主键的唯一性
  * 支持所有数据类型（数组除外）
  * 复合主键

#### 对象主键 ObjectId

  * 默认的文档主键
  * 可以快速生成的12字节id
  * 包含创建时间（前4个字节，创建的时间，精确到秒）
  * 大部分情况下，可以认为对象主键的顺序就是创建时间的顺序
  * 特殊情况

  **`多个文档同一秒钟存储到数据库中，对象主键的精确度没有办法区分这几个文档被创建的顺序`**

  **`由于对象主键是在客户端驱动生成的，如果各个客户端的系统时间不同，也会使得对象主键的顺序和文档创建时间的顺序不匹配`**

#### 创建文档

  * db.collection.insert()
  * db.collection.save()
  * 创建多个文档
  
  进入 shell，实战

  ```sh
  use test # 使用 test 数据库
  show collections # 查看 test 数据库中的集合
  ```

  开始创建第一个文档

  ```sh
  db.collection.insertOne()

  db.<collection>.insertOne(
    <document>,
    {
      writeConcern: <document>
    }
  )
  # 这里的 <collection> 要替换成文档将要写入的集合
  # 这里的 <document> 要替换成将要写入的文档的本身
  # 这里的 writeConcern 文档定义了本次文档创建操作的安全写级别
  # 简单来说，安全写级别用来判断一次数据库写入操作是否成功
  # 安全写级别越高，丢失数据的风险就越低，然而写入操作的延迟也可能更高
  # 如果不提供 writeConcern 文档，mongoDB 使用默认的安全写级别

  # 准备写入数据库的文档

  {
    _id: "account1",
    name: "alice",
    balance: 100 # 余额
  }

  # 将文档写入 accounts 集合

  > db.accounts.insertOne({
  ...     _id: "account1",
  ...     name: "alice",
  ...     balance: 100
  ...   })
  { "acknowledged" : true, "insertedId" : "account1" }

  # 注意返回结果

  ```

  