package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"goapi/app/models"
	"goapi/app/response"
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/helpers"
	"goapi/pkg/kline"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	mail_notice "goapi/pkg/notice/mail"
	"goapi/pkg/pprof"
	"goapi/pkg/redis"
	kline_common "goapi/service/kline-common"
	"goapi/service/kline-hsb/hsb/api"
	kconfig "goapi/service/kline-hsb/hsb/config"
	"goapi/service/kline-hsb/hsb/wsdata"
	"math"
	"strings"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init(kconfig.ServiceLog)
	// 初始化 SQL
	bootstrap.SetupDB()
}

func main() {
	// pprof 分析
	pprof.Debug(":6064")
	bootstrap.SetupRedis(15)
	defer bootstrap.RedisClose()
	CodeData := kline_common.GetData(kconfig.KlineSource)
	start := "ok"
	for {
		if start == "ok" {
			logger.Info(CodeData)
			for _, obj := range CodeData {
				logger.Info("缓存 " + obj.Symbol)
				go SocketHsb(obj)
			}
			start = "no"
		}
		time.Sleep(time.Second)
	}
}

func SocketHsb(info kline_common.Info) {
	callback := make(chan struct{}, 1)
	var StartSub = true
	// 资产休市状态
	if kline.CheckSleep(kconfig.KlineSource, info.ContractCode) {
		time.Sleep(time.Second) // 休息1秒
		StartSub = false
		// 关闭
		callback <- struct{}{}
	}
	// 开始订阅
	if StartSub {
		wsClient := new(api.WSPriceClient).Init("", kline_common.GetProxy(kconfig.KlineSource), callback)
		wsClient.SubPriceKLine(info, func(resp *wsdata.SubPriceKLineResponse) {
			defer func() {
				if r := recover(); r != nil {
					mail_notice.Notice("msgServiceAnomaly", map[string]interface{}{
						"name": kconfig.KlineSource,
						"r":    r,
					})
					// 关闭
					callback <- struct{}{}
				}
			}()
			wsData := *resp
			rData := wsData.Pb.Tick            // {"p":"USDJPY","t":1650526129,"q":0,"open":127.944,"low":127.944,"high":127.944,"close":127.944}
			SystemTs := time.Now().UnixMilli() // 开始时间 系統生成的Ts
			// RequestId 定义为 Ts 时间戳
			logger.RequestId = fmt.Sprintf("%v", rData.Time)
			if rData.Time <= 0 {
				// 当前数据源匹配就入库
				if response.KlineSource(info.ContractCode) == kconfig.KlineSource {
					// 休市
				}
				logger.Error(
					errors.New(fmt.Sprintf("订阅失败 当前订阅的 topic 为：%v,返回的数据时间为：%v", info.ContractCode, resp)),
				)
				return
			}
			ContractCode := info.ContractCode
			UpstreamTs := int64(rData.Time * 1000) // 上游的ts
			data := RedisData(SystemTs, resp)
			TimeDiff := SystemTs - UpstreamTs
			logger.Info(data)
			logger.Info(fmt.Sprintf("ShiCha=%v", TimeDiff)) // 时差提示
			logger.Info(kline_common.TimeDiffTips(kconfig.KlineSource, SystemTs))
			redisData, MarshalJSONErr := json.Marshal(data) //转换成JSON返回的是byte[]
			if MarshalJSONErr != nil {
				logger.Error(MarshalJSONErr)
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
			if response.KlineSource(ContractCode) == kconfig.KlineSource {
				var KlineData models.KlineLog
				err := mysql.DB.Debug().
					Model(models.KlineLog{}).
					Select("log_id").
					Where("contract_code", ContractCode).
					Where("upstream_ts", UpstreamTs).
					Where("kline_source", kconfig.KlineSource).
					Find(&KlineData).Error
				if err != nil {
					logger.Error(err)
				}
				if KlineData.LogId > 0 {
					return
				}
				lineLogRes := LogData(wsData, ContractCode, SystemTs, UpstreamTs)
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
		timeKey := fmt.Sprintf("connect:%v:%v:time", kconfig.KlineSource, info.Symbol)
		timeStr, _ := redis.Client.Get(timeKey)
		var lastRuntime int64
		if len(timeStr) > 0 {
			// 上次运行时间
			UnixMilli := helpers.StringToInt64(timeStr)
			// 上次
			lastRuntime = time.Now().UnixMilli() - UnixMilli
		}
		num := redis.Client.Incr(fmt.Sprintf("connect:%v:%v:num", kconfig.KlineSource, info.Symbol))
		_, _ = redis.Client.Add(timeKey, time.Now().UnixMilli(), 0)
		logger.Info(fmt.Sprintf("%v:已关闭正在第【%v】次重新启动Goroutine", info.Symbol, num))
		go mail_notice.Notice("msgRebootSocket", map[string]interface{}{
			"service":     kconfig.KlineSource,
			"symbol":      info.Symbol,
			"num":         num,
			"lastRuntime": lastRuntime,
		})
		go SocketHsb(info)
	}
}

func LogData(wsData wsdata.SubPriceKLineResponse, ContractCode string, SystemTs, UpstreamTs int64) models.KlineLog {
	var klineLog models.KlineLog
	rData := wsData.Pb.Tick                                                             // {"id":153104,"open":24138,"price":23794,"seq":48345335,"time":1650874872,"tick_deep":[],"price_close":24147,"digits":3}
	klineLog.ContractCode = strings.ReplaceAll(strings.ToUpper(ContractCode), "/", "-") // 合约代码 BTC-USD
	klineLog.Ch = fmt.Sprintf("market.%v.index.1min", klineLog.ContractCode)            // kilne返回日期时间戳
	klineLog.Ts = SystemTs                                                              // 自己生成的ts
	klineLog.UpstreamTs = UpstreamTs                                                    // 上游的ts
	klineLog.ID = SystemTs                                                              // 自己生成的ts/1000
	klineLog.Mrid = 0.00                                                                // ID
	klineLog.COpen = rData.Open / math.Pow(10, rData.Digits)                            // 起始峰值
	klineLog.CClose = rData.Price / math.Pow(10, rData.Digits)                          // 结束峰值
	klineLog.High = 0                                                                   // 最高峰值
	klineLog.Low = 0                                                                    // 最低峰值
	klineLog.Amount = 0                                                                 // 成交金额 TODO
	klineLog.Vol = 0                                                                    // 成交手 TODO
	klineLog.CCount = 0                                                                 // 交易笔数 TODO
	klineLog.TsDate = time.Unix(0, SystemTs*int64(time.Millisecond))                    // Kline返回日期
	klineLog.CreateDate = time.Now()                                                    // 创建日期
	klineLog.KlineSource = kconfig.KlineSource                                          // K线来源hsb
	return klineLog
}

// RedisData 缓存进redis的数据
func RedisData(SystemTs int64, resp *wsdata.SubPriceKLineResponse) response.SubIndexKLineResponse {
	var data response.SubIndexKLineResponse
	wsData := *resp
	rData := wsData.Pb.Tick // {"p":"USDJPY","t":1650526129,"q":0,"open":127.944,"low":127.944,"high":127.944,"close":127.944}
	ContractCode := kline_common.GetSymbolById(helpers.StringToInt(fmt.Sprintf("%v", rData.Id)))
	data.Ch = fmt.Sprintf("market.%v.index.1min", ContractCode)
	data.Ts = SystemTs                                                                             // 系統生成的Ts
	data.TsDate = time.Unix(0, SystemTs*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999") // 系統生成的Ts
	data.Source = kconfig.KlineSource
	data.Tick.Id = SystemTs
	data.Tick.Vol = "0"
	data.Tick.Count = 0
	data.Tick.Open = fmt.Sprintf("%v", rData.Open/math.Pow(10, rData.Digits))
	data.Tick.Close = fmt.Sprintf("%v", rData.Price/math.Pow(10, rData.Digits))
	data.Tick.Low = ""
	data.Tick.High = ""
	data.Tick.Amount = "0"
	return data
}
