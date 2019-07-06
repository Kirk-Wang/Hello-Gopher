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