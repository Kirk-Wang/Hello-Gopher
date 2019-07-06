# 数据处理工具

* mongoexport
* mongoimport

### mongoexport
* 将数据导出为 json 或 csv 格式文件
* 需要对操作的数据库具备 read 权限

"创建执行 mongoexport 的用户"
```sh
docker exec -it a811efa08b1d mongo -u myUserAdmin -p passwd --authenticationDatabase admin
use admin
db.createUser(
  {
    user: "readUser",
    pwd: "passwd",
    roles: [ "readAnyDatabase" ]
  }
)
```


