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
  "读取用户姓名是null的银行账户文档"
  ```sh
  db.accounts.find({
    name:{
      $type: "null"
    }
  })
  ```
  "也可以使用对应的 BSON 类型序号作为 $type 操作符的参数"
  ```sh
  db.accounts.find({ _id: { $type: 2 } })
  ```
  
  数组操作符：$all, $elemMatch

  "创建包含数组和嵌套数组的文档"

  ```sh
  db.accounts.insert([
    {
      name: "jack",
      balance: 2000,
      contact: ["11111111", "Alabama", "US"]
    },
    {
      name: "karen",
      balance: 2500,
      contact: [
        ["22222222","33333333"],
        "Beijing",
        "China"
      ]
    },
  ])
  ```
  "读取联系地址位于中国北京的银行账户文档"
  ```sh
  db.accounts.find({ contact: { $all: ["China", "Beijing"] } })
  ```
  "读取联系电话包含22222222和33333333的银行账户文档"
  ```sh
  db.accounts.find({ contact: { $all: [ ["22222222","33333333"] ] }})
  ```
  "读取联系电话范围在10000000至20000000之间的银行账户文档"
  ```sh
  db.accounts.find({
    contact: {
      $elemMatch: {
        $gt: "10000000",
        $lt: "20000000"
      }
    }
  })
  ```

  "读取包含一个在10000000至20000000之间，和一个在20000000至30000000之间的联系电话的银行账户文档"
  ```sh
  db.accounts.find({
    contact: {
      $all: [
        { $elemMatch: { $gt: "10000000", $lt: "20000000"} },
        { $elemMatch: { $gt: "20000000", $lt: "30000000"} },
      ]}
  })
  ```
  运算操作符
  
  $regex 匹配满足正则表达式的文档

  "读取用户姓名以 c 或者 j 开头的银行账户文档"
  ```sh
  db.accounts.find({ name: { $in: [ /^c/, /^j/ ] } })
  ```

  "读取用户姓名包含LIE(不区分大小写)的银行账户文档"
  ```sh
  db.accounts.find({name: {$regex:/LIE/,$options:'i'} })
  ```

  文档游标

  db.collection.find()返回一个文档集合游标

  在不迭代游标的情况下，只列出前 20 个文档
  ```sh
  var myCursor = db.accounts.find();
  myCursor
  ```
  我们也可以直接使用游标下标直接访问文档集合中的某一文档

  ```sh
  var myCursor = db.accounts.find();
  myCursor[1]
  ```

  游历完游标中所有的文档之后，或者在10分钟之后，游标便会自动关闭

  可以使用 noCursorTimeout() 函数来保持游标一直有效
  ```sh
  var myCursor = db.accounts.find().noCursorTimeout()
  ```
  在这之后，在不遍历游标的情况下，你需要主动关闭游标
  ```sh
  myCursor.close()
  ```

  游标函数
  ```sh
  cursor.hasNext()
  cursor.next()

  var myCursor = db.accounts.find({name:"george"})
  while(myCursor.hasNext()) {
    printjson(myCursor.next())
  }

  cursor.forEach()
  var myCursor = db.accounts.find({name:"george"})
  myCursor.forEach(printjson)

  cursor.limit()
  db.accounts.find({name:"george"}).limit(1)
  db.accounts.find({name:"george"}).limit(0) # 0 不会有任何效果

  cursor.skip()
  db.accounts.find({name:"george"}).skip(1)

  cursor.count()
  cursor.sort()
  ```
  cursor.count(<applySkipLimit>)

  默认情况下，<applySkipLimit> 为 false, 即 cursor.count() 不会考虑cursor.skip() 和 cursor.limit() 的效果
  ```sh
  db.accounts.find({name:"george"}).limit(1).count()
  # 3
  db.accounts.find({name:"george"}).limit(1).count(true)
  # 1
  ```
  在不提供筛选条件时，cursor.count() 会从集合的元数据 Metadata 中取得结果
  ```sh
  db.accounts.find().count()
  # 13
  ```
  当数据库分布式结构较为复杂时，元数据中的文档数量可能不准确

  在这种情况下，应该避免应用不提供筛选条件的 cursor.count() 函数，而使用聚合管道来计算文档数量

  cursor.sort(<document>)

  这里的<document>定义了排序的要求{ field: ordering }

  1表示由小及大的正向排序，-1表示逆向排序

  "按照余额从大到小，用户姓名按字母排序的方式排列银行账户文档"

  ```sh
  db.accounts.find().sort({ balance: -1, name: 1 })
  ```

  "读取余额最大的银行账户文档"
  ```sh
  db.accounts.find().sort( { balance: -1 } ).limit(1)
  ```
  cursor.skip() 在 cursor.limit() 之前执行
  ```sh
  db.accounts.find().limit(5).skip(3) # 会先执行 skip(3)
  > 5 篇文档
  ```
  cursor.sort() 在 cursor.skip() 和 cursor.limit() 之前执行
  ```sh
  db.accounts.find().skip(3).limit(5).sort({balance: -1})
  ```

  文档投影
  ```sh
  db.collection.find(<query>, <projection>)
  ```
  可以有选择性的返回文档中部分字段
  ```sh
  { field: inclusion }
  ```
  1 表示返回字段，0表示不返回字段

  "只返回银行账户文档中的用户姓名"
  ```sh
  db.accounts.find({}, {name:1})
  ```

  "只返回银行账户文档中的用户姓名(不包括文档主键)"
  ```sh
  db.accounts.find({}, {name:1, _id: 0})
  ```
  "不返回银行账户文档中的用户姓名(也不返回文档主键)"
  ```sh
  db.accounts.find({}, {name:0, _id: 0})

  db.accounts.find({}, {name:1, blance: 0, _id: 0}) # 报错
  ```
  除了文档主键之外，我们不可以在投影文档中混合使用包含和不包含这两种投影操作

  在数组字段上使用投影

  $slice 操作符可以返回数组字段中的部分元素
  ```sh
  db.accounts.find({}, {_id:0, name: 1, contract: 1})
  db.accounts.find({}, {_id:0, name: 1, contract: { $slice: 1} })
  db.accounts.find({}, {_id:0, name: 1, contract: { $slice: -1} })
  db.accounts.find({}, {_id:0, name: 1, contract: { $slice: [1, 2]} }) # skip & limit

  ```

  $elemMatch 和 $ 操作符可以返回数组字段中满足筛选条件的第一个元素
  ```sh
  db.accounts.find({},{
    _id: 0,
    name: 1,
    contact: {
      $elemMatch: {
        $gt: "Alabama"
      }
    }
  })
  ```
  共用筛选条件
  ```sh
  db.accounts.find(
    { contract: { $gt: "Alabama" } },
    { _id: 0, name: 1, "contact.$": 1 }
  )
  ```

  总结一下

  使用 db.collection.find() 读取文档
  * 比较操作符
  * 逻辑操作符
  * 字段操作符
  * 数组操作符
  * 运算操作符
  * 文档游标函数
  * 文档投影操作

