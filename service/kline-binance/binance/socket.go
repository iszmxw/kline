package binance

import (
	"encoding/json"
	"errors"
	"fmt"
	"goapi/app/models"
	"goapi/app/response"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	mail_notice "goapi/pkg/notice/mail"
	"goapi/pkg/redis"
	"goapi/service/kline-binance/binance/api"
	"goapi/service/kline-binance/binance/config"
	"goapi/service/kline-binance/binance/wsdata"
	kline_common "goapi/service/kline-common"
	"strings"
	"time"
)

func SocketBinance(info kline_common.Info) {
	callback := make(chan struct{}, 1)
	// 开始订阅
	{
		streamName := fmt.Sprintf("%s@indexPriceKline_%v", strings.ToLower(info.Symbol), "1m")
		wsmkClient := new(api.WSPriceClient).Init("dstream.binance.com", streamName, kline_common.GetProxy(config.KlineSource), callback)
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
			runStr := time.Now().UnixMilli() // 开始时间
			// RequestId 定义为 Ts 时间戳
			logger.RequestId = fmt.Sprintf("%v", resp.E1)
			// bian网推送的回来的数据转换为字符串
			strData, err := json.Marshal(resp) //转换成JSON返回的是byte[]
			if err != nil {
				logger.Error(err)
				return
			}
			if len(strData) <= 0 {
				logger.Error(errors.New(fmt.Sprintf(
					"订阅失败 当前订阅的 topic 为：%v,返回的数据为：%v",
					resp.Ps, resp)))
				return
			}
			logger.Info(resp)
			ContractCode := info.ContractCode
			resp.Ps = info.ContractCode
			SystemTs := time.Now().UnixMilli() // 系統生成的Ts
			UpstreamTs := resp.E1              // 上游的ts
			var data response.SubIndexKLineResponse
			data.Ch = fmt.Sprintf("market.%v.index.1min", ContractCode)
			data.Ts = SystemTs                                                                             // 系統生成的Ts
			data.TsDate = time.Unix(0, SystemTs*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999") // 系統生成的Ts
			data.Source = config.KlineSource
			data.Tick.Id = SystemTs
			data.Tick.Vol = "0"
			data.Tick.Count = 0
			data.Tick.Open = resp.K.O
			data.Tick.Close = fmt.Sprintf("%v", resp.K.C)
			data.Tick.Low = resp.K.L1
			data.Tick.High = resp.K.H
			data.Tick.Amount = "0"
			tips := fmt.Sprintf(
				"bian时差：系统时间：%v(%v)<br>bian的 ts 时间：%v(%v)",
				time.Now().UnixMilli(),
				time.Now().Format("2006-01-02 15:04:05"),
				UpstreamTs,
				time.Unix(UpstreamTs/1000, 0).Format("2006-01-02 15:04:05"),
			)
			logger.Info(tips)
			shicha := SystemTs - UpstreamTs
			logger.Info(fmt.Sprintf("shicha=%v", shicha))
			msgData, MarshalJSONErr := json.Marshal(data) //转换成JSON返回的是byte[]
			if MarshalJSONErr != nil {
				logger.Error(MarshalJSONErr)
				return
			}
			// 收集数据缓存到redis
			key := fmt.Sprintf("%v:%v", config.KlineSource, data.Ch)
			_, redisErr := redis.Client.Add(key, msgData, 0)
			if redisErr != nil {
				logger.Error(redisErr)
				return
			}
			// 当前数据源匹配就入库
			if response.KlineSource(ContractCode) == config.KlineSource {
				var KlineData models.KlineLog
				err = mysql.DB.Debug().
					Model(models.KlineLog{}).
					Select("log_id").
					Where("contract_code", ContractCode).
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
				cerr1 := mysql.DB.Debug().Model(models.KlineLog{}).Create(&lineLogRes).Error
				if cerr1 != nil {
					logger.Error(cerr1)
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

			runEnd := time.Now().UnixMilli() // 开始时间
			// 计算一次存储数据消耗的时间
			logger.Info(fmt.Sprintf(
				"开始时间：%v(%v),结束时间%v(%v),消耗%vms",
				runStr,
				time.Unix(runStr/1000, 0).Format("2006-01-02 15:04:05"),
				runEnd,
				time.Unix(runEnd/1000, 0).Format("2006-01-02 15:04:05"),
				runEnd-runStr,
			))

		}, "")
	}
	select {
	case <-callback:
		// 重启服务
		kline_common.RebootSocketLog(config.KlineSource, info)
		go SocketBinance(info)
	}
}

func LogData(resp *wsdata.SubPriceKLineResponse, SystemTs, UpstreamTs int64) models.KlineLog {
	var klineLog models.KlineLog
	klineLog.ContractCode = strings.ReplaceAll(strings.ToUpper(resp.Ps), "/", "-") // 合约代码 BTC-USD
	klineLog.Ch = fmt.Sprintf("market.%v.index.1min", klineLog.ContractCode)       // kilne返回日期时间戳
	klineLog.Ts = SystemTs                                                         // 自己生成的ts
	klineLog.UpstreamTs = UpstreamTs                                               // 上游的ts
	klineLog.ID = SystemTs                                                         // 自己生成的ts/1000
	klineLog.Mrid = 0.00                                                           // ID
	klineLog.COpen = helpers.StrToFloat64(resp.K.O)                                // 起始峰值
	klineLog.CClose = helpers.StrToFloat64(resp.K.C)                               // 结束峰值
	klineLog.High = helpers.StrToFloat64(resp.K.H)                                 // 最高峰值
	klineLog.Low = helpers.StrToFloat64(resp.K.L1)                                 // 最低峰值
	klineLog.Amount = 0                                                            // 成交金额 TODO
	klineLog.Vol = 0                                                               // 成交手 TODO
	klineLog.CCount = 0                                                            // 交易笔数 TODO
	klineLog.TsDate = time.Unix(0, SystemTs*int64(time.Millisecond))               // Kline返回日期
	klineLog.CreateDate = time.Now()                                               // 创建日期
	klineLog.KlineSource = config.KlineSource                                      // K线来源binance
	return klineLog
}
