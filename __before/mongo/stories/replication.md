# 复制集

* 高可用性
* 数据安全
* 分流/分工

### 复制集节点

* 主节点负责处理所有的写入请求
* 主节点（默认）和副节点都可以处理读取请求
* 副节点从主节点（或者符合条件的副节点）出复制数据
* 每个节点都会向其他节点发送心跳请求
* 每隔2秒发送一次，超过10秒则请求超时（默认）
* 复制集中最多可以有50个节点

### 复制集选举

* 候选节点发起选举，每个节点投票给比自己更同步的节点
* 得到超过半数选票的候选节点会当选为主节点
* 复制集中最多可以有7个投票节点

### 触发选举的事件

* 主节点与副节点之间的心跳请求超时
* 复制集初始化
* 新节点加入复制集

### 投票机节点

* 没有数据
* 可以投票
* 不能成为主节点

### 初始同步

* 数据库
* 集合
* 索引
* 文档

### 写库纪录同步

主节点
  * 文档A：新增
  * 文档B：更新1
  * 文档A：更新1

副节点（作同样操作）
  * local.oplog.rs

### 写库纪录

* 写库日志中的纪录可以被重复使用
* 多个线程分批次使用日志纪录
* 写库日志的大小和文档大小不一定成正比

### 创建复制集

* 创建 docker network
```sh
docker network create web
docker network ls
```

docker-compose.yml

```yml
version: '3.1'

services:

  mongo:
    image: mongo:4.0.6
    restart: always
    networks:
      - web
    ports:
      - 27017:27017
    command: mongod --replSet rs
    volumes:
      - ./data:/data/db

  mongo2:
    image: mongo:4.0.6
    restart: always
    networks:
      - web
    ports:
      - 27018:27017
    command: mongod --replSet rs
    volumes:
      - ./data2:/data/db
  
  mongo3:
    image: mongo:4.0.6
    restart: always
    networks:
      - web
    ports:
      - 27019:27017
    command: mongod --replSet rs
    volumes:
      - ./data3:/data/db

  mongo-express:
    image: mongo-express
    restart: always
    networks:
      - web
    ports:
      - 8081:8081

networks:
  web:
    external: true
```

```sh
docker exec -it f202f6ae60f4 mongo

rs.initiate(
  {
    _id: "rs",
    members: [
      { _id: 0, host: "mongo:27017" },
      { _id: 1, host: "mongo2:27017" },
      { _id: 2, host: "mongo3:27017" },
    ]
  }
)
```

