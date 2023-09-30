package binomo

import (
	"encoding/json"
	"fmt"
	"goapi/app/models"
	"goapi/app/response"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/notice/mail"
	"goapi/pkg/redis"
	"goapi/service/kline-binomo/binomo/api"
	"goapi/service/kline-binomo/binomo/config"
	"goapi/service/kline-binomo/binomo/wsdata"
	"goapi/service/kline-common"
	"sync"
	"time"
)

func Init() *api.WSPriceClient {
	api.Once.Do(func() {
		api.App = new(api.WSPriceClient)
		wc := api.WsInit()
		wc.Callback = make(chan int, 1)
		// 启动WebSocket连接
		wc.WsConnect(KlineCommon.GetProxy(config.KlineSource))
		api.App.WsClient = wc
	})
	fmt.Println(api.App)
	return api.App
}

func SocketBinomo(info KlineCommon.Info) {
	App := Init()
	//  3 定时检测数据健康状况，不健康通知外部
	// 开始订阅
	{
		logger.Info(info)
		App.SubPriceKLine(info, func(resp *wsdata.SubPriceKLineResponse) {
			SystemTs := time.Now().UnixMilli()                                   // 系統生成的Ts
			logger.RequestId = fmt.Sprintf("%v@%v", SystemTs, info.ContractCode) // RequestId 定义为 Ts 时间戳
			defer func() {
				if r := recover(); r != nil {
					NoticeMail.Notice("msgServiceAnomaly", map[string]interface{}{
						"name": config.KlineSource,
						"r":    r,
					})
					// 关闭
					App.WsClient.Callback <- 1
				}
			}()
			UpstreamTs := resp.CreatedAt.UnixMilli() // 上游的ts
			l := App.List.Back()
			if l != nil {
				App.List.Remove(l)
				if l.Value == UpstreamTs {
					logger.Info("相等了")
					return
				}
			}
			App.List.PushFront(UpstreamTs)
			TimeDiff := SystemTs - UpstreamTs
			if TimeDiff >= 30*1000 {
				logger.Info("时差大于30秒直接不存储，认为数据源不健康")
				logger.Info(fmt.Sprintf("ShiCha=%v", TimeDiff)) // 时差提示
				logger.Info(KlineCommon.TimeDiffTips(config.KlineSource, UpstreamTs))
				// 时差大于30秒直接不存储，认为数据源不健康
				return
			}
			data := RedisData(SystemTs, UpstreamTs, resp)
			data.Ch = fmt.Sprintf("market.%v.index.1min", info.ContractCode)
			logger.Info(data)
			redisData, err := json.Marshal(data) //转换成JSON返回的是byte[]
			if err != nil {
				logger.Error(err)
				return
			}
			// 收集数据缓存到redis
			key := fmt.Sprintf("%v:%v", config.KlineSource, data.Ch)
			logger.Info("0001"+data.Ch, data)
			_, redisErr := redis.Client.Add(key, redisData, 0)
			if redisErr != nil {
				logger.Error(redisErr)
				return
			}
			// 当前数据源匹配就入库
			KlineSource := response.KlineSource(info.ContractCode)
			logger.Info(KlineSource, info.ContractCode)
			if KlineSource == config.KlineSource {
				var KlineData models.KlineLog
				err = mysql.DB.Debug().
					Model(models.KlineLog{}).
					Select("log_id").
					Where("contract_code", info.ContractCode).
					Where("ts", SystemTs).
					Where("upstream_ts", UpstreamTs).
					Where("c_close", data.Tick.Close).
					Where("kline_source", config.KlineSource).
					Find(&KlineData).Error
				if err != nil {
					logger.Error(err)
				}
				logger.Info("0002"+data.Ch, KlineData)
				if KlineData.LogId > 0 {
					logger.Info("0004", "已存在")
					return
				}
				lineLogRes := LogData(resp, info.ContractCode, SystemTs, UpstreamTs)
				if err = mysql.DB.Model(models.KlineLog{}).Create(&lineLogRes).Error; err != nil {
					logger.Error(err)
					return
				}
				logger.Info("0003"+data.Ch, lineLogRes)
			}
			// 插针后的数据缓存到另外一张表
			RecoverSource, err := redis.Client.Get(fmt.Sprintf("recover_source:%v", info.ContractCode))
			if err != nil && err.Error() != "redis: nil" {
				logger.Error(err)
				return
			}
			if RecoverSource == config.KlineSource {
				var KlineData models.KlineRecoverLog
				err = mysql.DB.Debug().
					Model(models.KlineRecoverLog{}).
					Select("log_id").
					Where("contract_code", info.ContractCode).
					Where("ts", SystemTs).
					Where("upstream_ts", UpstreamTs).
					Where("c_close", data.Tick.Close).
					Where("kline_source", config.KlineSource).
					Find(&KlineData).Error
				if err != nil {
					logger.Error(err)
				}
				logger.Info("0002"+data.Ch, KlineData)
				if KlineData.LogId > 0 {
					logger.Info("0004", "已存在")
					return
				}
				lineLogRes := LogData(resp, info.ContractCode, SystemTs, UpstreamTs)
				if err = mysql.DB.Model(models.KlineRecoverLog{}).Create(&lineLogRes).Error; err != nil {
					logger.Error(err)
					return
				}
				logger.Info("0003"+data.Ch, lineLogRes)
			}
			logger.Info(KlineCommon.ConsumeTime(SystemTs, time.Now().UnixMilli()))
		})
	}
	select {
	case msg, ok := <-App.WsClient.Callback:
		if !ok {
			logger.Info(ok)
			return
		}
		logger.Info(msg)
		// 重启服务
		KlineCommon.RebootSocketLog(config.KlineSource, info)
		api.Once = sync.Once{} // 创建一个新的 sync.Once 实例
		go SocketBinomo(info)
	}
}

