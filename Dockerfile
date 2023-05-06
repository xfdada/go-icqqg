# 第一阶段：构建应用程序
FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

# 移动到工作目录：/build
WORKDIR /build

# 安装依赖项
COPY go.mod go.sum ./

RUN go mod download

# 复制应用程序源代码并构建二进制文件
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

# 第二阶段：构建最终的 Docker 镜像
FROM alpine:latest

# 安装必要的运行时依赖项
WORKDIR /build/
RUN apk update --no-cache && apk add --no-cache ca-certificates && apk add tzdata

# 复制应用程序二进制文件
COPY --from=builder /build/app .
COPY ./public /build/public
COPY ./runtime /build/runtime
COPY ./resource /build/resource
COPY ./config.yaml /build/config.yaml
# 暴露应用程序的端口
EXPOSE 8080

# 运行应用程序
CMD ["/build/app"]