#### 更新文档
* db.collection.update()
* db.collection.findAndModify()
* db.collection.save()
* 更新文档操作符
* 更新多个文档

"将 alice 的账户余额更改为 123"
```sh
db.account.update({name:"alice"}, { name:"alice", balance: 123 })
```

文档主键_id是不可以更改的

更新整篇文档的操作只能应用在单一文档上

"查看 jack 的银行账户"
```sh
db.accounts.find({name:"jack"}).pretty()
```

"更新 jack 的银行账户余额和开户信息"(新增开户信息字段)

```sh
db.accounts.update(
  { name: "jack" },
  { $set:  
    {
      balance: 3000,
      info: {
        dateOpened: new Date("2016-05-18T16:00:00Z"),
        branch: "branch1"
      }
    } 
  }
)

db.accounts.find({name:"jack"}).pretty()
```

"更新 jack 的银行账户的开户时间"
```sh
db.accounts.update(
  { name: "jack" },
  { $set: 
    {
      "info.dateOpened": new Date("2017-01-01T16:00:00Z"),
    } 
  }
)
db.accounts.find({name:"jack"}).pretty()
```

"更新 jack 的联系电话"
```sh
db.accounts.update(
  { name: "jack" },
  { $set: 
    {
      "contact.0": "66666666", # 通过下标来操作
    } 
  }
)
db.accounts.find({name:"jack"}).pretty()
```

"添加 jack 的联系方式"(第四个元素)
```sh
db.accounts.update(
  { name: "jack" },
  { $set: 
    {
      "contact.3": "new contact",
    } 
  }
)
db.accounts.find({name:"jack"}).pretty()
```

"再次添加 jack 的联系方式"(第6个元素)
```sh
db.accounts.update(
  { name: "jack" },
  { $set: 
    {
      "contact.5": "new contact",
    } 
  }
)
db.accounts.find({name:"jack"}).pretty()
```
注意：第5个元素将被设置为 null

删除字段

"删除 jack 的银行账户余额和开户地点"
```sh
db.accounts.update(
  { name: "jack" },
  {
    $unset: {
      balance: "",
      "info.branch":"",
    }
  }
)
db.accounts.find({name:"jack"}).pretty()
```
其实$unset命令中的赋值("")对操作结果并没有任何影响

"删除 jack 的银行开户时间"
```sh
db.accounts.update(
  { name: "jack" },
  {
    $unset: {
      "info.dateOpened":"this can be any value",
    }
  }
)
db.accounts.find({name:"jack"}).pretty()
```
如果$unset命令中的字段根本不存在，那么文档内容将保持不变
```sh
db.accounts.update(
  { name: "jack" },
  {
    $unset: {
      notExist: ""
    }
  }
)
db.accounts.find({name:"jack"}).pretty()
```
删除数组内的字段

