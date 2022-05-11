package response

import (
	"fmt"
	"goapi/pkg/logger"
	"goapi/pkg/redis"
)

// 默认数据源获取

func KlineSource(AssetsCode string) string {
	key := fmt.Sprintf("kline_source:%v", AssetsCode)
	Source, err := redis.Client.Get(key)
	if err != nil && err.Error() != "redis: nil" {
		logger.Error(err)
		Source = ""
	}
	if len(Source) > 0 {
		return Source
	}
	switch AssetsCode {
	case "ETH-USD":
		Source = "huobi"
		break
	case "BTC-USD":
		Source = "huobi"
		break
	case "EUR-USD":
		Source = "raceoption"
		break
	case "GBP-USD":
		Source = "raceoption"
		break
	case "AUD-CAD":
		Source = "raceoption"
		break
	case "USD-JPY":
		Source = "raceoption"
		break
	case "GOLD":
		Source = "raceoption"
		break
	case "SILVER":
		Source = "raceoption"
		break
	case "Netflix":
		Source = "qxbroker"
		break
	case "Microsoft":
		Source = "qxbroker"
		break
	case "Intel":
		Source = "qxbroker"
		break
	default:
		Source = "unknown"
		break
	}
	return Source
}
