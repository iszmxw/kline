package olymptrade

import (
	"encoding/json"
	"errors"
	"fmt"
	"goapi/app/models"
	"goapi/app/response"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	mail_notice "goapi/pkg/notice/mail"
	"goapi/pkg/redis"
	kline_common "goapi/service/kline-common"
	"goapi/service/kline-olymptrade/olymptrade/api"
	"goapi/service/kline-olymptrade/olymptrade/config"
	"goapi/service/kline-olymptrade/olymptrade/wsdata"
	"time"
)

func SocketOlymptrade(info kline_common.Info) {
	callback := make(chan struct{}, 1)
	var TradingAssets models.TradingAssets
	var StartSub = true
	errs := mysql.DB.Debug().Model(models.TradingAssets{}).
		Where("assets_code", info.ContractCode).
		//Where("assets_status", "1"). // 资产状态 1正常 2休市 3下架
		Select("id,assets_code,assets_status").
		Find(&TradingAssets).Error
	if errs != nil {
		logger.Error(errs)
		time.Sleep(time.Second) // 休息1秒
		StartSub = false
		// 关闭
		callback <- struct{}{}
	}
	// 资产不存在或者休市状态
	if TradingAssets.Id <= 0 || TradingAssets.AssetsStatus != 1 {
		logger.Info(TradingAssets)
		time.Sleep(time.Second * 2) // 休息 2 秒
		StartSub = false
		// 关闭
		callback <- struct{}{}
	}
	// 开始订阅
	if StartSub {
		// 必须传入登录session
		_, session := new(api.SocketSession).GetToken(false)
		for len(session) == 0 {
			// 不断获取 session ，直到取到为止
			_, session = new(api.SocketSession).GetToken(false)
			time.Sleep(time.Second * 3) // 休息 3 秒
		}
		time.Sleep(time.Second * 5) // 休息 5 秒
		wsmkClient := new(api.WSPriceClient).Init("", session, kline_common.GetProxy(config.KlineSource), callback)
		wsmkClient.SubPriceKLine(info, func(resp *wsdata.SubPriceKLineResponse) {
			defer func() {
				if r := recover(); r != nil {
					mail_notice.Notice("msgServiceAnomaly", map[string]interface{}{
						"name": config.KlineSource,
						"r":    r,
					})
					// 关闭
					callback <- struct{}{}
				}
			}()
			wsData := *resp
			dType := wsData[0].E    // 读取E的类型
			rData := wsData[0].D[0] // {"p":"USDJPY","t":1650526129,"q":0,"open":127.944,"low":127.944,"high":127.944,"close":127.944}
			if dType == 1 {
				return
			}
			SystemTs := time.Now().UnixMilli()             // 开始时间 系統生成的Ts
			logger.RequestId = fmt.Sprintf("%v", SystemTs) // RequestId 定义为 Ts 时间戳
			if rData.T <= 0 {
				// 当前数据源匹配就入库
				if response.KlineSource(info.ContractCode) == config.KlineSource {
					logger.Info("休市")
				}
				logger.Error(
					errors.New(fmt.Sprintf("订阅失败 当前订阅的 topic 为：%v,返回的数据时间为：%v", info.ContractCode, resp)),
				)
				return
			}
			UpstreamTs := int64(rData.T * 1000) // 上游的ts
			TimeDiff := SystemTs - UpstreamTs
			logger.Info(fmt.Sprintf("ShiCha=%v", TimeDiff)) // 时差提示
			logger.Info(kline_common.TimeDiffTips(config.KlineSource, UpstreamTs))
			data := RedisData(SystemTs, resp)
			logger.Info(data)
			redisData, MarshalJSONErr := json.Marshal(data) //转换成JSON返回的是byte[]
			if MarshalJSONErr != nil {
				logger.Error(MarshalJSONErr)
				return
			}
			// 收集数据缓存到redis
			key := fmt.Sprintf("%v:%v", config.KlineSource, data.Ch)
			_, redisErr := redis.Client.Add(key, redisData, 0)
			if redisErr != nil {
				logger.Error(redisErr)
				return
			}

			// 当前数据源匹配就入库
			if response.KlineSource(info.ContractCode) == config.KlineSource {
				var KlineData models.KlineLog
				err := mysql.DB.Debug().
					Model(models.KlineLog{}).
					Select("log_id").
					Where("contract_code", info.ContractCode).
					Where("upstream_ts", UpstreamTs).
					Where("kline_source", config.KlineSource).
					Find(&KlineData).Error
				if err != nil {
					logger.Error(err)
				}
				if KlineData.LogId > 0 {
					return
				}
				lineLogRes := LogData(wsData, info.ContractCode, SystemTs, UpstreamTs)
				if err = mysql.DB.Debug().Model(models.KlineLog{}).Create(&lineLogRes).Error; err != nil {
					logger.Error(err)
					return
				}
			}

			//// 历史存入 elastic
			//putRes := esdk.IndexPut("kline_log", helpers.Int64ToString(KlineData.LogId), KlineData)
			//if !putRes {
			//	logger.Error(errors.New("历史存入 elastic 失败"))
			//} else {
			//	logger.Info("elastic-success")
			//}
			// 计算一次存储数据消耗的时间
			logger.Info(kline_common.ConsumeTime(SystemTs, time.Now().UnixMilli()))
		}, "")
	}
	select {
	case <-callback:
		// 重启服务
		kline_common.RebootSocketLog(config.KlineSource, info)
		go SocketOlymptrade(info)
	}
}

