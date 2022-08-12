package main

import (
	"flag"
	"fmt"
	"goapi/bootstrap"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/logger"
	"time"

	"github.com/gin-contrib/pprof"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("binary-web")
	// 初始化 SQL
	bootstrap.SetupDB()
}

func main() {
	APP_PORT := flag.Int64("APP_PORT", conf.GetInt64("app.port"), "服务端口")
	flag.Parse()
	// 初始化 SQL
	// 定义日志目录
	logger.Info("初始化 SQL")
	bootstrap.SetupDB()
	// 初始化 Redis
	logger.Info("初始化 Redis")
	db := conf.GetInt("redis.db")
	bootstrap.SetupRedis(db)
	defer bootstrap.RedisClose()
	// 初始化路由绑定
	logger.Info("加载路由")
	router := bootstrap.SetupRouteWeb()
	pprof.Register(router) // 开启 pprof
	// 启动路由
	logger.Info("启动 Web 路由")
	logger.Info(fmt.Sprintf("当前环境:%v", *APP_PORT))
	_ = router.Run(fmt.Sprintf(":%v", *APP_PORT))
}
