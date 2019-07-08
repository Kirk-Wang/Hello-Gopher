# 事务

### ACID 特性

* Atomicity 原子性
* Consistency 一致性
* Isolation 隔离性
* Durability 持久性

#### Atomicity 原子性

由A向转账50元（两次操作，不可分割）
* A(50)-50
* B(50)+50
* ----
* A:0
* B:100
  * 代表转账成功
