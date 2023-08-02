package main

import (
	"encoding/json"
	"fmt"
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/logger"
	"goapi/pkg/pprof"
	"goapi/pkg/redis"
	KlineCommon "goapi/service/kline-common"
	"goapi/service/kline-common/heartbeat"
	"goapi/service/kline-mql5/mql5"
	kconfig "goapi/service/kline-mql5/mql5/config"
	"goapi/service/kline-mql5/mql5/wsdata"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init(kconfig.ServiceLog)
	// 初始化 SQL
	bootstrap.SetupDB()
}

func main() {
	// pprof 分析
	pprof.Debug(":6071")
	bootstrap.SetupRedis(15)
	defer bootstrap.RedisClose()
	CodeData := KlineCommon.GetData(kconfig.KlineSource)
	// 启动监控，监控数据源活跃状态，设置参数每个30秒检测一次
	go heartbeat.Start(30, kconfig.KlineSource, CodeData)
	start := "ok"
	for {
		if start == "ok" {
			start = "no"
			logger.Info(CodeData)
			for _, obj := range CodeData {
				logger.Info("缓存 " + obj.Symbol)
				go SubPriceKLine(obj)
			}
		}
		time.Sleep(time.Second)
	}
}

func SubPriceKLine(info KlineCommon.Info) {
	pushKey := fmt.Sprintf("%s:%s", "list-mql5", info.ContractCode)
	for {
		data := redis.Client.RPop(pushKey)
		if len(data) <= 0 {
			time.Sleep(time.Millisecond * 300)
			continue
		}
		var mql5data wsdata.Mql5
		err := json.Unmarshal([]byte(data), &mql5data)
		if err != nil {
			time.Sleep(time.Millisecond * 300)
			continue
		}
		mql5.SocketMql5(info, mql5data)
	}
}
