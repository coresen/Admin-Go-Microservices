# 使用官方 Go 镜像作为构建环境
FROM golang:1.24.2

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制全部源代码
COPY . .

# 构建应用 (假设你的 main 文件是 user.go)
RUN go build -o /app/bin/user ./user.go

# 使用轻量级 Alpine 镜像作为运行时
FROM alpine:3.18

# 安装必要的运行时依赖
RUN apk add --no-cache ca-certificates tzdata

# 设置工作目录
WORKDIR /app

# 从构建阶段复制二进制文件和配置文件
COPY --from=builder /app/bin/user .
COPY --from=builder /app/etc/user.yaml ./etc/

# 暴露服务端口 (与 user.yaml 中的端口一致)
EXPOSE 8888

# 启动服务
ENTRYPOINT ["./user", "-f", "etc/user.yaml"]