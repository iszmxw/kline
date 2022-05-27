package main

import (
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/logger"
	"goapi/pkg/pprof"
	kline_common "goapi/service/kline-common"
	"goapi/service/kline-common/heartbeat"
	"goapi/service/kline-raceoption/raceoption"
	"goapi/service/kline-raceoption/raceoption/api"
	kconfig "goapi/service/kline-raceoption/raceoption/config"
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
	pprof.Debug(":6061")
	bootstrap.SetupRedis(15)
	defer bootstrap.RedisClose()
	CodeData := kline_common.GetData(kconfig.KlineSource)
	start := "ok"
	rate := true
	for {
		if start == "ok" {
			for _, obj := range CodeData {
				logger.Info("缓存 " + obj.Symbol)
				go raceoption.SocketRaceOption(obj)
			}
			start = "no"
		}
		if rate {
			rate = false
			// 费率采集处理
			go api.GetRate()
			// 启动监控，监控数据源活跃状态，设置参数每个30秒检测一次
			go heartbeat.Start(30, kconfig.KlineSource, CodeData, true)
		}
		time.Sleep(time.Second)
	}
}
