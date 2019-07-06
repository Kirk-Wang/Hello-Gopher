# 数据库安全

* 认证
* 验证身份

* 授权
* 操作权限

创建第一个用户
```sh
use admin
db.createUser({
  user: "myUserAdmin",
  pwd: "passwd",
  roles: [ "userAdminAnyDatabase" ]
})
```

启用身份认证
```sh
docker-compose down
# command: mongod --auth 加入 docker-compose.yml
docker-compose up -d
```

使用用户名和密码进行身份验证
```sh
# authenticationDatabase
# 对应的验证数据库
mongo -u "myUserAdmin" -p "passwd" --authenticationDatabase "admin"

docker exec -it a811efa08b1d mongo -u myUserAdmin -p passwd --authenticationDatabase admin

> use test
> db.accounts.find()
Error: error: {
        "ok" : 0,
        "errmsg" : "not authorized on test to execute command { find: \"accounts\", filter: {}, lsid: { id: UUID(\"4065e318-7334-426d-a7e0-2a31e210bf23\") }, $db: \"test\" }",
        "code" : 13,
        "codeName" : "Unauthorized"
# 有管理的权限，但并没有读取的权限
```

使用 db.auth() 进行身份验证
```sh
docker exec -it a811efa08b1d mongo
> db
test
> use admin;
switched to db admin
> db.auth("myUserAdmin","passwd")
1
```

授权

权限
* = 在哪里 + 做什么
```sh
{
  resource: {
    db: "test",
    collection: ""
  },
  actions: [
    "find",
    "update"
  ]
}
# 在test数库，你可以进行 find & update

{
  resource: {
    cluster: true
  },
  actions: [
    "shutdown"
  ]
}
# 可以停止整个集群
```

角色
* 角色 = 一组权限的集合
* read - 读取当前数据库中所有非系统集合
* readWrite - 读写当前数据库中所有非系统集合 