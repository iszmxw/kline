package binary_monitoring

import (
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/logger"
	"goapi/pkg/pprof"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("binary_monitoring")
	// 初始化 SQL
	bootstrap.SetupDB()
}

func main() {
	// pprof 分析
	pprof.Debug(":6069")
	bootstrap.SetupRedis(15)
	defer bootstrap.RedisClose()

}
