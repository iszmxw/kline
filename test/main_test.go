package test

import (
	"database/sql"
	"fmt"
	"goapi/bootstrap"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/logger"
	"goapi/pkg/xtime"
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

	offset := xtime.GmtOffset("GMT+07:00")
	fmt.Println(offset)
	// 时区设置
	timeZone := time.FixedZone("CST", offset)
	// 获取本周的最后一天日期
	endDate := xtime.GetEndDateOfATimeZoneWeek(0, timeZone)
	fmt.Println("获取本周的最后一天日期:", endDate)
	// 获取上周的最后一天日期
	endDateLastWeek := xtime.GetEndDateOfATimeZoneWeek(1, timeZone)
	fmt.Println("获取上周的最后一天日期:", endDateLastWeek)
	// 获取上上周的最后一天日期
	endDateTwoWeeksAgo := xtime.GetEndDateOfATimeZoneWeek(2, timeZone)
	fmt.Println("获取上上周的最后一天日期:", endDateTwoWeeksAgo)
	// 其他周的最后一天日期依此类推...
	//
	//RES := commonUtils.GetVipInfo("SMO", "Indonesia", 1)
	//fmt.Println(RES)
	//fmt.Println(RES)
	//fmt.Println(RES)
}
