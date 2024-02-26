# 使用适用于 x86_64 架构的基础镜像
FROM golang:1.16 AS builder

# 设置工作目录
WORKDIR /app

# 拷贝 Go 项目文件到工作目录
COPY . .

# 设置目标操作系统和架构为 Linux x86_64
ENV GOARCH=amd64
ENV GOOS=linux

# 构建 Go 应用
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o myhealthcheck .

# 使用轻量的 alpine 镜像作为运行时镜像
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 从构建阶段复制二进制文件到运行时镜像
COPY --from=builder /app/myhealthcheck .

# 暴露服务端口
EXPOSE 8080

# 运行应用程序
CMD ["./myhealthcheck"]