"删除 jack 的联系电话"(长度一样，元素的值被重置为 null)
```sh
db.accounts.update(
  { name: "jack" },
  { $unset:
      {
        "contact.0": ""
      } 
  }
)
db.accounts.find({name:"jack"}).pretty()
```
如果$rename命令要重命名的字段并不存在，那么文档内容不会被改变
```sh
db.accounts.update(
  { name: "jack" },
  { $rename:
      {
        "notExist": "name"
      } 
  }
)
```
如果新的字段名已经存在，那么原有的这个字段会被覆盖
```sh
db.accounts.find({name:"jack"}).pretty()
db.accounts.update(
  { name: "jack" },
  { $rename:
      {
        "name": "contact"
      } 
  }
)
db.accounts.find({contact:"jack"}).pretty()
```
当 $rename 命令中的新字段存在的时候，$rename 命令会先 $unset 新旧字段，然后$set新字段

重命名内嵌文档的字段

"更新 karen 的银行账户的开户时间和联系方式"
```sh
db.accounts.update(
  { name: "karen" },
  { $set: {
      info: {
        "dateOpened": new Date("2017-01-01T16:00:00Z"),
        branch: "branch1"
      },
      "contact.3": {
        primaryEmail: "xxx@gmail.com",
        secondaryEmail: "yyy@gmail.com"
      }
    } 
  }
)
db.accounts.find({name:"karen"}).pretty()
```

"更新账户余额和开户地点字段在文档中的位置"
```sh
db.accounts.update(
  { name: "karen" },
  { $rename:
      {
        "info.branch": "branch",
        "balance": "info.balance"
      } 
  }
)
db.accounts.find({name:"karen"}).pretty()
```
重命名数组中内嵌文档的字段

"更新karen的联系方式"
```sh
db.accounts.update(
  { name: "karen" },
  { $rename:
      {
        "contact.3.primaryEmail": "primaryEmail"
      } 
  }
)
# 这里会报错
db.accounts.find({name:"karen"}).pretty()
```
$rename 命令中的旧字段和新字段都不可以指向数组元素

这一点和之前介绍过的 $set 和 $unset 命令不同

$set 和 $unset 命令都可以应用在数组元素上

更新字段值

“更新 david 的账户余额”
```sh
db.accounts.find({name:"david"}).pretty()
db.accounts.update(
  { name: "david" },
  { $inc:
      {
        balance: -0.5
      } 
  }
)
db.accounts.find({name:"david"}).pretty()
db.accounts.update(
  { name: "david" },
  { $mul:
      {
        balance: 0.5
      } 
  }
)
db.accounts.find({name:"david"}).pretty()
```
$inc 和 $mul 只能应用在数字字段上

如果被更新的字段不存在，$inc 会创建字段，并且将字段值设为命令中的增减值，而 $mu1 会创建字段，但是把字段值设为0

"更新karen的账户余额"
```sh
db.accounts.find({name:"karen"}, { name:1,info:1, _id:0 }).pretty()
db.accounts.update(
  { name: "karen" },
  { $min:
      {
        "info.balance": 5000
      } 
  }
)
db.accounts.update(
  { name: "karen" },
  { $max:
      {
        "info.balance": 5000
      } 
  }
)
db.accounts.update(
  { name: "karen" },
  { $min:
      {
        "info.dateOpened": ISODate("2013-01-01T16:00:00Z")
      } 
  }
)
```
如果被更新的字段不存在……
```sh
db.accounts.update(
  { name: "karen" },
  { $min:
      {
        notYetExist: 10
      } 
  }
)
db.accounts.find({name:"karen"}, { name:1,info:1, _id:0 }).pretty()
```
$min和$max命令会创建字段，并且将字段设为命令中的更新值

如果被更新的字段类型和更新值类型不一致……
```sh
db.accounts.update(
  { name: "karen" },
  { $min:
      {
        "info.balance":null
      } 
  }
)
```
如果被更新的字段类型和更新值类型不一致，$min 和 $max 命令会按照 BSON 数据类型排序规则进行比较

* Null (最小)
* Numbers(ints, longs, doubles, decimals)
* Symbol, String
* Object
* Array
* BinData
* ObjectId
* Boolean
* Date
* Timestamp
* Regular Expression (最大)

数组更新操作符

* $addToSet
* $pop
* $pull
* $pullAll
* $push

"查看 karen 的银行账户文档"
```sh
db.accounts.find({name:"karen"},{name:1, contact: 1, _id:0}).pretty()
```

"向 karen 的账户文档中添加联系方式"
```sh
db.accounts.update(
  { name: "karen" },
  { $addToSet: { contact: "China" } }
)
```

如果要插入的值已经存在数组字段中，则 $addToSet 不会再添加重复(完全匹配包括数组中的顺序)值

"向 karen 的账户文档中添加新的联系方式"
```sh
db.accounts.update(
  { name: "karen" },
  { 
    $addToSet: { 
      contact: {
        "secondaryEmail" : "yyy@gmail.com",
        "primaryEmail" : "xxx@gmail.com"
      }
    } 
  }
)
```

"向 karen 的账户文档中添加多个联系方式"

```sh
db.accounts.update(
  {name: "karen" },
  { $addToSet: { contact: [ "contact1", "contact2" ] } }
)
```

$addToSet 会将数组插入被更新的数组字段中，成为内嵌数组

