package main

import (
	"fmt"
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/logger"
	mail_notice "goapi/pkg/notice/mail"
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
	mail_notice.Notice("msgAssetSleep", map[string]interface{}{
		"assets_code": "Apple",
		"msg":         fmt.Sprintf("%v,休市完毕，准备开市交易", "Apple"),
	})
	//go kline.TickerCheckTodaySleep(true)
	time.Sleep(time.Second * 125)

}
