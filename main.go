package main

import (
	"database/sql"
	"flag"
	"fmt"
	"goapi/bootstrap"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/logger"
	"goapi/pkg/notice/bark"
	"os/exec"
	"strings"
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
// @termsOfService http://127.0.0.1/docs/index.html

// @contact.name 追梦小窝
// @contact.url http://github.com/iszmxw
// @contact.email mail@54zm.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

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
	// 查询当前服务器ip
	res, err := exec.Command("curl", "ip.sb").Output()
	if err != nil {
		fmt.Println("执行命令出错：", err)
		return
	}
	ip := strings.Trim(string(res), "\n\n")
	// bark 通知
	bark.Notice(ip, "xw-client 启动成功")
	_ = router.Run(fmt.Sprintf(":%v", *AppPort))
}