如果想要将多个元素直接添加到数组字段中，则需要使用$each操作符

```sh
db.accounts.update(
  {name: "karen" },
  { $addToSet: { contact: { $each: [ "contact1", "contact2" ] } } }
)
db.accounts.find({name:"karen"},{name:1, contact: 1, _id:0}).pretty()
```

"从 karen 的账户文档中删除最后一个联系方式"
```sh
db.accounts.update(
  { name: "karen" },
  { $pop: { contact: 1 } }
)
```

"从 karen 的账户文档中删除联系方式"(内嵌数组操作)
```sh
db.accounts.update(
  { name: "karen" },
  { $pop: { "contact.5": -1 } }
)
```

"继续从 karen 的账户文档中删除联系方式"
```sh
db.accounts.update(
  { name: "karen" },
  { $pop: { "contact.5": -1 } }
)
```
删除掉数组中的最后一个元素后，会留下空数组

注意一点，$pop操作符只能应用在数组字段上
```sh
db.accounts.update(
  { name: "karen" },
  { $pop: { "contact.1": -1 } }
)
```

从数组字段中删除特定元素

"将 karen 的账户文档复制为 lawrence 的账户文档"
```sh
db.accounts.find(
  { name: "karen" },
  { _id: 0 }
).forEach( function(doc) {
    var newDoc = doc;
    newDoc.name = "lawrence";
    db.accounts.insert(newDoc)
  }
)
db.accounts.find({name: "lawrence"}).pretty()
```

"从 karen 的联系方式中删去包含 'hi' 字母的元素"
```sh
db.accounts.update(
  { name: "karen" },
  { $pull: { contact: { $regex: /hi/ } } }
)
```
"从 karen 的联系方式中删去电话号码 22222222"

数组里面的整个内嵌数组全被干掉
```sh
db.accounts.update(
  { name: "karen" },
  { $pull: {
      contact: {
        $elemMatch: { $eq: "22222222" }
      }
    } 
  }
)
```
$pullAll 命令只会删去字段和字段排列顺序都完全匹配的文档元素

$pull 命令会删去包含指定的文档字段和字段值的文档元素，字段排列顺序不需要完全匹配
```sh
db.accounts.update(
  { name: "lawrence" },
  { 
    $pull: {
      contact: { "primaryEmail" : "xxx@gmail.com"}
    } 
  }
)
db.accounts.find({name: "lawrence"}).pretty()
```

向数组字段中添加元素

$push 和 $addToSet 命令相似，但是 $push 命令的功能更强大

和 $addToSet 命令一样，如果 $push 命令中指定的数组字段不存在，这个字段会被添加到原文档中
```sh
db.accounts.update(
  { name: "lawrence" },
  { $push: {
      newArray: "new element"
    } 
  }
)
db.accounts.find({name: "lawrence"}, { name:1, newArray:1, _id:0 }).pretty()
```
和 $addToSet 相似，$push 操作符也可以和 $each 搭配使用
```sh
db.accounts.update(
  { name: "lawrence" },
  { $push: {
      newArray: { $each: [ 2, 3, 4 ] }
    } 
  }
)
```
$push 和 $each 操作符还可以和更多的操作符搭配使用，实现比 $addToSet 更复杂的操作

使用 $position 操作符将元素插入到数组的指定位置
```sh
db.accounts.update(
  { name: "lawrence" },
  { $push: {
      newArray: {
        $each: [ "pos1", "pos2" ],
        $position: 0
      }
    } 
  }
)
db.accounts.find({ name: "lawrence" }).pretty()
db.accounts.update(
  { name: "lawrence" },
  { $push: {
      newArray: {
        $each: [ "pos3", "pos4" ],
        $position: -1
      }
    } 
  }
)
```

在上面这个例子中，$position: -1 和 $position: 5对应的位置是相同的

使用 $sort 对数组进行排序
```sh
db.accounts.update(
  { name: "lawrence" },
  { 
    $push: {
      newArray: {
        $each: [ "sort1" ],
        $sort: 1
      }
    } 
  }
)
```

如果插入的元素是内嵌文档，也可以根据内嵌文档的字段值排序
```sh
db.accounts.update(
  { name: "lawrence" },
  { 
    $push: {
      newArray: {
        $each:  [
          { key: "sort", value: 100 },
          { key: "sort", value: 200 }
        ],
        $sort: { value: -1 }
      },
    } 
  } 
)
```
如果不想插入元素，只想对文档中的数组字段进行排序……
```sh
db.accounts.update(
  { name: "lawrence" },
  {
    $push: {
      newArray: {
        $each: [],
        $sort: -1
      }
    }
  }
)
```

使用 $slice 来截取部分数组(到着数8个留下来)
```sh
db.accounts.update(
  { name: "lawrence" },
  { 
    $push: {
      newArray: {
        $each: [ "slice1" ],
        $slice: -8
      }
    } 
  }
)

db.accounts.find({ name: "lawrence" }).pretty()
```

如果不想插入元素，只想截取文档中的数组字段……
```sh
db.accounts.update(
  { name: "lawrence" },
  {
    $push: {
      newArray: {
        $each: [],
        $slice: 6
      }
    }
  }
)
```

