FROM golang:1.21.4-alpine

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    GOPROXY=goproxy.io \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

# 复制go.mod，go.sum并且下载依赖
COPY go.* ./
RUN go mod tidy

# 复制项目内的所有内容并构建
COPY . .

RUN go build main.go

EXPOSE 8181
ENTRYPOINT [ "./main" ]