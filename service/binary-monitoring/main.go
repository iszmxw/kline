package main

import (
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/logger"
	"goapi/pkg/pprof"
	"goapi/service/binary-monitoring/monitoring"
	"goapi/service/binary-monitoring/queue"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 5.5*3600) // 东七
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("binary_monitoring")
	// 初始化 SQL
	bootstrap.SetupDB()
}

func main() {
	// pprof 分析
	pprof.Debug(":6069")
	bootstrap.SetupRedis(15)
	defer bootstrap.RedisClose()
	go queue.Start()
	//monitoring.CreateInterferenceNeedle(time.Now().UnixMilli(), 0.66103, 0.66096, response.TradingAssetsMonitoring{
	//	ID:                      29,
	//	PointNumber:             5,
	//	AssetsCode:              "AUD-USD",
	//	IsMonitoring:            1,
	//	AreaTimer:               0,
	//	IntervalTimer:           60,
	//	InadvanceTimer:          20,
	//	LossThreshold:           0,
	//	EarlyexecutionTimer:     24,
	//	EarlyexecutionFrequency: 5,
	//	RecoverTimer:            0,
	//	RecoverFrequency:        0,
	//	ShockAmplitude:          0.0001,
	//	VolatilityAmplitude:     0.0001,
	//	VolatilityFrequency:     2,
	//})
	for {
		// 每隔一秒钟执行一次
		time.Sleep(time.Second)
		go monitoring.Start()
	}
}
