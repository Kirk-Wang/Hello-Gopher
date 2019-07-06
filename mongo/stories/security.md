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
```

