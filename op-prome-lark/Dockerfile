# 第一阶段：构建阶段
FROM golang:1.23.0-alpine AS builder

# 设置工作目录
WORKDIR /app

# 将 go.mod 和 go.sum 文件复制到工作目录
COPY go.mod go.sum ./

# 下载依赖（此时不需要复制整个源代码）
RUN go mod download

# 将源代码复制到工作目录
COPY . .

# 编译 Go 应用
RUN go build -o op-prome-feishu .

# 第二阶段：运行阶段
FROM alpine:latest

# 安装需要的运行时依赖 (如果有的话，举个例子安装 libc)
RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /root/

# 从构建阶段复制二进制文件到最终镜像
COPY --from=builder /app/op-prome-feishu .

# 开放应用监听的端口（例如，假设应用运行在 8080 端口）
EXPOSE 8080

# 设置容器启动命令
CMD ["./op-prome-feishu"]

