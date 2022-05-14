package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"goapi/app/models"
	"goapi/app/response"
	"goapi/bootstrap"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	mail_notice "goapi/pkg/notice/mail"
	"goapi/pkg/redis"
	"goapi/service/kline-olymptrade/olymptrade/api"
	kconfig "goapi/service/kline-olymptrade/olymptrade/config"
	"goapi/service/kline-olymptrade/olymptrade/utils/code"
	"goapi/service/kline-olymptrade/olymptrade/wsdata"
	"net/http"
	_ "net/http/pprof"
	"runtime"
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
	if conf.GetBool("app.pprof") {
		runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪
		runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪
		go func() {
			// 启动一个 http server，注意 pprof 相关的 handler 已经自动注册过了
			if err := http.ListenAndServe(":6063", nil); err != nil {
				logger.Error(err)
			}
		}()
	}
	bootstrap.SetupRedis("15")
	defer bootstrap.RedisClose()
	var CodeData []wsdata.Info
	CodeData = code.GetData()
	start := "ok"
	for {
		if start == "ok" {
			logger.Info(CodeData)
			for _, obj := range CodeData {
				logger.Info("缓存 " + obj.Symbol)
				go SocketOlymptrade(obj)
			}
			start = "no"
		}
		time.Sleep(time.Second)
	}
}

