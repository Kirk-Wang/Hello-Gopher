
# [Introduction to MongoDB](https://docs.mongodb.com/manual/introduction/)

### 使用 Docker 一秒本地搭建 Mongodb  & mongo-express 环境

编辑 docker-compose.yml
```sh
vim docker-compose.yml
```
```yml
version: '3.1'

services:

  mongo:
    image: mongo:4.0.6
    restart: always
    volumes:
      - ./data:/data/db
    environment:
      MONGO_INITDB_ROOT_USERNAME: root
      MONGO_INITDB_ROOT_PASSWORD: 123456

  mongo-express:
    image: mongo-express
    restart: always
    ports:
      - 8081:8081
    environment:
      ME_CONFIG_MONGODB_ADMINUSERNAME: root
      ME_CONFIG_MONGODB_ADMINPASSWORD: 123456
```

启动
```sh
docker-compose up -d
```

进入 mongo-express，[http://localhost:8081](http://localhost:8081)，对 database 进行一系列的操作

