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

