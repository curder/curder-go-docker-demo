## go docker demo

go项目运行在docker中

## 构建镜像

```
docker build -t curder/go-docker-demo .
```

## 运行镜像

```
docker run -p 8080:80 -it curder/go-docker-demo
```

> 将容器里的`80`端口映射到宿主机的`8080`端口。

## 测试

```
curl http://127.0.0.1:8080

hello docker!
```