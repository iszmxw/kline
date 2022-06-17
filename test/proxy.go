package main

import (
	"fmt"
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/logger"
	kline_common "goapi/service/kline-common"
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
	fmt.Println(kline_common.GetProxy("hsb"))
}