// LogData 存入 kline_log 的数据
func LogData(resp *wsdata.SubPriceKLineResponse, ContractCode string, SystemTs, UpstreamTs int64) models.KlineLog {
	var klineLog models.KlineLog
	klineLog.ContractCode = ContractCode                            // 合约代码 BTC-USD
	klineLog.Ch = fmt.Sprintf("market.%v.index.1min", ContractCode) // kilne返回日期时间戳
	klineLog.Ts = SystemTs                                          // 自己生成的ts
	klineLog.UpstreamTs = UpstreamTs                                // 上游的ts
	klineLog.ID = SystemTs                                          // 自己生成的ts/1000
	klineLog.Mrid = 0.00                                            // ID
	klineLog.COpen = resp.Ask                                       // 起始峰值
	klineLog.CClose = 0                                             // 结束峰值
	Close, err := helpers.FormatFloat(helpers.StrToFloat64(fmt.Sprintf("%v", klineLog.CClose)), 10)
	if err != nil {
		klineLog.CClose = 0.00
	} else {
		klineLog.CClose = Close
	}
	klineLog.High = resp.Ask                                                       // 最高峰值
	klineLog.Low = resp.Ask                                                        // 最低峰值
	klineLog.Amount = 0                                                            // 成交金额 TODO
	klineLog.Vol = 0                                                               // 成交手 TODO
	klineLog.CCount = 0                                                            // 交易笔数 TODO
	klineLog.TsDate = time.Unix(0, time.Now().UnixMilli()*int64(time.Millisecond)) // Kline返回日期
	klineLog.CreateDate = time.Now()                                               // 创建日期
	klineLog.KlineSource = config.KlineSource                                      // K线来源iqoption
	return klineLog
}

// RedisData 缓存进redis的数据
func RedisData(SystemTs, UpstreamTs int64, resp *wsdata.SubPriceKLineResponse) response.SubIndexKLineResponse {
	var data response.SubIndexKLineResponse
	data.Ts = SystemTs                                                                             // 系統生成的Ts
	data.TsDate = time.Unix(0, SystemTs*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999") // 系統生成的Ts
	data.Source = config.KlineSource
	data.Tick.Id = UpstreamTs
	data.Tick.Vol = ""
	data.Tick.Count = 0
	data.Tick.Open = fmt.Sprintf("%v", resp.Ask)
	Close, err := helpers.FormatFloat(helpers.StrToFloat64(fmt.Sprintf("%v", resp.Ask)), 10)
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
