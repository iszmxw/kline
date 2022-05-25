package main

import (
	"goapi/config"
	"goapi/pkg/logger"
	kline_common "goapi/service/kline-common"
	"time"
)

func init() {
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("test")
}

func main() {
	goroutine()
	time.Sleep(time.Hour)
}

func goroutine() {
	defer kline_common.KlineRecover("异常", "asasa", "asasas")
	c := make(chan int, 10)
	c <- 1
	c <- 2
	close(c)
	c <- 3
}
