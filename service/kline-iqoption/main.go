package main

import (
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/logger"
	"goapi/pkg/pprof"
	"goapi/service/kline-common"
	"goapi/service/kline-iqoption/iqoption"
	kconfig "goapi/service/kline-iqoption/iqoption/config"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init(kconfig.ServiceLog)
	// 初始化 SQL
	bootstrap.SetupDB()
}

func main() {
	// pprof 分析
	pprof.Debug(":6066")
	bootstrap.SetupRedis(15)
	defer bootstrap.RedisClose()
	CodeData := KlineCommon.GetData(kconfig.KlineSource)
	// 启动监控，监控数据源活跃状态，设置参数每个30秒检测一次
	//go heartbeat.Start(30, kconfig.KlineSource, CodeData)
	start := "ok"
	for {
		if start == "ok" {
			start = "no"
			logger.Info(CodeData)
			for _, obj := range CodeData {
				logger.Info("缓存 " + obj.Symbol)
				go iqoption.SocketIqoption(obj)
			}
		}
		time.Sleep(time.Second)
	}
}
