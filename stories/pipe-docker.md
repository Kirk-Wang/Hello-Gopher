#### 利用 Docker，先让它快速跑起来

下载 [SQLite](https://www.sqlite.org/download.html)，我的是 Mac，大家对应着下载。

[SQLite Database Browser](https://github.com/sqlitebrowser/sqlitebrowser)

我的是 Mac, 已经自带了。

1. 在项目根目录，利用 DB Browser for SQLite 创建一个 pipe.db。

2. 我这边更新下它的 Dockerfile，主要是升级到 1.8.6
```sh
FROM alpine:3.7
LABEL maintainer = "abcdsxg@gmail.com"

ENV PIPE_VERSION 1.8.6
ENV GLIBC_VERSION 2.28-r0

WORKDIR /opt/

RUN set -ex && \
    apk --no-cache add ca-certificates && \
    wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub && \
    wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/glibc-${GLIBC_VERSION}.apk && \
    apk add glibc-${GLIBC_VERSION}.apk && \
    wget -O pipe${PIPE_VERSION}.zip https://github.com/b3log/pipe/releases/download/v${PIPE_VERSION}/pipe-v${PIPE_VERSION}-linux.zip && \
    unzip pipe${PIPE_VERSION}.zip && \
    chmod +x pipe && \
    rm -f pipe${PIPE_VERSION}.zip glibc-${GLIBC_VERSION}.apk

CMD ["/opt/pipe"]
```

3. 安装 docker & docker-compose 很简单，大家官网 step by step 跟着走就 OK 了。
```sh
docker-compose --help # 看下帮助
# up                 Create and start containers
# down               Stop and remove containers, networks, images, and volumes

docker-compose up --help
#     -d, --detach               Detached mode: Run containers in the background,
#                                print new container names. Incompatible with
#                                --abort-on-container-exit.

docker-compose up -d # 起来，然后背后运行
docker-compose down # 销毁
```

4. 进入 [http://localhost:5897](http://localhost:5897)

5. 可以先使用本地账号初始化
