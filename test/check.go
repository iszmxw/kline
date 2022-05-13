package main

import (
	"goapi/bootstrap"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/health"
	"goapi/pkg/logger"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("apiService")
}

func main() {
	// 初始化 SQL
	// 定义日志目录
	logger.Info("初始化 SQL")
	bootstrap.SetupDB()
	// 初始化 Redis
	logger.Info("初始化 Redis")
	db := conf.GetString("redis.db")
	bootstrap.SetupRedis(db)
	defer bootstrap.RedisClose()
	start := true
	for {
		if start {
			// 启动监控，监控数据源活跃状态，设置参数每个30秒检测一次
			go health.Heartbeat(30)
		}
		start = false
	}
}
