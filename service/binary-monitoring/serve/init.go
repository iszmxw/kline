package serve

import (
	"encoding/json"
	"goapi/app/response"
	"goapi/pkg/logger"
)

type Service struct {
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

// 初始化参数

func (p *Service) Init(config response.TradingAssetsMonitoring) bool {
	// 最后交易的开始时间节点
	p.LastTradeBeginTimestamp = 0
	// 最后交易的结束时间节点
	p.LastTradeEndTimestamp = 0
	// 买涨的人数
	p.upNum = 0
	// 买涨的金额
	p.upAmounts = 0.00
	// 买跌的人数
	p.downNum = 0
	// 买跌的金额
	p.downAmounts = 0.00
	// HClose 最大 close
	p.HClose = 1000000
	// LClose 最小 close
	p.LClose = 0.00
	p.winLoss = 0.00
	p.WinType = ""
	p.Peoples = make([]string, 0)
	p.OrderIds = make([]string, 0)
	// 插针配置
	var AssetsInsertFrame response.AssetsInsertFrame
	err := json.Unmarshal([]byte(config.AssetsInsertFrame), &AssetsInsertFrame)
	if err != nil {
		logger.Info(err)
		return false
	}
	p.config = response.TradingAssetsMonitoringConfig{
		ID:           config.ID,
		AssetsCode:   config.AssetsCode,
		PointNumber:  config.PointNumber,
		IsMonitoring: config.IsMonitoring,
		//ShockAmplitude:      config.ShockAmplitude,
		VolatilityAmplitude: config.VolatilityAmplitude,
		VolatilityFrequency: config.VolatilityFrequency,
		AssetsInsertFrame: response.AssetsInsertFrame{
			AreaTimer:               AssetsInsertFrame.AreaTimer,
			IntervalTimer:           AssetsInsertFrame.IntervalTimer,
			InAdvanceTimer:          AssetsInsertFrame.InAdvanceTimer,
			LossThreshold:           AssetsInsertFrame.LossThreshold,
			EarlyExecutionTimer:     AssetsInsertFrame.EarlyExecutionTimer,
			EarlyExecutionFrequency: AssetsInsertFrame.EarlyExecutionFrequency,
			RecoverTimer:            AssetsInsertFrame.RecoverTimer,
			RecoverFrequency:        AssetsInsertFrame.RecoverFrequency,
		},
	}
	return true
}