func LogData(wsData wsdata.SubPriceKLineResponse, ContractCode string, SystemTs, UpstreamTs int64) models.KlineLog {
	var klineLog models.KlineLog
	rData := wsData[0].D[0]                                                        // {"p":"USDJPY","t":1650526129,"q":0,"open":127.944,"low":127.944,"high":127.944,"close":127.944}
	klineLog.ContractCode = ContractCode                                           // 合约代码 BTC-USD
	klineLog.Ch = fmt.Sprintf("market.%v.index.1min", klineLog.ContractCode)       // kilne返回日期时间戳
	klineLog.Ts = SystemTs                                                         // 自己生成的ts
	klineLog.UpstreamTs = UpstreamTs                                               // 上游的ts
	klineLog.ID = SystemTs                                                         // 自己生成的ts/1000
	klineLog.Mrid = 0.00                                                           // ID
	klineLog.COpen = 0                                                             // 起始峰值
	klineLog.CClose = rData.Close                                                  // 结束峰值
	klineLog.High = rData.High                                                     // 最高峰值
	klineLog.Low = rData.Low                                                       // 最低峰值
	klineLog.Amount = 0                                                            // 成交金额 TODO
	klineLog.Vol = 0                                                               // 成交手 TODO
	klineLog.CCount = 0                                                            // 交易笔数 TODO
	klineLog.TsDate = time.Unix(0, time.Now().UnixMilli()*int64(time.Millisecond)) // Kline返回日期
	klineLog.CreateDate = time.Now()                                               // 创建日期
	klineLog.KlineSource = config.KlineSource                                      // K线来源olymptrade
	return klineLog
}

// RedisData 缓存进redis的数据
func RedisData(SystemTs int64, resp *wsdata.SubPriceKLineResponse) response.SubIndexKLineResponse {
	var data response.SubIndexKLineResponse
	wsData := *resp
	rData := wsData[0].D[0] // {"p":"USDJPY","t":1650526129,"q":0,"open":127.944,"low":127.944,"high":127.944,"close":127.944}
	ContractCode := kline_common.Convert(config.KlineSource, rData.P).ContractCode
	data.Ch = fmt.Sprintf("market.%v.index.1min", ContractCode)
	data.Ts = SystemTs
	data.TsDate = time.Unix(0, SystemTs*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999") // 系統生成的Ts
	data.Source = config.KlineSource
	data.Tick.Id = SystemTs
	data.Tick.Vol = "0"
	data.Tick.Count = 0
	data.Tick.Open = fmt.Sprintf("%v", rData.Open)
	data.Tick.Close = fmt.Sprintf("%v", rData.Close)
	data.Tick.Low = fmt.Sprintf("%v", rData.Low)
	data.Tick.High = fmt.Sprintf("%v", rData.High)
	data.Tick.Amount = "0"
	return data
}
