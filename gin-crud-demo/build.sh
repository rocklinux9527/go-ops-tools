#!/bin/bash
# 获取当前设备的操作系统类型
os=$(uname -s)

# 根据操作系统类型执行不同的编译命令
if [ "$os" == "Darwin" ]; then
    # Mac 平台
    go build -ldflags="-s -w" -o ./cmd/gin-crud-demo-macos main.go
    chmod +x ./cmd/gin-crud-demo-macos
elif [ "$os" == "Linux" ]; then
    # Linux 平台
    GOOS=linux GOARCH=amd64  go build -o ./cmd/gin-crud-demo-linux main.go
    chmod +x ./cmd/gin-crud-demo-linux
else
    echo "Unsupported operating system: $os"
    exit 1
fi
