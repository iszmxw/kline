package main

import (
	"fmt"
	"goapi/config"
	"goapi/pkg/logger"
	"math"
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
}

func main() {
	//data := trade_sleep.GetTimeIsSleep("Apple")
	////res := xjson.GetConfig("./service/kline-olymptrade/olymptrade/config/source_time/Adidas.json")
	////fmt.Println(res)
	//fmt.Println(data)
	fmt.Println(math.Floor(58.99999))
}
