package main

import (
	"fmt"
	"goapi/bootstrap"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/logger"
	"goapi/pkg/redis"
	KlineCommon "goapi/service/kline-common"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("test")
}

func main() {
	// 初始化 SQL
	// 定义日志目录
	logger.Info("初始化 SQL")
	bootstrap.SetupDB()
	// 初始化 Redis
	logger.Info("初始化 Redis")
	db := conf.GetInt("redis.db")
	bootstrap.SetupRedis(db)
	defer bootstrap.RedisClose()
	for true {
		go func() {
			SystemTs := time.Now().UnixMilli()
			// 插针后的数据缓存到另外一张表
			RecoverSource, _ := redis.Client.Get(fmt.Sprintf("recover_source:%v", "test"))
			if len(RecoverSource) == 0 {
				logger.Info(KlineCommon.ConsumeTime(SystemTs, time.Now().UnixMilli()))
				return
			}
		}()
	}
}
