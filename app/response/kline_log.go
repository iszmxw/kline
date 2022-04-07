package response

import (
	"github.com/gostudys/huobi_contract_golang/src/ws/response"
	"goapi/app/models"
	"goapi/pkg/helpers"
	"time"
)

// KlineLog  K线数据表
type KlineLog struct {
	LogId        int64              `json:"log_id"`        //  k线表主键id
	ContractCode string             `json:"contract_code"` //  合约代码 btc-usd
	Ch           string             `json:"ch"`            //  k线类型请求kline的wss请求地址
	Ts           int64              `json:"ts"`            //  kilne返回日期时间戳
	ID           int64              `json:"id"`            //  返回的id
	Mrid         float64            `json:"mrid"`          //  id
	COpen        float64            `json:"c_open"`        //  起始峰值
	CClose       float64            `json:"c_close"`       //  结束峰值
	High         float64            `json:"high"`          //  最高峰值
	Low          float64            `json:"low"`           //  最低峰值
	Amount       float64            `json:"amount"`        //  成交金额
	Vol          float64            `json:"vol"`           //  成交手
	CCount       float64            `json:"c_count"`       //  交易笔数
	TsDate       helpers.TimeNormal `json:"ts_date"`       //  kline返回日期
	CreateDate   helpers.TimeNormal `json:"create_date"`   //  创建日期
	KlineSource  string             `json:"kline_source"`  //  k线来源 huobi
}

type WsIndex struct {
	Ch   string `json:"ch"`
	Ts   int64  `json:"ts"`
	Tick struct {
		Id     int     `json:"id"`
		Vol    string  `json:"vol"`
		Count  float64 `json:"count"`
		Open   string  `json:"open"`
		Close  string  `json:"close"`
		Low    string  `json:"low"`
		High   string  `json:"high"`
		Amount string  `json:"amount"`
	} `json:"tick"`
}

// KlineLogData 组装数据
func KlineLogData(Topic []string, data *response.SubIndexKLineResponse) *models.KlineLog {
	return &models.KlineLog{
		ContractCode: Topic[1],                                      // 合约代码 BTC-USD
		Ch:           data.Ch,                                       // k线类型请求kline的wss请求地址
		Ts:           data.Ts,                                       // kilne返回日期时间戳
		ID:           data.Tick.Id,                                  // 返回的id
		Mrid:         0,                                             // 返回的id
		COpen:        helpers.StrToFloat64(data.Tick.Open),          // 起始峰值
		CClose:       helpers.StrToFloat64(data.Tick.Close),         // 结束峰值
		High:         helpers.StrToFloat64(data.Tick.High),          // 最高峰值
		Low:          helpers.StrToFloat64(data.Tick.Low),           // 最低峰值
		Amount:       helpers.StrToFloat64(data.Tick.Amount),        // 成交金额
		Vol:          helpers.StrToFloat64(data.Tick.Vol),           // 成交手
		CCount:       data.Tick.Count,                               // 交易笔数
		TsDate:       time.Unix(0, data.Ts*int64(time.Millisecond)), // kline返回日期
		CreateDate:   time.Now(),                                    // 创建日期
		KlineSource:  "huobi",                                       // k线来源 huobi
	}
}
