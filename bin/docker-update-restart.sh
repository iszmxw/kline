#!/bin/bash

# 设置脚本在任何命令返回非零退出码时立即退出
set -e

# 检查是否提供了服务名称作为参数
if [ -z "$1" ]; then
  echo "请提供服务名称作为参数。示例："
  echo " client / web / email / kline"
  exit 1
fi

# 设置变量
SERVICE_NAME="xw-$1"
BINARY_DIR="/home/yezhiming/binary"

# 设置工作目录
cd "$BINARY_DIR"

# 拉取最新镜像
pull_image() {
  echo -e "\033[35m拉取最新镜像...\033[0m"
  docker-compose pull "$SERVICE_NAME" # 使用docker-compose拉取指定服务的最新镜像
  echo -e "\033[35m镜像拉取完成...\033[0m"
}

# 休息片刻
rest() {
  echo -e "\033[35m休息片刻 3s ...\033[0m"
  sleep 3 # 休息3秒钟
}

# 单独更新并重启指定服务
restart_service() {
  echo -e "\033[35m更新并重启服务 $SERVICE_NAME ...\033[0m"
  docker-compose up -d --no-deps "$SERVICE_NAME" # 更新并以后台方式启动指定服务，不重建依赖的服务
}

# 清理不再使用的镜像
cleanup_images() {
  echo -e "\033[35m清理不再使用的镜像...\033[0m"
  docker image prune -f # 清理不再使用的镜像
}

# 获取当前目录和本机 IP 地址
get_info() {
  CURRENT_DIR=$(pwd) # 获取当前目录
  LOCAL_IP=$(curl -s ip.sb) # 获取本机 IP 地址

  echo "当前目录: $CURRENT_DIR"
  echo "本机 IP 地址: $LOCAL_IP"
}

# 执行流程
pull_image
rest
restart_service
cleanup_images
get_info
