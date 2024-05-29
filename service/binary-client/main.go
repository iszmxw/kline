package main

import (
	"database/sql"
	"flag"
	"fmt"
	"goapi/bootstrap"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/notice/wechat"
	"time"

	"github.com/gin-contrib/pprof"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("binary-client")
}

// @title 用户端接口服务
// @version 3.0
// @description 3.0版本，基于之前的2.0改造的
// @termsOfService http://127.0.0.1:7104/docs/index.html

// @contact.name 追梦小窝
// @contact.email mail@54zm.com

// @host 127.0.0.1
// @BasePath
func main() {
	AppPort := flag.Int64("APP_PORT", conf.GetInt64("app.port"), "服务端口")
	flag.Parse()
	// 初始化 SQL
	logger.Info("初始化 SQL")
	DB := bootstrap.SetupDB()
	// 关闭数据库链接
	defer func(s *sql.DB) {
		err := s.Close()
		if err != nil {
			logger.Error(err)
		}
	}(DB)
	// 初始化 Redis
	logger.Info("初始化 Redis")
	db := conf.GetInt("redis.db")
	bootstrap.SetupRedis(db)
	defer bootstrap.RedisClose()
	bootstrap.SetupMongoDB()
	// 初始化路由绑定
	logger.Info("加载 client 路由")
	router := bootstrap.SetupRoute()
	pprof.Register(router) // 开启 pprof
	// 启动路由
	logger.Info("启动路由")
	logger.Info(fmt.Sprintf("当前环境:%v", *AppPort))
	// wechat 通知
	wechat.Notice(helpers.IP(), "xw-client 启动成功")
	_ = router.Run(fmt.Sprintf(":%v", *AppPort))
}
