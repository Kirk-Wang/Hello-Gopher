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