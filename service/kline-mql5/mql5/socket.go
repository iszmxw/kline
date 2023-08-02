package mql5

import (
	"encoding/json"
	"fmt"
	"goapi/app/models"
	"goapi/app/response"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
	KlineCommon "goapi/service/kline-common"
	kconfig "goapi/service/kline-mql5/mql5/config"
	"goapi/service/kline-mql5/mql5/wsdata"
	"time"
)

func SocketMql5(info KlineCommon.Info, mql5 wsdata.Mql5) {
	SystemTs := time.Now().UnixMilli()
	logger.RequestId = fmt.Sprintf("%v@%v", SystemTs, mql5.Name) // RequestId 定义为 Ts 时间戳
	UpstreamTs := mql5.TimeMsc                                   // 上游的ts
	TimeDiff := SystemTs - UpstreamTs
	if TimeDiff >= 30*1000 {
		logger.Info("时差大于30秒直接不存储，认为数据源不健康")
		logger.Info(fmt.Sprintf("ShiCha=%v", TimeDiff)) // 时差提示
		logger.Info(KlineCommon.TimeDiffTips(kconfig.KlineSource, UpstreamTs))
		// 时差大于30秒直接不存储，认为数据源不健康
		return
	}
	data := RedisData(SystemTs, UpstreamTs, mql5)
	data.Ch = fmt.Sprintf("market.%v.index.1min", info.ContractCode)
	redisData, err := json.Marshal(data) //转换成JSON返回的是byte[]
	if err != nil {
		logger.Error(err)
		return
	}
	// 收集数据缓存到redis
	key := fmt.Sprintf("%v:%v", kconfig.KlineSource, data.Ch)
	_, redisErr := redis.Client.Add(key, redisData, 0)
	if redisErr != nil {
		logger.Error(redisErr)
		return
	}
	// 当前数据源匹配就入库
	//KlineSource := response.KlineSource(info.ContractCode)
	if true {
		var KlineData models.KlineLog
		mysql.DB.
			Model(models.KlineLog{}).
			Select("log_id").
			Where("contract_code", info.ContractCode).
			Where("ts", SystemTs).
			Where("upstream_ts", UpstreamTs).
			Where("c_close", data.Tick.Close).
			Where("kline_source", kconfig.KlineSource).
			Find(&KlineData)
		if KlineData.LogId > 0 {
			logger.Info("0004", "已存在")
			return
		}
		lineLogRes := LogData(mql5, info.ContractCode, SystemTs, UpstreamTs)
		if err = mysql.DB.Model(models.KlineLog{}).Create(&lineLogRes).Error; err != nil {
			logger.Error(err)
			return
		}
	}
	// 插针后的数据缓存到另外一张表
	RecoverSource, _ := redis.Client.Get(fmt.Sprintf("recover_source:%v", info.ContractCode))
	if len(RecoverSource) == 0 {
		logger.Info(KlineCommon.ConsumeTime(SystemTs, time.Now().UnixMilli()))
		return
	}
	if kconfig.KlineSource == RecoverSource {
		var KlineData models.KlineRecoverLog
		err = mysql.DB.
			Model(models.KlineRecoverLog{}).
			Select("log_id").
			Where("contract_code", info.ContractCode).
			Where("ts", SystemTs).
			Where("upstream_ts", UpstreamTs).
			Where("c_close", data.Tick.Close).
			Where("kline_source", kconfig.KlineSource).
			Find(&KlineData).Error
		if err != nil {
			logger.Error(err)
		}
		if KlineData.LogId > 0 {
			logger.Info("0004", "已存在")
			return
		}
		lineLogRes := LogData(mql5, info.ContractCode, SystemTs, UpstreamTs)
		if err = mysql.DB.Model(models.KlineRecoverLog{}).Create(&lineLogRes).Error; err != nil {
			logger.Error(err)
			return
		}
	}
	logger.Info(KlineCommon.ConsumeTime(SystemTs, time.Now().UnixMilli()), SystemTs, "---", time.Now().UnixMilli())
}

// LogData 存入 kline_log 的数据
func LogData(mql5 wsdata.Mql5, ContractCode string, SystemTs, UpstreamTs int64) models.KlineLog {
	var klineLog models.KlineLog
	klineLog.ContractCode = ContractCode                            // 合约代码 BTC-USD
	klineLog.Ch = fmt.Sprintf("market.%v.index.1min", ContractCode) // kilne返回日期时间戳
	klineLog.Ts = SystemTs                                          // 自己生成的ts
	klineLog.UpstreamTs = UpstreamTs                                // 上游的ts
	klineLog.ID = SystemTs                                          // 自己生成的ts/1000
	klineLog.Mrid = 0.00                                            // ID
	klineLog.COpen = mql5.Bid                                       // 起始峰值
	klineLog.CClose = mql5.Ask                                      // 结束峰值
	Close, err := helpers.FormatFloat(helpers.StrToFloat64(fmt.Sprintf("%v", klineLog.CClose)), 10)
	if err != nil {
		klineLog.CClose = 0.00
	} else {
		klineLog.CClose = Close
	}
	klineLog.High = 0                                                              // 最高峰值
	klineLog.Low = 0                                                               // 最低峰值
	klineLog.Amount = 0                                                            // 成交金额 TODO
	klineLog.Vol = 0                                                               // 成交手 TODO
	klineLog.CCount = 0                                                            // 交易笔数 TODO
	klineLog.TsDate = time.Unix(0, time.Now().UnixMilli()*int64(time.Millisecond)) // Kline返回日期
	klineLog.CreateDate = time.Now()                                               // 创建日期
	klineLog.KlineSource = kconfig.KlineSource                                     // K线来源iqoption
	return klineLog
}

// RedisData 缓存进redis的数据
func RedisData(SystemTs, UpstreamTs int64, mql5 wsdata.Mql5) response.SubIndexKLineResponse {
	var data response.SubIndexKLineResponse
	data.Ts = SystemTs                                                                             // 系統生成的Ts
	data.TsDate = time.Unix(0, SystemTs*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999") // 系統生成的Ts
	data.Source = kconfig.KlineSource
	data.Tick.Id = UpstreamTs
	data.Tick.Vol = ""
	data.Tick.Count = 0
	data.Tick.Open = fmt.Sprintf("%v", mql5.Bid)
	Close, err := helpers.FormatFloat(mql5.Ask, 10)
	if err != nil {
		data.Tick.Close = "0.00"
	} else {
		data.Tick.Close = fmt.Sprintf("%v", Close)
	}
	data.Tick.Low = ""
	data.Tick.High = ""
	data.Tick.Amount = "0"
	return data
}
