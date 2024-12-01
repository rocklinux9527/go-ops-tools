#!/bin/bash
# 获取当前设备的操作系统类型
os=$(uname -s)

# 根据操作系统类型执行不同的编译命令
if [ "$os" == "Darwin" ]; then
    # Mac 平台
    go build -ldflags="-s -w" -o ./cmd/op-prom-feishu-macos main.go
    chmod +x ./cmd/op-prom-feishu
elif [ "$os" == "Linux" ]; then
    # Linux 平台
    GOOS=linux GOARCH=amd64  go build -o ./cmd/op-prom-feishu-linux main.go
    chmod +x ./cmd/op-prom-feishu
else
    echo "Unsupported operating system: $os"
    exit 1
fi
