### MongoDB 是什么？

存储 `文档` 的 `非关系型` 数据库

数据库--》集合--》文档

#### 一行命令在 Docker 中运行 MongoDB
```sh
docker pull mongo:4 # 下载 MongoDB 的官方 docker 镜像
docker images # 查看下载的镜像

docker run --name mymongo -v /mymongo/data:/data/db -d mongo:4
# --name mymongo --> 容器名字
# -v /mymongo/data:/data/db --> 挂载数据目录
# -d --> 后台运行容器

docker ps # 查看 docker 容器状态

docker logs mymongo # 查看数据库服务器日志
```


