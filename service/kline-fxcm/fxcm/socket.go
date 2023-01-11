package fxcm

import (
	"encoding/json"
	"fmt"
	socketio "github.com/graarh/golang-socketio"
	"goapi/app/models"
	"goapi/app/response"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	NoticeMail "goapi/pkg/notice/mail"
	"goapi/pkg/redis"
	"goapi/service/kline-common"
	"goapi/service/kline-common/source"
	"goapi/service/kline-fxcm/fxcm/api"
	"goapi/service/kline-fxcm/fxcm/config"
	"goapi/service/kline-fxcm/fxcm/wsdata"
	"strings"
	"time"
)

func SocketFxcm(info KlineCommon.Info) {
	wsClient := new(api.FXCM)
	wsClient.CloseChanCallback = make(chan KlineCommon.Info, 20)
	if err := wsClient.Start(); err != nil {
		logger.Info(err)
		go SocketFxcm(info)
		return
	}
	// 开启定时器30秒检测一次数据活跃状态
	wsClient.OnceReboot.Do(func() {
		defer KlineCommon.KlineRecover("TickerReboot", config.KlineSource, info.Symbol)
		ticker := time.NewTicker(30 * time.Second) // 每隔 30 秒检查一次
		KeyStatus := fmt.Sprintf("status:%v:market.%v.index.1min", config.KlineSource, info.ContractCode)
		for range ticker.C {
			status := source.CheckStatus(config.KlineSource, info.ContractCode)
			if status == "4" { // 延迟过高关闭重启 status:qxbroker:market.Microsoft.index.1min
				logger.Info(KeyStatus + "当前数据不健康正在重启中" + info.Symbol)
				goto STOP
			}
		}
	STOP:
		// 关闭
		ticker.Stop()
		wsClient.CloseChanCallback <- info
	})

	// 开始订阅
	{
		wsClient.GetTicker(info)
		// 接收数据
		if wsClient.Socket == nil {
			go SocketFxcm(info)
			return
		}
		_ = wsClient.Socket.On(info.Symbol, func(c *socketio.Channel, args interface{}) {
			var stream wsdata.DataStream
			err := json.Unmarshal([]byte(args.(string)), &stream)
			if err != nil {
				logger.Error(err)
				wsClient.CloseChanCallback <- info
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
					wsClient.CloseChanCallback <- info
				}
			}()

			UpstreamTs := stream.Updated // 上游的ts
			TimeDiff := SystemTs - UpstreamTs
			logger.Info(fmt.Sprintf("ShiCha=%v", TimeDiff)) // 时差提示
			logger.Info(KlineCommon.TimeDiffTips(config.KlineSource, UpstreamTs))
			data := RedisData(SystemTs, stream)
			data.Ch = fmt.Sprintf("market.%v.index.1min", info.ContractCode)
			redisData, err := json.Marshal(data) //转换成JSON返回的是byte[]
			if err != nil {
				logger.Error(err)
				wsClient.CloseChanCallback <- info
				return
			}
			// 收集数据缓存到redis
			key := fmt.Sprintf("%v:%v", config.KlineSource, data.Ch)
			logger.Info("0001"+data.Ch, data)
			_, redisErr := redis.Client.Add(key, redisData, 0)
			if redisErr != nil {
				logger.Error(redisErr)
				wsClient.CloseChanCallback <- info
				return
			}
			// 当前数据源匹配就入库
			if response.KlineSource(info.ContractCode) == config.KlineSource {
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
					return
				}
				lineLogRes := LogData(stream, SystemTs)
				if err = mysql.DB.Model(models.KlineLog{}).Create(&lineLogRes).Error; err != nil {
					logger.Error(err)
					return
				}
				logger.Info("0003"+data.Ch, lineLogRes)
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
	case Symbol, ok := <-wsClient.CloseChanCallback:
		if ok {
			// 重启服务
			logger.Info("重启", Symbol)
			KlineCommon.RebootSocketLog(config.KlineSource, Symbol)
			go SocketFxcm(Symbol)
		} else {
			logger.Info("!ok", "重启服务失败")
		}
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
	Close, err := helpers.FormatFloat(klineLog.CClose, 10)
	if err != nil {
		klineLog.CClose = 0.00
	} else {
		klineLog.CClose = Close
	}
	klineLog.High = stream.Rates[2]                                                // 最高峰值
	klineLog.Low = stream.Rates[3]                                                 // 最低峰值
	klineLog.Amount = 0                                                            // 成交金额 TODO
	klineLog.Vol = stream.Rates[0]                                                 // 成交手 TODO
	klineLog.CCount = 0                                                            // 交易笔数 TODO
	klineLog.TsDate = time.Unix(0, time.Now().UnixMilli()*int64(time.Millisecond)) // Kline返回日期
	klineLog.CreateDate = time.Now()                                               // 创建日期
	klineLog.KlineSource = config.KlineSource                                      // K线来源raceoption
	return klineLog
}

// RedisData 缓存进redis的数据
func RedisData(SystemTs int64, stream wsdata.DataStream) response.SubIndexKLineResponse {
	var data response.SubIndexKLineResponse
	data.Ts = SystemTs                                                                             // 系統生成的Ts
	data.TsDate = time.Unix(0, SystemTs*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999") // 系統生成的Ts
	data.Source = config.KlineSource
	data.Tick.Id = stream.Updated
	data.Tick.Vol = fmt.Sprintf("%v", stream.Rates[0])
	data.Tick.Count = 0
	data.Tick.Open = fmt.Sprintf("%v", stream.Rates[1])
	data.Tick.Close = fmt.Sprintf("%v", (stream.Rates[0]+stream.Rates[1])/2)
	Close, err := helpers.FormatFloat(helpers.StrToFloat64(data.Tick.Close), 10)
	if err != nil {
		data.Tick.Close = "0.00"
	} else {
		data.Tick.Close = fmt.Sprintf("%v", Close)
	}
	data.Tick.Low = fmt.Sprintf("%v", stream.Rates[3])
	data.Tick.High = fmt.Sprintf("%v", stream.Rates[2])
	data.Tick.Amount = "0"
	return data
}