向数组字段中添加元素

$position, $sort, $slice 可以一起使用

这三个操作符的执行顺序是：
* $position
* $sort
* $slice

写在命令中的操作符顺序并不重要，并不会影响命令的执行顺序
```sh
db.accounts.update(
  { name: "lawrence" },
  {
    $push: {
      newArray: {
        $each: [ "push1", "push2" ],
        $position: 2,
        $sort: -1,
        $slice: 5
      }
    }
  }
)
```

更新数组中的特定元素

$是数组中第一个符合筛选条件的数组元素的占位符

搭配更新操作符使用，可以对满足筛选条件的数组元素进行更新
```sh
db.accounts.find({ name: "lawrence" }, { name: 1, newArray: 1, _id: 0 }).pretty()

db.accounts.update(
  {
    name: "lawrence",
    newArray: "pos2"
  },
  {
    $set: {
      "newArray.$": "updated"
    }
  }
)
```
更新数组中的所有元素

搭配更新操作符使用，可以对数组中所有的元素进行更新

```sh
db.accounts.find(
  { name: "lawrence" },
  { name: 1, contact: 1, _id:0 }
).pretty()

db.accounts.update(
  { name: "lawrence" },
  {
    $set: { "contact.0.$[]": "88888888" }
  }
)
```

更新多个文档

到目前为止，我们在 update 命令中使用的筛选条件只对应于一篇文档

在默认情况下，即使筛选条件对应了多篇文档，update 命令仍然只会更新*一篇*文档
```sh
# 更新的始终只有一篇

db.accounts.update(
  {},
  {
    $set: { currency: "USD" }
  }
)
```

使用 multi 选项来更新多个符合筛选条件的文档
```sh
db.accounts.update(
  {},
  {
    $set: { currency: "USD" }
  },
  { multi: true }
)
```
注意，MongoDB 只能保证*单个*文档操作的原子性，不能保证*多个*文档操作的原子性

更新多个文档的操作虽然在单一线程中执行，但是线程在执行过程中可能被挂起，以便其他线程也有机会对数据进行操作

如果需要保证多个文档操作时的原子性，就需要使用 MongoDB 4.0 版本引入的实物功能进行操作

在默认情况下，如果 update 命令中的筛选条件没有匹配任何文档，则不会进行任何操作

将 upsert 选项设置为 true，如果 update 命令中的筛选条件没有匹配任何文档，则会创建新文档

“查看 maggie 的银行账户文档”
```sh
db.accounts.find({name: "maggie"}, {name: 1, balance: 1, _id: 0})

db.accounts.update(
  { name: "maggie" },
  {
    $set: { balance: 700}
  },
  { upsert: true }
)
```

不过，如果无法从筛选条件中推断出确定的字段值，新创建的文档将不会包含筛选条件涉及的字段
```sh
db.accounts.update(
  { balance: { $gt: 20000 } },
  { $set: { name: "nick" } },
  { upsert: true }
)
```
"查看 nick 的银行账户文档"
```sh
db.accounts.find({ name: "nick" }, {_id: 0})
{ "name" : "nick" } ### 只有一个字段
```

如果 document 文档中包含了 _id 字段，save() 命令将会调用 db.collection.update() 命令(upsert: true)
```sh
db.accounts.save(
  { _id: "account1", name: "alice", balance: 100 }
)
db.accounts.find({name: "alice"})
db.accounts.save(
  { _id: "account2", name: "oliver", balance: 100 }
)
db.accounts.find({name: "oliver"})
```

总结一下
* 使用 db.collection.update() 命令更新整篇文档
* 使用 db.collection.update() 命令更新文档中的特定字段
* 文档更新操作符
* 数组更新操作符
* 使用 db.collection.update() 命令更新多篇文档
* 使用 db.collection.update() 命令更新或者创建文档
* 使用 db.collection.save() 命令更新文档

#### 删除文档
* 删除集合
* 删除特定文档
* db.collection.remove()

```sh
# 查看银行账户文档
db.accounts.find(
  {},
  { name:1, balance: 1, _id: 0 }
).sort( { balance: 1 } )

# 删除余额为50的银行账户文档
db.accounts.remove({balance: 50})
```
在默认情况下，remove 命令会删除所有符合筛选条件的文档

如果只想删除满足筛选条件的*第一篇*文档，可以使用 justOne 选项

"删除一篇余额小于100的银行账户文档"

```sh
db.accounts.remove(
  { balance: { $lt: 100 } },
  { justOne: true }
)
```
"删除集合内的所有文档"
```sh
db.accounts.remove({})
```
"删除集合"

drop 命令可以删除整个集合，包括集合中的所有文档，以及集合的索引

```sh
show collections
db.accounts.drop()
show collections
```

如果集合中的文档数量很多，使用 remove 命令删除所有文档的效率不高。这种情况下，更加有效率的方法，是使用 drop 命令删除集合，然后再创建空集合并创建索引

总结一下
* 使用 db.collection.remove() 命令删除文档
* 使用 db.collection.drop() 命令删除集合

