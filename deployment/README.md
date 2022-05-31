# Docker Deployment

## Server

**Need mongodb on your server first.**

### Download or Compile

Download by [releases](https://github.com/SukiEva/aldb/releases)

### Create Dockerfile

```
FROM alpine:latest
MAINTAINER SukiEva<dev.suki@outlook.com>
ENV VERSION 1.0

# 在容器根目录 创建一个 apps 目录
WORKDIR /apps

# 挂载容器目录
# VOLUME ["/apps"]

# 拷贝当前目录下可执行文件
COPY server /apps/server

# 拷贝配置文件到容器中
COPY config.json /apps/config.json

# 设置时区为上海
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

# 设置编码
ENV LANG C.UTF-8

# 暴露端口
EXPOSE 8888

# 运行golang程序的命令
ENTRYPOINT /apps/server
```

### Build Image

```bash
docker build -t aldb .
```

### Put File

```bash
mkdir /data/aldb
mv server /data/aldb/server
mv config.json /data/aldb/config.json
chmod 0777 /data/aldb/server
```


### Run

- change the port:port

```bash
docker run -p 8888:8888 -v /data/aldb:/apps --name aldb --network host --env GIN_MODE=release -d aldb
```

## Web

