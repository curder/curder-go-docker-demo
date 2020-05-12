## go docker demo

go项目运行在docker中

## 构建镜像

```
docker build -t curder/go-docker-demo .
```

## 运行镜像

```
docker run -p 8080:8081 -it curder/go-docker-demo
```