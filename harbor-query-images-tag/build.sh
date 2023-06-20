#!/bin/bash
# 获取当前设备的操作系统类型
os=$(uname -s)

# 根据操作系统类型执行不同的编译命令
if [ "$os" == "Darwin" ]; then
    # Mac 平台
    go build -ldflags="-s -w" -o ./cmd/harbor-image-get-tag-macos main.go
    chmod +x ./cmd/harbor-image-get-tag-macos
elif [ "$os" == "Linux" ]; then
    # Linux 平台
    GOOS=linux GOARCH=amd64  go build -o ./cmd/harbor-image-get-tag-linux main.go
    chmod +x ./cmd/harbor-image-get-tag-linux
else
    echo "Unsupported operating system: $os"
    exit 1
fi
