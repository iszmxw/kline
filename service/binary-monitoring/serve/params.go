package serve

import (
	"encoding/json"
	"fmt"
	"goapi/app/models"
	"goapi/app/response"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/service/binary-monitoring/config"
	"time"
)

// RedisData 缓存进redis的数据
func redisData(AssetsCode, Source string, SystemTs int64, close float64) string {
	var data response.SubIndexKLineResponse
	data.Ch = fmt.Sprintf("market.%v.index.1min", AssetsCode)
	data.Ts = SystemTs                                                                             // 系統生成的Ts
	data.TsDate = time.Unix(0, SystemTs*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999") // 系統生成的Ts
	data.Source = Source
	data.Tick.Id = SystemTs
	data.Tick.Vol = ""
	data.Tick.Count = 0
	data.Tick.Open = "0.00"
	data.Tick.Close = fmt.Sprintf("%v", close)
	data.Tick.Low = ""
	data.Tick.High = ""
	data.Tick.Amount = "0"
	strData, err := json.Marshal(data) //转换成JSON返回的是byte[]
	if err != nil {
		logger.Error(err)
		return ""
	}
	return string(strData)
}

// LogData 存入 kline_log 的数据
func LogData(AssetsCode string, data response.SubIndexKLineResponse) models.KlineLog {
	var klineLog models.KlineLog
	klineLog.ContractCode = AssetsCode                              // 合约代码 BTC-USD
	klineLog.Ch = fmt.Sprintf("market.%v.index.1min", AssetsCode)   // kilne返回日期时间戳
	klineLog.Ts = data.Ts                                           // 自己生成的ts
	klineLog.UpstreamTs = data.Ts                                   // 上游的ts
	klineLog.ID = data.Ts                                           // 自己生成的ts/1000
	klineLog.Mrid = 0.00                                            // ID
	klineLog.COpen = 0.00                                           // 起始峰值
	klineLog.CClose = helpers.StrToFloat64(data.Tick.Close)         // 结束峰值
	klineLog.High = 0                                               // 最高峰值
	klineLog.Low = 0                                                // 最低峰值
	klineLog.Amount = 0                                             // 成交金额 TODO
	klineLog.Vol = 0                                                // 成交手 TODO
	klineLog.CCount = 0                                             // 交易笔数 TODO
	klineLog.TsDate = time.Unix(0, data.Ts*int64(time.Millisecond)) // Kline返回日期
	klineLog.CreateDate = time.Now()                                // 创建日期
	klineLog.KlineSource = config.KlineSource                       // K线来源iqoption
	return klineLog
}
