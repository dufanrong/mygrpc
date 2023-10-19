#!/bin/bash

# 检查是否有 root 权限运行脚本
if [ "$EUID" -ne 0 ]; then
  echo "请以 root 权限运行此脚本"
  exit 1
fi

# 指定要查找的端口列表
ports=("50051" "50052" "50053" "50054" "50055")

for port in "${ports[@]}"; do
  # 检查系统中是否有占用指定端口的进程
  if lsof -i :$port -t &>/dev/null; then
    # 找到占用指定端口的进程并结束它
    lsof -i :$port -t | xargs kill -9
    echo "已结束占用 $port 端口的进程"
  else
    echo "系统中没有占用 $port 端口的进程"
  fi
done
