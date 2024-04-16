





## 启动

### 1. 使用Docker

1. 打包镜像
```go
docker build . --file=Dockerfile --network=host --platform=linux/amd64 -t bbs_web:v1.0.0
```

2. 运行

```go
docker run -d -p 8181:8181 8899:8899 --rm bbs_web:v1.0.0
```