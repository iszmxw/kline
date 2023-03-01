package main

import (
	"fmt"
	"github.com/google/uuid"
	"goapi/bootstrap"
	"goapi/config"
	goadjust "goapi/pkg/adjust"
	conf "goapi/pkg/config"
	"goapi/pkg/logger"
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

	fmt.Println()
	fmt.Println(uuid.NewUUID())
	fmt.Println(uuid.NewUUID())
	fmt.Println(uuid.NewUUID())

	goadjust.SubmitEvent(goadjust.EventParams{
		EventName:        "trade",
		AppUserId:        "7FF235ACCD4",
		IsRepeatedSubmit: "0",
		Mobile:           "",
		EventBindId:      "482bd535-7fcd-49ea-9678-9b9d15e81737",
		Email:            "1360791050@qq.com",
		FlowJsonString:   "",
		DeviceBrand:      "Xiaomi",
	}, "India", "adjust")

	time.Sleep(60 * time.Second)
}