### 聚合操作

数据分析

* 单一用途的聚合方法
* Map Reduce
* 聚合管道 db.collection.aggregate()

聚合表达式

* 用来操作输入文档的 "公式"
* 经聚合表达式计算出的值可以被赋予输出文档中的字段
* 字段路径，系统变量，文本，表达式对象，操作符

聚合阶段

* 聚合阶段有顺序地排列在聚合管道中
* 绝大多数聚合阶段可以反复出现（$out 和 $geoNear 除外）
* 数据库层面和集合层面

聚合操作符

* 用来构建聚合表达式

#### 聚合表达式

介绍几种常见的表达式

字段路径表达式
* $<field> - 使用 $ 来指示字段路径
* $<field>.<sub-field> - 使用 $ 和 . 来指示内嵌文档字段路径
* $name - 指示银行账户文档中客户姓名的字段
* $info.dateOpened - 指示银行账户文档中开户日期的字段

系统变量表达式
* $$<variable> - 使用$$来指示系统变量
* $$CURRENT - 指示管道中当前操作的文档
  * - $$CURRENT.<field>和$<field>是等效的

常量表达式
* $literal: <value> - 指示常量<value>
* $literal: "$name" - 指示常量字符串"$name"
  * 这里的 $ 被当作常量处理，而不是字段路径表达式

#### 聚合管道阶段
* $project - 对输入文档进行再次投影
* $match - 对输入文档进行筛选
* $limit - 筛选出管道内前 N 篇文档
* $skip - 跳过管道内前 N 篇文档
* $unwind - 展开输入文档中的数组字段
* $sort - 对输入文档进行排序
* $lookup - 对输入文档进行查询操作
* $group - 对输入文档进行分组
* $out - 将管道中的文档输出

$poject

"先创建几个文档"
```sh
db.accounts.insertMany([
  {
    name: { firstName: "alice", lastName: "wong" },
    balance: 50
  },
  {
    name: { firstName: "bob", lastName: "yang" },
    balance: 20
  }
])
```
"对银行账户文档进行重新投影"
```sh
db.accounts.aggregate([
  {
    $project: {
      _id: 0,
      balance: 1,
      clientName: "$name.firstName"
    }
  }
])

db.accounts.aggregate([
  {
    $project: {
      _id: 0,
      balance: 1,
      clientName: "$name.firstName"
    }
  }
])

db.accounts.aggregate([
  {
    $project: {
      _id: 0,
      balance: 1,
      nameArray: [
        "$name.firstName",
        "$name.middleName",
        "$name.lastName"
      ]
    }
  }
])
```

$project 是一个很常用的聚合阶段

可以用来灵活地控制输出文档的格式

也可以用来剔除不相关的字段，以优化聚合管道操作的性能

$match

$match 中使用的文档筛选语法和读取文档是的筛选语法相同

"对银行账户文档进行筛选"
```sh
db.accounts.aggregate([
  {
    $match: {
      "name.firstName": "alice"
    }
  }
])

db.accounts.aggregate([
  {
    $match: {
      $or: [
        { balance: { $gt: 40, $lt: 80 } },
        { "name.lastName": "yang" }
      ]
    }
  }
])
```

"将筛选和投影阶段结合在一起"
```sh
db.accounts.aggregate([
  {
    $match: {
      $or: [
        { balance: { $gt: 40, $lt: 80 } },
        { "name.lastName": "yang" }
      ]
    }
  },
  {
    $project: {
      _id: 0
    }
  }
])
```

$match 也是一个很常用的聚合阶段

应该尽量在聚合管道的开始阶段应用$match

这样可以减少后续阶段中需要处理的文档数量，优化聚合操作的性能

$limit & $skip

"筛选第一篇银行账户文档"
```sh
db.accounts.aggregate([
  { $limit: 1 }
])
```

"跳过第一篇银行账户文档"
```sh
db.accounts.aggregate([
  { $skip: 1 }
])
```
$unwind

"向现有的银行账户文档中加入数组字段"
```sh
db.accounts.update(
  { "name.firstName": "alice" },
  { 
    $set: {
      currency: [ "CNY", "USD" ]
    }
  }
)
db.accounts.update(
  { "name.firstName": "bob" },
  { 
    $set: { currency: "GBP" }
  }
)
```

"将文档中的货币种类数组展开"
```sh
db.accounts.aggregate([
  {
    $unwind: {
      path: "$currency"
    }
  }
])
```
"展开数组时添加元素位置"
```sh
db.accounts.aggregate([
  {
    $unwind: {
      path: "$currency",
      includeArrayIndex: "ccyIndex"
    }
  }
])
```

"再添加几个文档"
```sh
db.accounts.insertMany([
  {
    name: { firstName: "charlie", lastName: "gordon" },
    balance: 100
  },
  {
    name: { firstName: "david", lastName: "wu" },
    balance: 200,
    currency: []
  },
  {
    name: { firstName: "charlie", lastName: "kim" },
    balance: 20,
    currency: null
  }
])
```
"将文档中的货币种类数组展开"
```sh
db.accounts.aggregate([
  {
    $unwind: {
      path: "$currency"
    }
  }
])
```
发现上面的文档剔除掉了

