package models

import (
	"goapi/pkg/config"
	"time"
)

// KlineLog  K线数据表
type KlineLog struct {
	LogId        int64     `json:"log_id"`        //  k线表主键id
	ContractCode string    `json:"contract_code"` //  合约代码 btc-usd
	Ch           string    `json:"ch"`            //  k线类型请求kline的wss请求地址
	Ts           int64     `json:"ts"`            //  kilne返回日期时间戳
	ID           int64     `json:"id"`            //  返回的id
	Mrid         float64   `json:"mrid"`          //  id
	COpen        float64   `json:"c_open"`        //  起始峰值
	CClose       float64   `json:"c_close"`       //  结束峰值
	High         float64   `json:"high"`          //  最高峰值
	Low          float64   `json:"low"`           //  最低峰值
	Amount       float64   `json:"amount"`        //  成交金额
	Vol          float64   `json:"vol"`           //  成交手
	CCount       float64   `json:"c_count"`       //  交易笔数
	TsDate       time.Time `json:"ts_date"`       //  kline返回日期
	CreateDate   time.Time `json:"create_date"`   //  创建日期
	KlineSource  string    `json:"kline_source"`  //  k线来源 huobi
}

func (m *KlineLog) TableName() string {
	prefix := config.GetString("database.mysql.prefix")
	table := config.GetString("table.kline_log")
	return prefix + table
}
