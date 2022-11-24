package heartbeat

import (
	"fmt"
	"goapi/pkg/logger"
	"goapi/service/kline-binance/binance"
	"goapi/service/kline-common"
	"goapi/service/kline-common/source"
	"goapi/service/kline-hsb/hsb"
	"goapi/service/kline-olymptrade/olymptrade"
	"goapi/service/kline-qxbroker/qxbroker"
	"goapi/service/kline-raceoption/raceoption"
	"time"
)

// 心跳 heartbeat 健康

func Start(num time.Duration, Source string, CodeData []KlineCommon.Info) {
	time.Sleep(time.Minute * 5)                 //等待当前数据订阅5分钟后开始心跳检测
	ticker := time.NewTicker(num * time.Second) // 每隔 num 秒检查一次
	for range ticker.C {
		for _, info := range CodeData {
			// 休市的数据源不检测状态
			//if helpers.InArray(Source, []string{"olymptrade", "hsb"}) && kline.CheckSourceSleep(Source, info.ContractCode) {
			//	continue
			//}
			status := source.CheckStatus(Source, info.ContractCode)
			logger.Info(fmt.Sprintf("数据源【%v】对应资产【%v】健康状态：【%v】", Source, info.ContractCode, status))
		}
	}
}

// 心跳检测到异常了，重启协诚
func reboot(Source string, info KlineCommon.Info) {
	switch Source {
	case "raceoption":
		go raceoption.SocketRaceOption(info)
		break
	case "qxbroker":
		go qxbroker.SocketQxbroker(info)
		break
	case "olymptrade":
		go olymptrade.SocketOlymptrade(info)
		break
	case "hsb":
		go hsb.SocketHsb(info)
		break
	case "binance":
		go binance.SocketBinance(info)
		break
	default:
		break
	}
}