"展开数组时保留空数组或不存在数组的文档"
```sh
db.accounts.aggregate([
  {
    $unwind: {
      path: "$currency",
      preserveNullAndEmptyArrays: true
    }
  }
])
```

"对银行账户文档进行排序"
```sh
db.accounts.aggregate([
  { 
    $sort: {
      balance: 1,
      "name.lastName": -1
    } 
  }
])
```

$lookup

使用单一字段值进行查询

"增加一个集合用来储存外汇数据"
```sh
db.forex.insertMany([
  { ccy: "USD", rate: 6.91, date: new Date("2018-12-21") },
  { ccy: "GBP", rate: 8.72, date: new Date("2018-08-21") },
  { ccy: "CNY", rate: 1.0, date: new Date("2018-12-21") }
])
```
"将查询到的外汇汇率写入银行账户文档"
```sh
db.accounts.aggregate([
  {
    $lookup: {
      from: "forex",
      localField: "currency",
      foreignField: "ccy",
      as: "forexData"
    }
  }
])
```
如果 localField 是一个数组字段……
```sh
db.accounts.aggregate([
  {
    $unwind: {
      path: "$currency"
    }
  },
  {
    $lookup: {
      from: "forex",
      localField: "currency",
      foreignField: "ccy",
      as: "forexData"
    }
  }
])
```
使用复杂条件进行查询

对查询集合中的文档使用聚合阶段进行处理时，如果需要参考管道文档中的字段

则必须使用 let 参数对字段进行声明

"将特定日期外汇汇率写入银行账户文档"
```sh
db.accounts.aggregate([
  {
    $lookup: {
      from: "forex",
      pipeline: [
        {
          $match: {
            date: new Date("2018-12-21")
          }
        }
      ],
      as: "forexData"
    }
  }
])
```

注意，在这个例子中，查询条件和管道文档之间，其实并没有直接的联系

这种查询被称作不相关查询，$lookup 从 3.6版本开始支持不相关查询

"将特定日期外汇汇率写入余额大于 100 的银行账户文档"
```sh
db.accounts.aggregate([
  {
    $lookup: {
      from: "forex",
      let: { bal: "$balance" },
      pipeline: [
        {
          $match: {
            $expr: {
              $and: [
                { $eq: [ "$date", new Date("2018-12-21") ] },
                { $gt: [ "$$bal", 100 ] }
              ]
            }
          }
        }
      ],
      as: "forexData"
    }
  }
])
```

$group

定义分组规则

可以使用聚合操作符来定义新字段

"增加一个集合用来储存股票交易记录"
```sh
db.transactions.insertMany([
  {
    symbol: "600519",
    qty: 100,
    price: 567.4,
    currency: "CNY"
  },
  {
    symbol: "AMZN",
    qty: 1,
    price: 1377.5,
    currency: "USD"
  },
  {
    symbol: "AAPL",
    qty: 2,
    price: 150.7,
    currency: "USD"
  }
])
```

"按照交易货币来分组交易记录"
```sh
db.transactions.aggregate([
  {
    $group: {
      _id: "$currency"
    }
  }
])
```

不使用聚合操作符的情况下，$group可以返回管道文档中某一字段的所有（不重复的）值
```sh
db.transactions.aggregate([
  {
    $group: {
      _id: "$currency",
      totalQty: { $sum: "$qty" },
      totalNotional: { $sum: { $multiply: [ "$price", "$qty" ] } },
      avgPrice: { $avg: "$price" },
      count: { $sum: 1 },
      maxNotional: { $max: { $multiply: [ "$price", "$qty" ] } },
      minNotional: { $min: { $multiply: [ "$price", "$qty" ] } }
    }
  }
])
```

"使用聚合操作符计算所有文档聚合值"(没分组，但是进行了聚合运算)
```sh
db.transactions.aggregate([
  {
    $group: {
      _id: null,
      totalQty: { $sum: "$qty" },
      totalNotional: { $sum: { $multiply: [ "$price", "$qty" ] } },
      avgPrice: { $avg: "$price" },
      count: { $sum: 1 },
      maxNotional: { $max: { $multiply: [ "$price", "$qty" ] } },
      minNotional: { $min: { $multiply: [ "$price", "$qty" ] } }
    }
  }
])
```
"使用聚合操作符创建数组字段"
```sh
db.transactions.aggregate([
  {
    $group: {
      _id: "$currency",
      symbols: { $push: "$symbol" }
    }
  }
])
```

$out 

"将聚合管道中的文档写入一个新集合"
```sh
db.transactions.aggregate([
  {
    $group: {
      _id: "$currency",
      symbols: { $push: "$symbol" }
    }
  },
  {
    $out: "output"
  }
])
```

"查看 output 集合"
```sh
db.output.find()
```

"将聚合管道中的文档写入一个已存在的集合"

会覆盖已存在集合的内容
```sh
db.transactions.aggregate([
  {
    $group: {
      _id: "$symbol",
      totalNotional: { 
        $sum: {
          $multiply: [ "$price", "$qty" ]
        }
      }
    }
  },
  {
    $out: "output"
  }
])
```

