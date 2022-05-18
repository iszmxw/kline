package main

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/bootstrap"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/logger"
	"goapi/pkg/notice/telegram"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("telegram-bot")
	// 初始化 SQL
	bootstrap.SetupDB()
}

func main() {
	// 通知
	Info := cmap.New().Items()
	Info["title"] = fmt.Sprintf("环境【%v】%v %v 服务：恢复市场通知", conf.GetString("app.env"), "BTC-USD", "huobi")
	Info["content"] = "休市已结束，市场正在恢复，请您查收"
	telegram.Notice(Info)
}
