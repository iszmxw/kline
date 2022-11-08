package main

import (
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/logger"
	"goapi/pkg/pprof"
	"goapi/service/kline-common"
	"goapi/service/kline-common/heartbeat"
	"goapi/service/kline-fxcm/fxcm"
	"goapi/service/kline-fxcm/fxcm/api"
	kconfig "goapi/service/kline-fxcm/fxcm/config"
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
	CodeData := KlineCommon.GetData(kconfig.KlineSource)
	// 费率采集处理
	//go api.GetRate(10) // 每隔10秒收集一次赔率
	// 启动监控，监控数据源活跃状态，设置参数每个30秒检测一次
	go heartbeat.Start(30, kconfig.KlineSource, CodeData)
	start := "ok"
	wsClient := new(api.FXCM)
	if err := wsClient.Start(); err != nil {
		logger.Info(err)
		return
	}
	for {
		if start == "ok" {
			wsClient.CloseChanCallback = make(chan string, 20)
			for _, obj := range CodeData {
				go fxcm.SocketFxcm(wsClient, obj)
			}
			start = "no"
		}
		time.Sleep(time.Second)
	}
}
