package main

import (
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/kline"
	"goapi/pkg/logger"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("test")
	// 初始化 SQL
	bootstrap.SetupDB()
}

func main() {
	bootstrap.SetupRedis(0)
	defer bootstrap.RedisClose()
	kline.UpdateSleep("kline-system", "GBP-JPY", "1", "1")
	time.Sleep(20)

}
