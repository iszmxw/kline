package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gostudys/huobi_contract_golang/src/ws/response"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/bootstrap"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/email"
	"goapi/pkg/esdk"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
	"goapi/service/kline-raceoption/raceoption/api"
	"goapi/service/kline-raceoption/raceoption/wsdata"
	"net/http"
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
	logger.Init("kline-raceoption")
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
			if err := http.ListenAndServe(":6061", nil); err != nil {
				fmt.Println(err)
			}
		}()
	}
	bootstrap.SetupRedis("15")
	defer bootstrap.RedisClose()
	var CodeData []wsdata.Info
	CodeData = append(CodeData,
		wsdata.Info{Symbol: "EUR/USD", Desc: "欧元/美元", ContractCode: "EUR-USD"},
		wsdata.Info{Symbol: "GBP/USD", Desc: "英镑/美元", ContractCode: "GBP-USD"},
		wsdata.Info{Symbol: "AUD/CAD", Desc: "澳大利亚元/加拿大元", ContractCode: "AUD-CAD"},
		wsdata.Info{Symbol: "USD/JPY", Desc: "美元/日元", ContractCode: "USD-JPY"},
		wsdata.Info{Symbol: "XAU/USD", Desc: "伦敦金", ContractCode: "GOLD"},
		wsdata.Info{Symbol: "XAG/USD", Desc: "伦敦银", ContractCode: "SILVER"},
	)
	start := "ok"
	for {
		if start == "ok" {
			logger.Info(CodeData)
			for _, obj := range CodeData {
				logger.Info("缓存 " + obj.Symbol)
				go SocketRaceOption(obj)
			}
			start = "no"
		}
		time.Sleep(time.Second)
	}
}

