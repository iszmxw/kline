package main

import (
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/logger"
	"goapi/service/kline-raceoption/raceoption"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("kline-raceoption")
	// 初始化 SQL
	bootstrap.SetupDB()
}

//var addr = flag.String("addr", "trading.raceoption.com", "服务器地址")
//var subMsg = flag.String("subMsg", `{"operation":"SUBSCRIBE.TICK","id":"EUR/USD","param":""}`, "订阅主题")

//var subMsg1 = flag.String("subMsg1", `{"sub":"market.ETH-USD.index.1min","id":"EUR/USD","param":""}`, "订阅主题")

func main() {
	bootstrap.SetupRedis("15")
	defer bootstrap.RedisClose()
	//flag.Parse()
	var CodeData []raceoption.Info
	CodeData = append(CodeData,
		raceoption.Info{Symbol: "EUR/USD", Desc: "欧元/美元", Referer: "https://finance.sina.com.cn/money/forex/hq/EURUSD.shtml"},
		raceoption.Info{Symbol: "GBP/USD", Desc: "英镑/美元", Referer: "https://finance.sina.com.cn/money/forex/hq/GBPUSD.shtml"},
		raceoption.Info{Symbol: "AUD/CAD", Desc: "澳大利亚元/加拿大元", Referer: "https://finance.sina.com.cn/money/forex/hq/AUDCAD.shtml"},
		raceoption.Info{Symbol: "USD/JPY", Desc: "美元/日元", Referer: "https://finance.sina.com.cn/money/forex/hq/USDJPY.shtml"},
		raceoption.Info{Symbol: "XAU/USD", Desc: "伦敦金", Referer: "https://finance.sina.com.cn/futures/quotes/XAU.shtml"},
		raceoption.Info{Symbol: "XAG/USD", Desc: "伦敦银", Referer: "https://finance.sina.com.cn/futures/quotes/XAG.shtml"},
	)
	logger.Info(CodeData)
	raceoption.Client(`{"operation":"SUBSCRIBE.TICK","id":"XAU/USD","param":""}`)
}
