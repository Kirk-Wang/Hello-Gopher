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
# readAnyDatabase 只在 admin 数据库中才提供
```

导出 csv 文件
```sh
mongoexport --db test --collection accounts --type=csv --fields name,balance --out opt/backups/accounts.csv -u readUser -p passwd --authenticationDatabase admin
```

查看导出文件
```sh
cat opt/backups/accounts.csv
```

导出内嵌文档字段
```sh
mongoexport --db test --collection accounts --type=csv --fields name.firstName,name.lastName,balance --out opt/backups/accounts.csv -u readUser -p passwd --authenticationDatabase admin
```

查看导出文件
```sh
cat opt/backups/accounts.csv
```

