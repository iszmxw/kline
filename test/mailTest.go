package main

import (
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/logger"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("apiService")
}

func main() {
	bootstrap.SetupRedis("1")
	defer bootstrap.RedisClose()
	time.Sleep(time.Second * 100)
}
