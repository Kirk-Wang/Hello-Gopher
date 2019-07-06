# 常见故障诊断

响应事件增长
* 对于一般的 web 服务来说，响应时间应该在 200ms 以内
* 对于一般的 mongodb 请求来说，响应时间应该在 100ms 以内

合适的索引
* 使用 explain() 查看索引的有效性

工作集超出 RAM 的大小
* 使用 mongostat 查看服务器状态

模拟缓存不足的情镜

mongo.conf
```sh
storage:
  dbPath: /data/db
  wiredTiger:
    engineConfig:
      cacheSizeGB: 0.25
```

docker-compose.yml

```sh
command: mongod -f /data/db/mongo.conf
```

导入大量数据
```sh
docker exec -it fe1bb8f3934b bash
/data/db/scripts/load-large-dataset.sh
```

使用 mongostat 监控服务器进程状态
```sh
mongostat --host localhost --port 27017 -o "command,dirty,used,vsize,res,conn,time"

# 如果 used 远远大于 dirty ,就要看看工作集大小了
```

连接失败

默认情况下，mongod 进程可以支持多达 65536 个连接

不恰当的配置可能限制连接数

查看支持的连接数
```sh
docker exec -it ffb49ad941e9 mongo
> db.serverStatus().connections
{ "current" : 5, "available" : 838855, "totalCreated" : 6 }
```

查看数据库服务器配置文件（如：mongo.conf）
```sh
net:
  maxIncomingConnections: 200
storage:
  dbPath: /data/db
  wiredTiger:
    engineConfig:
      cacheSizeGB: 0.25
```

第二种情况：查看 ulimit 配置
```sh
docker exec -it ffb49ad941e9 mongo

ulimit -a
#open files -> 提升mongo.conf已无用，可提升它
```