func SocketOlymptrade(info wsdata.Info) {
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
		time.Sleep(time.Second) // 休息1秒
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
			time.Sleep(time.Second * 3) // 休息3秒
		}
		wsmkClient := new(api.WSPriceClient).Init("", session, kconfig.Proxy(), callback)
		wsmkClient.SubPriceKLine(info, func(resp *wsdata.SubPriceKLineResponse) {
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
			dType := wsData[0].E    // 读取E的类型
			rData := wsData[0].D[0] // {"p":"USDJPY","t":1650526129,"q":0,"open":127.944,"low":127.944,"high":127.944,"close":127.944}
			if dType == 1 {
				return
			}
			if dType == 2 {
				logger.Info(rData)
			}
			runStr := time.Now().UnixMilli() // 开始时间
			// RequestId 定义为 Ts 时间戳
			logger.RequestId = fmt.Sprintf("%v", rData.T)
			// olymptrade网推送的回来的数据转换为字符串
			strData, err := json.Marshal(resp) //转换成JSON返回的是byte[]
			if err != nil {
				logger.Error(err)
				return
			}
			if len(strData) <= 0 {
				logger.Error(errors.New(fmt.Sprintf(
					"订阅失败 当前订阅的 topic 为：%v,返回的数据为：%v",
					info.ContractCode, resp)))
				return
			}
			logger.Info(resp)
			if rData.T <= 0 {
				// 当前数据源匹配就入库
				if response.KlineSource(info.ContractCode) == kconfig.KlineSource {
					err = mysql.DB.Debug().Model(models.TradingAssets{}).Where("assets_code", info.ContractCode).Update("assets_status", "2").Error
					if err != nil {
						logger.Error(err)
					}
					// 通知
					mail_notice.Notice("msgNoSocket", map[string]interface{}{
						"service": kconfig.KlineSource,
						"symbol":  info.ContractCode,
					})
				}
				logger.Error(errors.New(fmt.Sprintf(
					"订阅失败 当前订阅的 topic 为：%v,返回的数据时间为：%v",
					info.ContractCode, resp)))
				return
			}
			ContractCode := info.ContractCode
			SystemTs := time.Now().UnixMilli()  // 系統生成的Ts
			UpstreamTs := int64(rData.T * 1000) // 上游的ts
			var data response.SubIndexKLineResponse
			data.Ch = fmt.Sprintf("market.%v.index.1min", ContractCode)
			data.Ts = SystemTs
			data.TsDate = time.Unix(0, SystemTs*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999") // 系統生成的Ts
			data.Source = kconfig.KlineSource
			data.Tick.Id = SystemTs
			data.Tick.Vol = "0"
			data.Tick.Count = 0
			data.Tick.Open = fmt.Sprintf("%v", rData.Open)
			data.Tick.Close = fmt.Sprintf("%v", rData.Close)
			data.Tick.Low = fmt.Sprintf("%v", rData.Low)
			data.Tick.High = fmt.Sprintf("%v", rData.High)
			data.Tick.Amount = "0"

			tips := fmt.Sprintf(
				"olymptrade时差：系统时间：%v(%v)<br>olymptrade的 ts 时间：%v(%v)",
				time.Now().UnixMilli(),
				time.Now().Format("2006-01-02 15:04:05"),
				data.Ts,
				time.Unix(data.Ts/1000, 0).Format("2006-01-02 15:04:05"),
			)
			logger.Info(tips)
			shicha := time.Now().UnixMilli() - data.Ts
			logger.Info(fmt.Sprintf("shicha=%v", shicha))
			// 时间差大于 10 秒发送邮件通知
			if shicha/1000 >= 20 {
				mail_notice.Notice("msgTimeDiff", map[string]interface{}{
					"ch":      data.Ch,
					"tips":    tips,
					"service": kconfig.KlineSource,
					"shicha":  shicha,
				})
			}
			msgData, MarshalJSONErr := json.Marshal(data) //转换成JSON返回的是byte[]
			if MarshalJSONErr != nil {
				logger.Error(MarshalJSONErr)
				return
			}
			// 收集数据缓存到redis
			key := fmt.Sprintf("%v:%v", kconfig.KlineSource, data.Ch)
			_, redisErr := redis.Client.Add(key, msgData, 0)
			if redisErr != nil {
				logger.Error(redisErr)
				return
			}

			// 当前数据源匹配就入库
			if response.KlineSource(info.ContractCode) == kconfig.KlineSource {
				var KlineData models.KlineLog
				err = mysql.DB.Debug().
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
				cerr1 := mysql.DB.Debug().Model(models.KlineLog{}).Create(&lineLogRes).Error
				if cerr1 != nil {
					logger.Error(cerr1)
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
		// 记录重连次数 连接
		num := redis.Client.Incr(fmt.Sprintf("connect:%v:%v", kconfig.KlineSource, info.Symbol))
		logger.Info(fmt.Sprintf("%v:已关闭正在第【%v】次重新启动Goroutine", info.Symbol, num))
		go SocketOlymptrade(info)
	}
}

func LogData(wsData wsdata.SubPriceKLineResponse, ContractCode string, SystemTs, UpstreamTs int64) models.KlineLog {
	var klineLog models.KlineLog
	rData := wsData[0].D[0]                                                             // {"p":"USDJPY","t":1650526129,"q":0,"open":127.944,"low":127.944,"high":127.944,"close":127.944}
	klineLog.ContractCode = strings.ReplaceAll(strings.ToUpper(ContractCode), "/", "-") // 合约代码 BTC-USD
	klineLog.Ch = fmt.Sprintf("market.%v.index.1min", klineLog.ContractCode)            // kilne返回日期时间戳
	klineLog.Ts = SystemTs                                                              // 自己生成的ts
	klineLog.UpstreamTs = UpstreamTs                                                    // 上游的ts
	klineLog.ID = SystemTs                                                              // 自己生成的ts/1000
	klineLog.Mrid = 0.00                                                                // ID
	klineLog.COpen = 0                                                                  // 起始峰值
	klineLog.CClose = rData.Close                                                       // 结束峰值
	klineLog.High = rData.High                                                          // 最高峰值
	klineLog.Low = rData.Low                                                            // 最低峰值
	klineLog.Amount = 0                                                                 // 成交金额 TODO
	klineLog.Vol = 0                                                                    // 成交手 TODO
	klineLog.CCount = 0                                                                 // 交易笔数 TODO
	klineLog.TsDate = time.Unix(0, time.Now().UnixMilli()*int64(time.Millisecond))      // Kline返回日期
	klineLog.CreateDate = time.Now()                                                    // 创建日期
	klineLog.KlineSource = kconfig.KlineSource                                          // K线来源olymptrade
	return klineLog
}
