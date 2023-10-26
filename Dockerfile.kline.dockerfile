# 使用官方的 Golang 镜像作为基础镜像
FROM golang:latest AS build

# 创建一个工作目录
WORKDIR /work

# 将本地代码拷贝到容器中
COPY . .

# 编译应用程序
# 打包 kline-binance K线图socket采集服务
RUN GOOS=linux GOARCH=amd64 go build -o iqoption-go-binance service/kline-binance/main.go && \
    GOOS=linux GOARCH=amd64 go build -o iqoption-go-hsb service/kline-hsb/main.go && \
    GOOS=linux GOARCH=amd64 go build -o iqoption-go-huobi service/kline-huobi/main.go && \
    GOOS=linux GOARCH=amd64 go build -o iqoption-go-mql5 service/kline-mql5/main.go && \
    GOOS=linux GOARCH=amd64 go build -o iqoption-go-iqoption service/kline-iqoption/main.go && \
    GOOS=linux GOARCH=amd64 go build -o iqoption-go-binomo service/kline-binomo/main.go && \
    GOOS=linux GOARCH=amd64 go build -o iqoption-go-olymptrade service/kline-olymptrade/main.go && \
    # 打包 binary-monitoring 插针服务
    GOOS=linux GOARCH=amd64 go build -o iqoption-go-monitoring service/binary-monitoring/main.go && \
    mv iqoption-go-* binary/kline

# 运行阶段 使用简洁的作为基础镜像
FROM --platform=linux/amd64 alpine:latest

# 设置工作目录
WORKDIR /home/yezhiming/binary/kline

# 将打包后的程序拷贝到容器中
COPY --from=build /work/binary/kline /home/yezhiming/binary/kline

# 创建配置文件挂载目录
RUN mkdir -p /home/yezhiming/binary/kline/config

# 运行应用程序
CMD ["sh","./startKline.sh"]