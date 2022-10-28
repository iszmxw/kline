package fxcm

import (
	"encoding/json"
	"fmt"
	socketio "github.com/graarh/golang-socketio"
	"goapi/app/models"
	"goapi/app/response"
	"goapi/pkg/kline/otc"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	NoticeMail "goapi/pkg/notice/mail"
	"goapi/pkg/redis"
	"goapi/service/kline-common"
	"goapi/service/kline-fxcm/fxcm/api"
	"goapi/service/kline-fxcm/fxcm/config"
	"goapi/service/kline-fxcm/fxcm/wsdata"
	"strings"
	"time"
)

func SocketFxcm(wsClient *api.FXCM, info KlineCommon.Info) {
	callback := make(chan struct{}, 1)
	// OTC 数据源读取时间段
	if otc.RaceOptionCheckOTC() {
		time.Sleep(time.Second) // 休息1秒
		// 休市切换 OTC 数据源读取时间段
		if !strings.Contains(info.Symbol, ".W") {
			info.Symbol = info.Symbol + ".W"
		}
		logger.Info("OTC时间")
	} else {
		if strings.Contains(info.Symbol, ".W") {
			info.Symbol = strings.ReplaceAll(info.Symbol, ".W", "") // 去除掉
		}
		logger.Info("正常时间")
	}
	// 开始订阅
	{
		wsClient.GetTicker(strings.ReplaceAll(info.Symbol, "/", "%2F"))
		// 接收数据
		_ = wsClient.Socket.On(info.Symbol, func(c *socketio.Channel, args interface{}) {
			var stream wsdata.DataStream
			err := json.Unmarshal([]byte(args.(string)), &stream)
			if err != nil {
				logger.Error(err)
				return
			}
			logger.Info(stream)
			SystemTs := time.Now().UnixMilli() // 开始时间 系統生成的Ts
			// RequestId 定义为 Ts 时间戳
			logger.RequestId = fmt.Sprintf("%v", SystemTs)
			defer func() {
				if r := recover(); r != nil {
					NoticeMail.Notice("msgServiceAnomaly", map[string]interface{}{
						"name": config.KlineSource,
						"r":    r,
					})
					// 关闭
					callback <- struct{}{}
				}
			}()

			UpstreamTs := stream.Updated // 上游的ts
			TimeDiff := SystemTs - UpstreamTs
			logger.Info(fmt.Sprintf("ShiCha=%v", TimeDiff)) // 时差提示
			logger.Info(KlineCommon.TimeDiffTips(config.KlineSource, UpstreamTs))
			data := RedisData(SystemTs, stream)
			logger.Info(data)
			redisData, err := json.Marshal(data) //转换成JSON返回的是byte[]
			if err != nil {
				logger.Error(err)
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
				err = mysql.DB.
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
				lineLogRes := LogData(stream, SystemTs)
				if err = mysql.DB.Model(models.KlineLog{}).Create(&lineLogRes).Error; err != nil {
					logger.Error(err)
					return
				}
			}

			//// 历史存入 elastic todo::暂时不开启
			//putRes := esdk.IndexPut("kline_log", helpers.Int64ToString(KlineData.LogId), KlineData)
			//if !putRes {
			//	logger.Error(errors.New("历史存入 elastic 失败"))
			//} else {
			//	logger.Info("elastic-success")
			//}

		})
	}
	select {
	case <-callback:
		// 重启服务
		KlineCommon.RebootSocketLog(config.KlineSource, info)
		go SocketFxcm(wsClient, info)
	}
}

// LogData 存入 kline_log 的数据
func LogData(stream wsdata.DataStream, SystemTs int64) models.KlineLog {
	var klineLog models.KlineLog
	klineLog.ContractCode = strings.ReplaceAll(strings.ToUpper(stream.Symbol), "/", "-") // 合约代码 BTC-USD
	klineLog.Ch = fmt.Sprintf("market.%v.index.1min", klineLog.ContractCode)             // kilne返回日期时间戳
	klineLog.Ts = SystemTs                                                               // 自己生成的ts
	klineLog.UpstreamTs = stream.Updated                                                 // 上游的ts
	klineLog.ID = SystemTs                                                               // 自己生成的ts/1000
	klineLog.Mrid = 0.00                                                                 // ID
	klineLog.COpen = stream.Rates[1]                                                     // 起始峰值
	klineLog.CClose = (stream.Rates[0] + stream.Rates[1]) / 2                            // 结束峰值
	klineLog.High = stream.Rates[2]                                                      // 最高峰值
	klineLog.Low = stream.Rates[3]                                                       // 最低峰值
	klineLog.Amount = 0                                                                  // 成交金额 TODO
	klineLog.Vol = stream.Rates[0]                                                       // 成交手 TODO
	klineLog.CCount = 0                                                                  // 交易笔数 TODO
	klineLog.TsDate = time.Unix(0, time.Now().UnixMilli()*int64(time.Millisecond))       // Kline返回日期
	klineLog.CreateDate = time.Now()                                                     // 创建日期
	klineLog.KlineSource = config.KlineSource                                            // K线来源raceoption
	return klineLog
}

// RedisData 缓存进redis的数据
func RedisData(SystemTs int64, stream wsdata.DataStream) response.SubIndexKLineResponse {
	var data response.SubIndexKLineResponse
	data.Ch = fmt.Sprintf("market.%v.index.1min", stream.Symbol)
	data.Ts = SystemTs                                                                             // 系統生成的Ts
	data.TsDate = time.Unix(0, SystemTs*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999") // 系統生成的Ts
	data.Source = config.KlineSource
	data.Tick.Id = stream.Updated
	data.Tick.Vol = fmt.Sprintf("%v", stream.Rates[0])
	data.Tick.Count = 0
	data.Tick.Open = fmt.Sprintf("%v", stream.Rates[1])
	data.Tick.Close = fmt.Sprintf("%v", (stream.Rates[0]+stream.Rates[1])/2)
	data.Tick.Low = fmt.Sprintf("%v", stream.Rates[3])
	data.Tick.High = fmt.Sprintf("%v", stream.Rates[2])
	data.Tick.Amount = "0"
	return data
}