"查看 output 集合"
```sh
db.output.find()
```

如果聚合管道操作遇到错误，管道阶段不会创建新集合或是覆盖已存在的集合内容

allowDiskUse

每个聚合管道阶段使用的内存不能超过 100MB 

如果数据量较大，为了防止聚合管道阶段超出内存上限并且抛出错误，可以启用allowDiskUse 选项

allowDisUse 启用之后，聚合阶段可以在内存容量不足时，将操作数据写入临时文件中，临时文件会被写入 dbPath 下的 _tmp 文件夹，dbPath的默认值为 /data/db

```sh
db.transactions.aggregate([
  {
    $group: {
      _id: "$currency",
      symbols: { $push: "$symbol" }
    }
  }],
  { allowDiskUse: true }
)
```

#### 聚合操作的优化

MongoDB 内部的一些优化

聚合阶段顺序优化

$project + $match
* $match 阶段会在 $project 阶段之前运行

$sort + $match
* $match 阶段会在 $sort 阶段之前运行

$project + $skip
* $skip 阶段会在 $project 阶段之前运行

聚合阶段合并优化

$sort + $limit

如果两者之间没有夹杂着会改变文档数量的聚合阶段，$sort和$limit阶段可以合并

$limit + $limit
$skip + $skip
$match + $match

连续的 $limit，$skip 或 $match 阶段排列在一起时，可以合并为一个阶段

$lookup + $unwind

连续排列在一起的 $lookup 和 $unwind 阶段，如果 $unwind 应用在 $lookup 阶段创建的 as 字段上，则两者合并

总结一下

使用 db.collection.aggregate() 命令进行聚合操作

使用聚合表达式

常用的聚合管道阶段

常用的聚合操作符

聚合操作的局限和优化

### 索引

* Index
* 合适的索引可以大大提升数据库搜索性能
* 集合层面的索引

对指定字段进行排序的数据结构(B-tree)

  * 更快的查询
  * 更快的排序

复合键索引可以对多个字段进行排序

复合键索引只能支持前缀子查询

* {A}       ~~{B}~~
* {A, B}    ~~{C}~~
* {A, B, C} ~~{B,C}~~

索引

* 对文档部分内容进行排序的数据结构
* 加快文档查询和文档排序的速度
* 复合键索引只能支持前缀子查询

索引操作

* db.collection.getIndexes()
* db.collection.createIndex()
* db.collection.dropIndex()

索引的类型

* 单键索引
* 复合键索引
* 多键索引

索引的特性

* 唯一性
* 稀疏性
* 生存时间

查询分析

* 检视索引的效果
* explain()

索引的选择

* 如何创建一个合适的索引
* 索引对数据库的写入操作的影响

#### 创建索引

db.collection.createIndex()

"创建一个新集合"

```sh
db.accountsWithIndex.insertMany([
  { name: "alice", balance: 50, currency: ["GBP", "USD"] },
  { name: "bob", balance: 20, currency: ["AUD", "USD"] },
  { name: "bob", balance: 300, currency: ["CNY"] },
])
```

"创建一个单键索引"

```sh
db.accountsWithIndex.createIndex({name: 1})
```

"列出集合中已存在的索引"

```sh
db.accountsWithIndex.getIndexes()
```

"创建一个复合键索引"

只支持前缀子查询

```sh
db.accountsWithIndex.createIndex({ name:1, balance: -1 })

db.accountsWithIndex.getIndexes()
```

"创建一个多键索引"（针对数组）
```sh
db.accountsWithIndex.createIndex({ currency: 1 })
```
数组字段中的每一个元素，都会在多键索引中创建一个键

* "AUD" --> {"bob"}
* "CNY" --> {"bob"}
* "GBP" --> {"alice"}
* "USD" --> {"alice"}
* "USD" --> {"bob"}

#### 索引的效果

db.collection.explain()

可以使用 explain 进行分析的命令包括 aggregate(), count(), distinct(), find(), group(), remove(), update()

"使用没有创建索引的字段进行搜索"
```sh
db.accountsWithIndex.explain().find({ balance: 100 })
# COLLSCAN -> 效果最差的搜索方式
```

"使用已经创建索引的字段进行搜索"
```sh
db.accountsWithIndex.explain().find({ name: "alice" })
# IXSCAN
```

"仅仅返回创建了索引的字段"
```sh
db.accountsWithIndex.explain().find({ name: "alice" }, { _id: 0, name: 1})
# FETCH 都不需要了，最大程度提升了查询效果
```

"使用已经创建索引的字段进行排序"
```sh
db.accountsWithIndex.explain().find().sort({ name: 1, balance: -1 })
```

"使用未创建索引的字段进行排序"
```sh
db.accountsWithIndex.explain().find().sort({ name: 1, balance: 1 })
# SORT -> 提示排序效率并不高
```

删除索引

db.collection.dropIndex()

如果需要更改某些字段上已经创建的索引

必须首先删除原有索引，再重新创建新的索引