func SocketRaceOption(info wsdata.Info) {
	callback := make(chan struct{}, 1)
	wsmkClient := new(api.WSPriceClient).Init("", callback)
	wsmkClient.SubPriceKLine(info, func(resp *wsdata.SubPriceKLineResponse) {
		runStr := time.Now().UnixMilli() // 开始时间
		// RequestId 定义为 Ts 时间戳
		logger.RequestId = fmt.Sprintf("%v", resp.TimeId)
		// raceoption网推送的回来的数据转换为字符串
		strData, err := json.Marshal(resp) //转换成JSON返回的是byte[]
		if err != nil {
			logger.Error(err)
			return
		}
		if len(strData) <= 0 {
			logger.Error(errors.New(fmt.Sprintf(
				"订阅失败 当前订阅的 topic 为：%v,返回的数据为：%v",
				resp.Symbol, resp)))
			return
		}
		logger.Info(resp)
		currentTime := time.Now()
		startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location()).Unix()
		Surplus := (time.Now().Unix() - startTime) % 8
		if resp.TimeId == "1" {
			if Surplus == 0 {
				go func() {
					Einfo := cmap.New().Items()
					Einfo["title"] = fmt.Sprintf("环境【%v】%v raceoption 服务通知：暂停socket推送", conf.GetString("app.env"), resp.Symbol)
					Einfo["content"] = "服务停止socket"
					err = email.SendEmail(Einfo, conf.GetString("notify.email"))
					if err != nil {
						logger.Error(err)
					}
				}()
			}
			logger.Error(errors.New(fmt.Sprintf(
				"订阅失败 当前订阅的 topic 为：%v,返回的数据时间为：%v",
				resp.Symbol, resp)))
			return
		} else {
			if Surplus == 0 {
				go func() {
					Einfo := cmap.New().Items()
					Einfo["title"] = fmt.Sprintf("环境【%v】%v raceoption 服务通知：当前socket数据恢复推送", conf.GetString("app.env"), resp.Symbol)
					Einfo["content"] = "服务恢复正常"
					err = email.SendEmail(Einfo, conf.GetString("notify.email"))
					if err != nil {
						logger.Error(err)
					}
				}()
			}
		}
		ContractCode := info.ContractCode
		resp.Symbol = info.ContractCode
		var data response.SubIndexKLineResponse
		data.Ch = fmt.Sprintf("market.%v.index.1min", ContractCode)
		data.Ts = helpers.StringToInt64(resp.TimeId)
		data.Tick.Id = helpers.StringToInt64(resp.TimeId)
		data.Tick.Vol = fmt.Sprintf("%v", resp.AskVolume)
		data.Tick.Count = 0
		data.Tick.Open = "0"
		data.Tick.Close = fmt.Sprintf("%v", resp.Strike)
		data.Tick.Low = fmt.Sprintf("%v", resp.Ask)
		data.Tick.High = fmt.Sprintf("%v", resp.Bid)
		data.Tick.Amount = "0"

		var KlineData models.KlineLog
		err = mysql.DB.Debug().
			Model(models.KlineLog{}).
			Select("log_id").
			Where("contract_code", ContractCode).
			Where("upstream_ts", data.Ts).
			Where("kline_source", "raceoption").
			Find(&KlineData).Error
		if err != nil {
			logger.Error(err)
		}
		if KlineData.LogId > 0 {
			return
		}
		lineLogRes := LogData(resp)
		cerr1 := mysql.DB.Debug().Model(models.KlineLog{}).Create(&lineLogRes).Error
		if cerr1 != nil {
			logger.Error(cerr1)
			return
		}

		tips := fmt.Sprintf(
			"raceoption时差：系统时间：%v(%v)<br>raceoption的 ts 时间：%v(%v)",
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
			go func() {
				Einfo := cmap.New().Items()
				Einfo["title"] = fmt.Sprintf("环境【%v】%v raceoption 服务通知：当前时间差为：%v秒大于 20秒", conf.GetString("app.env"), data.Ch, shicha/1000)
				Einfo["content"] = tips
				err = email.SendEmail(Einfo, conf.GetString("notify.email"))
				if err != nil {
					logger.Error(err)
				}
			}()
		}
		msgData, MarshalJSONErr := json.Marshal(data) //转换成JSON返回的是byte[]
		if MarshalJSONErr != nil {
			logger.Error(MarshalJSONErr)
			return
		}
		key := fmt.Sprintf("raceoption:%v", data.Ch)
		_, redisErr := redis.Add(key, msgData, 0)
		if redisErr != nil {
			logger.Error(redisErr)
			return
		}
		// 收集数据缓存到redis

		// 历史存入 elastic
		putRes := esdk.IndexPut("kline_log", helpers.Int64ToString(KlineData.LogId), KlineData)
		if !putRes {
			logger.Error(errors.New("历史存入 elastic 失败"))
		} else {
			logger.Info("elastic-success")
		}

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
	select {
	case <-callback:
		fmt.Println(info.Symbol + "：已关闭")
		fmt.Println(info.Symbol + "：正在重新启动 Goroutine")
		go SocketRaceOption(info)
	}
}

func LogData(resp *wsdata.SubPriceKLineResponse) models.KlineLog {
	var klineLog models.KlineLog
	klineLog.ContractCode = strings.ReplaceAll(strings.ToUpper(resp.Symbol), "/", "-") // 合约代码 BTC-USD
	klineLog.Ch = fmt.Sprintf("market.%v.index.1min", klineLog.ContractCode)           // kilne返回日期时间戳
	klineLog.Ts = time.Now().UnixMilli()                                               // 自己生成的ts
	klineLog.UpstreamTs = helpers.StringToInt64(resp.TimeId)                           // 上游的ts
	klineLog.ID = helpers.StringToInt64(resp.TimeId)                                   // 自己生成的ts/1000
	klineLog.Mrid = 0.00                                                               // ID
	klineLog.COpen = 0                                                                 // 起始峰值
	klineLog.CClose = resp.Strike                                                      // 结束峰值
	klineLog.High = resp.Bid                                                           // 最高峰值
	klineLog.Low = resp.Ask                                                            // 最低峰值
	klineLog.Amount = 0                                                                // 成交金额 TODO
	klineLog.Vol = helpers.StrToFloat64(resp.AskVolume)                                // 成交手 TODO
	klineLog.CCount = 0                                                                // 交易笔数 TODO
	klineLog.TsDate = time.Unix(0, time.Now().UnixMilli()*int64(time.Millisecond))     // Kline返回日期
	klineLog.CreateDate = time.Now()                                                   // 创建日期
	klineLog.KlineSource = "raceoption"                                                // K线来源raceoption
	return klineLog
}
