package config

import "goapi/app/response"

// KlineSource 来源
const KlineSource = "SYSTEM"

type Params struct {
	// 最后交易的开始时间节点
	LastTradeBeginTimestamp int64
	// 最后交易的结束时间节点
	LastTradeEndTimestamp int64
	// 买涨的人数
	upNum int64
	// 买涨的金额
	upAmounts float64
	// 买跌的人数
	downNum int64
	// 买跌的金额
	downAmounts float64
	// HClose 最大 close
	HClose float64
	// LClose 最小 close
	LClose float64
	// avgClose 最近一分钟平均 close
	avgClose float64
	// maxClose 一分钟内的最大 close
	maxClose float64
	// minClose 一分钟内的最小 close
	minClose float64
	// diffClose 差异 close
	diffClose float64
	winLoss   float64
	WinType   string
	Peoples   []string
	OrderIds  []string
	config    response.TradingAssetsMonitoringConfig
}
