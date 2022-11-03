package qxbroker

import (
	"encoding/json"
	"fmt"
	goRedis "github.com/go-redis/redis"
	"goapi/app/models"
	"goapi/app/response"
	"goapi/pkg/kline/otc"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/notice/mail"
	"goapi/pkg/redis"
	"goapi/service/kline-common"
	"goapi/service/kline-qxbroker/qxbroker/api"
	"goapi/service/kline-qxbroker/qxbroker/config"
	"goapi/service/kline-qxbroker/qxbroker/wsdata"
	"strings"
	"time"
)

func SocketQxbroker(info KlineCommon.Info) {
	//  3 定时检测数据健康状况，不健康通知外部
	//  2 上游直接断开了，通知外部已断开
	callback := make(chan int, 1)
	// OTC 数据源读取时间段
	if otc.QxbrokerCheckOTC() && !strings.Contains(info.Symbol, "_otc") {
		info.Symbol = info.Symbol + "_otc"
	}
	// 开始订阅
	{
		logger.Info(info)
		wsClient := new(api.WSPriceClient).Init("", KlineCommon.GetProxy(config.KlineSource), callback)
		wsClient.SubPriceKLine(info, func(resp *wsdata.SubPriceKLineResponse) {
			SystemTs := time.Now().UnixMilli()                                   // 系統生成的Ts
			logger.RequestId = fmt.Sprintf("%v@%v", SystemTs, info.ContractCode) // RequestId 定义为 Ts 时间戳
			resp.Symbol = info.ContractCode
			defer func() {
				if r := recover(); r != nil {
					NoticeMail.Notice("msgServiceAnomaly", map[string]interface{}{
						"name": config.KlineSource,
						"r":    r,
					})
					// 关闭
					callback <- 1
				}
			}()
			if resp.TimeId == 1 {
				// 数据源休市控制权，以及状态维护
				//
				//if response.KlineSource(info.ContractCode) == config.KlineSource {
				//	go kline.UpdateSleep(config.KlineSource, info.ContractCode)
				//}
				// 关闭
				callback <- 1
				return
			}
			UpstreamTs := int64(resp.TimeId * 1000) // 上游的ts
			// TODO::实时计算
			// 一分钟、五分钟、十分钟、十五分钟
			// ZADD "INRRank:2022-10-27" 17918.75 "https://fileserver.iqoptionproo.com/file/download/0a3d5de4-2eb2-4001-93ea-f7c051ec3e5e.png=INR=3ED16C397B5"
			minute := time.Unix(0, SystemTs*int64(time.Millisecond)).Format("2006-01-02 15:04")
			// 最大最小
			MaxAndMin := fmt.Sprintf("MaxAndMin:%v:%v", info.ContractCode, minute)
			redis.Client.ZAdd(MaxAndMin, goRedis.Z{
				Score:  resp.Close,
				Member: SystemTs,
			})
			// 开始结束
			OpenAndClose := fmt.Sprintf("OpenAndClose:%v:%v", info.ContractCode, minute)
			redis.Client.ZAdd(OpenAndClose, goRedis.Z{
				Score:  float64(SystemTs),
				Member: resp.Close,
			})
			redis.Client.Expire(MaxAndMin, 180)
			redis.Client.Expire(OpenAndClose, 180)
			// TODO::实时计算
			//TimeDiff := SystemTs - UpstreamTs
			//logger.Info(fmt.Sprintf("ShiCha=%v", TimeDiff)) // 时差提示
			//logger.Info(KlineCommon.TimeDiffTips(config.KlineSource, UpstreamTs))
			data := RedisData(SystemTs, UpstreamTs, resp)
			//logger.Info(data)
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
					Where("contract_code", resp.Symbol).
					Where("upstream_ts", UpstreamTs).
					Where("kline_source", config.KlineSource).
					Find(&KlineData).Error
				if err != nil {
					logger.Error(err)
				}
				if KlineData.LogId > 0 {
					return
				}
				lineLogRes := LogData(resp, SystemTs, UpstreamTs)
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
			// 计算一次存储数据消耗的时间
			//logger.Info(KlineCommon.ConsumeTime(SystemTs, time.Now().UnixMilli()))
		}, "")
	}
	select {
	case msg, ok := <-callback:
		if !ok {
			logger.Info(ok)
			return
		}
		logger.Info(msg)
		// 重启服务
		KlineCommon.RebootSocketLog(config.KlineSource, info)
		go SocketQxbroker(info)
	}
}

// LogData 存入 kline_log 的数据
func LogData(resp *wsdata.SubPriceKLineResponse, SystemTs, UpstreamTs int64) models.KlineLog {
	var klineLog models.KlineLog
	klineLog.ContractCode = strings.ReplaceAll(resp.Symbol, "/", "-")              // 合约代码 BTC-USD
	klineLog.Ch = fmt.Sprintf("market.%v.index.1min", klineLog.ContractCode)       // kilne返回日期时间戳
	klineLog.Ts = SystemTs                                                         // 自己生成的ts
	klineLog.UpstreamTs = UpstreamTs                                               // 上游的ts
	klineLog.ID = SystemTs                                                         // 自己生成的ts/1000
	klineLog.Mrid = 0.00                                                           // ID
	klineLog.COpen = 0                                                             // 起始峰值
	klineLog.CClose = resp.Close                                                   // 结束峰值
	klineLog.High = 0                                                              // 最高峰值
	klineLog.Low = 0                                                               // 最低峰值
	klineLog.Amount = 0                                                            // 成交金额 TODO
	klineLog.Vol = 0                                                               // 成交手 TODO
	klineLog.CCount = 0                                                            // 交易笔数 TODO
	klineLog.TsDate = time.Unix(0, time.Now().UnixMilli()*int64(time.Millisecond)) // Kline返回日期
	klineLog.CreateDate = time.Now()                                               // 创建日期
	klineLog.KlineSource = config.KlineSource                                      // K线来源qxbroker
	return klineLog
}

// RedisData 缓存进redis的数据
func RedisData(SystemTs, UpstreamTs int64, resp *wsdata.SubPriceKLineResponse) response.SubIndexKLineResponse {
	var data response.SubIndexKLineResponse
	data.Ch = fmt.Sprintf("market.%v.index.1min", resp.Symbol)
	data.Ts = SystemTs                                                                             // 系統生成的Ts
	data.TsDate = time.Unix(0, SystemTs*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999") // 系統生成的Ts
	data.Source = config.KlineSource
	data.Tick.Id = UpstreamTs
	data.Tick.Vol = ""
	data.Tick.Count = 0
	data.Tick.Open = "0"
	data.Tick.Close = fmt.Sprintf("%v", resp.Close)
	data.Tick.Low = ""
	data.Tick.High = ""
	data.Tick.Amount = "0"
	return data
}
