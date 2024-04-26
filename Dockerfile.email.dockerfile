# 使用官方的 Golang 镜像作为基础镜像
FROM golang:latest AS build

# 创建一个工作目录
WORKDIR /work

# 将本地代码拷贝到容器中
COPY . .

# 设置 Go 模块代理
ENV GOPROXY="https://goproxy.cn"

# 编译应用程序
RUN GOOS=linux GOARCH=amd64 go build -o iqoption-email service/binary-email/main.go && \
    mv iqoption-email binary/email

# 运行阶段 使用简洁的作为基础镜像
FROM --platform=linux/amd64 alpine:latest

# 安装 libc6-compat 包，提供 /lib64/ld-linux-x86-64.so.2 文件
RUN apk --no-cache add libc6-compat

# 设置工作目录
WORKDIR /home/yezhiming/binary/email

# 将打包后的程序拷贝到容器中
COPY --from=build /work/binary/email /home/yezhiming/binary/email

# 创建配置文件挂载目录
RUN mkdir -p /home/yezhiming/binary/email/config

# 暴露应用程序使用的端口
EXPOSE 7107

# 运行应用程序
CMD ["./iqoption-email", "--APP_PORT=7107", ">/dev/null", "2>error.log", "&"]
