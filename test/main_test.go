package test

import (
	"database/sql"
	"fmt"
	v3 "goapi/app/controllers/client/v3"
	"goapi/bootstrap"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/logger"
	"testing"
	"time"
)

func Init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("binary-client")
}

func Test_Main(t *testing.T) {
	Init()
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

	RES := v3.GetVipInfo("SMO", "Indonesia", 1)
	fmt.Println(RES)
	fmt.Println(RES)
	fmt.Println(RES)
}
