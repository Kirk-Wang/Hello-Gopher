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
  # "acknowledged": true 表示安全写级别被启用
  # 由于我们在 db.collection.insertOne() 命令中并没有提供 writeConcern 文档，
  # 这里显示的是 mongoDB 默认的安全写级别启用状态
  # "insertedId"显示了被写入的文档的_id

  > show collections
  accounts

  # db.collection.insertOne() 命令会自动创建相应的集合

  # 如果 db.collection.insertOne() 遇到了错误……
  # 处理抛出的错误

  > try{
      db.accounts.insertOne({
        _id: "account1",
        name: "bob",
        balance: 50
      })
    } catch(e) {
      print(e)
    }
  WriteError({
        "index" : 0,
        "code" : 11000,
        "errmsg" : "E11000 duplicate key error collection: test.accounts index: _id_ dup key: { : \"account1\" }",
        "op" : {
                "_id" : "account1",
                "name" : "bob",
                "balance" : 50
        }
  })

  # 自动生成_id
  # 省略创建文档中的 _id 字段
  > db.accounts.insertOne({name: "bob", balance: 50})
  {
        "acknowledged" : true,
        "insertedId" : ObjectId("5d148cf3aae4f17ec2d60582")
  }
  ```
  
  创建多个文档

  ```sh
  db.collection.insertMany()

  db.<collection>.insertMany(
    [<document1>,<document2>,...],
    {
      wirteConcern: <document>,
      ordered: <boolean>
    }
  )
  ```
  将多个需要创建的文档作为一个数组传入 `db.collection.insertMany()`

  ordered 参数用来决定 mongoDB 是否要按顺序来写入这些文档

  如果将 ordered 参数设置为 false，mongoDB 可以打乱文档的写入顺序，以便优化写入操作的性能

  ordered 参数的默认值为 true

  将文档写入 accounts 集合
  ```sh
  >db.accounts.insertMany([
    {
      name: "charlie",
      balance: 500
    },
    {
      name: "david",
      balance: 200
    }
  ])
  {
        "acknowledged" : true,
        "insertedIds" : [
                ObjectId("5d148f63aae4f17ec2d60583"),
                ObjectId("5d148f63aae4f17ec2d60584")
        ]
  }
  ```

  如果 db.collection.insertMany() 遇到了错误...

  在顺序写入时遇到了错误
  ```sh
  >try{
    db.accounts.insertMany([
      { _id: "account1", name: "charlie", balance: 500},
      { name: "david", balance: 200 }
    ])
  }catch(e) {
    print(e)
  }
  BulkWriteError({
        "writeErrors" : [
                {
                        "index" : 0,
                        "code" : 11000,
                        "errmsg" : "E11000 duplicate key error collection: test.accounts index: _id_ dup key: { : \"account1\" }",
                        "op" : {
                                "_id" : "account1",
                                "name" : "charlie",
                                "balance" : 500
                        }
                }
        ],
        "writeConcernErrors" : [ ],
        "nInserted" : 0, # 一条都没写进去
        "nUpserted" : 0,
        "nMatched" : 0,
        "nModified" : 0,
        "nRemoved" : 0,
        "upserted" : [ ]
  })
  ```

  在乱序写入是遇到错误

  ```sh
  >try{
    db.accounts.insertMany([
      { _id: "account1", name: "charlie", balance: 500},
      { name: "david", balance: 200 }
    ], { ordered: false })
  }catch(e) {
    print(e)
  }
  BulkWriteError({
        "writeErrors" : [
                {
                        "index" : 0,
                        "code" : 11000,
                        "errmsg" : "E11000 duplicate key error collection: test.accounts index: _id_ dup key: { : \"account1\" }",
                        "op" : {
                                "_id" : "account1",
                                "name" : "charlie",
                                "balance" : 500
                        }
                }
        ],
        "writeConcernErrors" : [ ],
        "nInserted" : 1, // ^_^,虽然报错，但有一篇写入了
        "nUpserted" : 0,
        "nMatched" : 0,
        "nModified" : 0,
        "nRemoved" : 0,
        "upserted" : [ ]
  })
  ```

  顺序写入时，一旦遇到错误，操作便会退出，剩余的文档无论正确与否，都不会被写入

  乱序写入时，即使某些文档造成了错误，剩余的正确文档仍然会被写入

  `db.collection.insert()` 创建单个或多个文档

  ```sh
  > db.accounts.insert({ name: "george", balance: 1000 })
  WriteResult({ "nInserted" : 1 })
  ```

  遇到错误
  ```sh
  > db.accounts.insert([
  ...       { _id: "account1", name: "george", balance: 1000 },
  ...       { name: "henry", balance: 2000 },
  ...     ])
  BulkWriteResult({
        "writeErrors" : [
                {
                        "index" : 0,
                        "code" : 11000,
                        "errmsg" : "E11000 duplicate key error collection: test.accounts index: _id_ dup key: { : \"account1\" }",
                        "op" : {
                                "_id" : "account1",
                                "name" : "george",
                                "balance" : 1000
                        }
                }
        ],
        "writeConcernErrors" : [ ],
        "nInserted" : 0,
        "nUpserted" : 0,
        "nMatched" : 0,
        "nModified" : 0,
        "nRemoved" : 0,
        "upserted" : [ ]
  })
  ```
  insertOne, insertMany 和 insert 的区别

  三个命令返回的结果文档格式是不一样的

  insertOne 和 insertMany 命令不支持 db.collection.explain() 命令
  
  insert 支持 db.collection.explain() 命令

  `db.collection.save()`

  他会调用 `db.collection.insert()`命令，所以返回的结果和它是一样的

  再来看一下文档主键 _id

  默认的对象主键 objectId

  ```sh
  > ObjectId()
  ObjectId("5d1498afc78d64aa5f9e45f0")
  > ObjectId("5d1498afc78d64aa5f9e45f0")
  ObjectId("5d1498afc78d64aa5f9e45f0")
  ```
  提取 ObjectId 的创建时间

  ```sh
  > ObjectId("5d1498afc78d64aa5f9e45f0").getTimestamp()
  ```

  复合主键

  可以使用文档作为文档主键, 复合主键仍然要满足文档主键的唯一性
  ```sh
  > db.accounts.insert({
  ...     _id: { accountNo: "001", type: "saving" },
  ...     name: "irene",
  ...     balance: 80
  ...   })
  WriteResult({ "nInserted" : 1 })
  > db.accounts.insert({
  ...     _id: { type: "saving", accountNo: "001" },
  ...     name: "irene",
  ...     balance: 80
  ...   })
  WriteResult({ "nInserted" : 1 })
  # 注意：字段顺序调换，同样能写入
  ```

#### 创建文档总结

  * 使用 db.collection.insertOne() 创建单一文档
  * 使用 db.collection.insertMany() 创建多个文档
  * 使用 db.collection.insert() 创建单一或多个文档
  * 创建文档命令返回的结果/错误
  * 使用 db.collection.save() 创建单一文档
  * 对象主键 ObjectId
  * 复合主键

#### 读取文档

  * db.collection.find()
  * 匹配查询
  * 查询操作符

  游标

  * 查询操作返回的结果游标
  * 游标的迭代与操作

  投射

  * 只返回部分字段
  * 内嵌文档的投射
  * 数组的投射

  演示：

  ```sh
  db.accounts.find()

  db.accounts.find().pretty() # 格式化

  db.accounts.find({name: "alice"}) # 匹配查询
  db.accounts.find({name: "alice", balance: 100})

  db.accounts.find({"_id.type":"saving"}) # 使用复合主键查询

  ```

  比较操作符

  ```sh
  db.accounts.find({name: {$eq: "alice"}}) # 比较运算符(等于)

  db.accounts.find({name: {$ne: "alice"}}) # 不等于
  db.accounts.find({balance: {$ne: 100}})
  # 注意： $ne 也会筛选出并不包含查询字段的文档
  db.accounts.find({"_id.type": { $ne: "saving" }})
  db.accounts.find({balance:{ $gt: 500 }}) # 大于

  # 读取用户名字排在 fred 之前的银行账户文档
  db.accounts.find({balance:{ $lt: "fred" }}) # 小于
  db.accounts.find({name:{$in:["alice", "charlie"]}}) # 属于 alice 和 charlie 的文档

  db.accounts.find({name:{$nin:["alice", "charlie"]}}) # 不属于
  db.accounts.find({"_id.type": { $nin: ["saving"] }}) # 注意：和 $ne 一样
  ```

  逻辑操作符 $not, $and, $or, $nor

  ```sh
  db.accounts.find({ balance: { $not: { $lt: 500 } } }) # 不小于500
  # $not 也会筛选出并不包含查询字段的文档
  db.accounts.find({ "_id.type": { $not: { $eq: "saving" } } })
  # 读取余额大于 100 并且用户姓名排在 fred 之后的银行账户文档
  db.accounts.find({
    $and: [
        {balance: { $gt:100 } },
        {name: { $gt: "fred"} }
    ]
  })

  # 当筛选条件应用在不同字段上时，可以省略 $and 操作符
  db.accounts.find({
    balance: { $gt:100 },
    name: { $gt: "fred"}
  })
  ```

  当筛选条件应用在同一个字段上时，也可以简化命令

  "读取余额大于100并且小于500的银行账户文档"
  ```sh
  db.accounts.find({ balance: { $gt: 100, $lt: 500 } })
  ```

  "读取属于 alice 或者 charlie 的银行账户文档"
  ```sh
  db.accounts.find({
    $or: [
      { name: { $eq: "alice" } },
      { name: { $eq: "charlie" } }
    ]
  })
  ```
  当所有筛选条件使用的都是 $eq 操作符时，$or 和 $in 的效果是相同的

  ```sh
  db.accounts.find( { name: { $in: [ "alice", "charlie" ] } } )
  ```
  "读取余额小于100或者大于500的银行账户文档"
  ```sh
  db.accounts.find({
    $or: [
      { balance: { $lt: 100 } },
      { balance: { $gt: 500 } }
    ]
  })
  ```

  "读取不属于 alice 和 charlie 且余额不小于 100 的银行账户文档"
  ```sh
  db.accounts.find({
    $nor:[
      { name: "alice" },
      { name: "charlie" },
      { balance: { $lt: 100 } },
    ]
  })
  ```

  $nor 也会筛选出并不包含查询字段的文档

  "读取账户类型不是储蓄账户且余额大于500的银行账户文档"
  ```sh
  db.accounts.find({
    $nor: [
      {"_id.type":"saving"},
      {balance: { $gt: 500 } }
    ]
  })
  ```
  字段操作符 $exists, $type

  "读取包含账户类型字段的银行账户文档"
  ```sh
  db.accounts.find({ "_id.type": { $exists: true } })
  ```
  回想一下，之前介绍的有些操作符会筛选出不包含查询字段的文档

  "读取账户类型不是支票账户的银行账户文档"
  ```sh
  db.accounts.find({ "_id.type": { $ne: "checking" } })
  ```
  如果增加一个 $exists 操作符，就可以得到更准确的筛选结果
  ```sh
  db.accounts.find({"_id.type": { $ne: "checking", $exists: true } })
  ```
  "读取文档主键是字符串的银行账户文档"
  ```sh
  db.accounts.find({ _id: { $type: "string" } })
  ```
  "读取文档主键是对象主键或者是复合主键的银行账户文档"
  ```sh
  db.accounts.find({
    _id: {
      $type: ["objectId", "object"]
    }
  })
  ```