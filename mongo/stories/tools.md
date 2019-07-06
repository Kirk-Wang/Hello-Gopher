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
docker exec -it a811efa08b1d bash

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

导出 json 文件
```sh
mongoexport --db test --collection accounts --type=json --fields name.firstName,name.lastName,balance --out opt/backups/accounts.json -u readUser -p passwd --authenticationDatabase admin
```

查看导出文件
```sh
cat opt/backups/accounts.json
```


导出 json 文件时，--fields选项是可选的
```sh
mongoexport --db test --collection accounts --type=json --out opt/backups/accounts.json -u readUser -p passwd --authenticationDatabase admin
```

查看导出文件
```sh
cat opt/backups/accounts.json
```

使用查询语句筛选导出文档
```sh
mongoexport --db test --collection accounts --type=json --fields name.firstName,name.lastName,balance --out opt/backups/accounts.json -u readUser -p passwd --authenticationDatabase admin --query '{balance:{$gte: 100}}'

cat opt/backups/accounts.json
```

使用 --host, --port 选项
```sh
mongoexport --db test --collection accounts --type=json --out opt/backups/accounts.json -u readUser -p passwd --authenticationDatabase admin --host localhost --port 27017

cat opt/backups/accounts.json
```

使用 --limit, --skip, --sort 选项
```sh
mongoexport --db test --collection accounts --type=json --fields name.firstName,name.lastName,balance --out opt/backups/accounts.json -u readUser -p passwd --authenticationDatabase admin --sort '{balance:1}' --limit 3 --skip 1

cat opt/backups/accounts.json
```

mongoimport

将数据由 json 或 csv 格式文件导入

需要对操作的数据库具备 readWrite 权限
```sh
docker exec -it a811efa08b1d mongo -u myUserAdmin -p passwd --authenticationDatabase admin
use admin
db.createUser(
  {
    user: "writeUser",
    pwd: "passwd",
    roles: [ "readWriteAnyDatabase" ]
  }
)
exit
# 登录
docker exec -it a811efa08b1d bash
```

查看导入文件
```sh
cat opt/backups/accounts.csv
```

导入csv文件
```sh
# headerline,告诉mongodb第一行不是数据
mongoimport --db test --collection importAccounts --type csv --headerline --file /opt/backups/accounts.csv -u writeUser -p passwd --authenticationDatabase admin
```

查看导入文档
```sh
mongo -u readUser -p passwd --authenticationDatabase admin --quiet --eval 'db.importAccounts.find()'
```