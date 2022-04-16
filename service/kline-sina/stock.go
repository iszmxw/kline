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
	"goapi/service/kline-sina/utils"
	"net/http"
	"runtime"
	"time"
)

type Info struct {
	Symbol  string `json:"symbol"`
	Desc    string `json:"desc"`
	Referer string `json:"Referer"`
}

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("kline-sina")
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
			if err := http.ListenAndServe(":6062", nil); err != nil {
				fmt.Println(err)
			}
		}()
	}
	bootstrap.SetupRedis("15")
	defer bootstrap.RedisClose()
	var CodeData []Info
	CodeData = append(CodeData,
		Info{"EUR-USD", "欧元/美元", "https://finance.sina.com.cn/money/forex/hq/EURUSD.shtml"},
		Info{"GBP-USD", "英镑/美元", "https://finance.sina.com.cn/money/forex/hq/GBPUSD.shtml"},
		Info{"AUD-CAD", "澳大利亚元/加拿大元", "https://finance.sina.com.cn/money/forex/hq/AUDCAD.shtml"},
		Info{"USD-JPY", "美元/日元", "https://finance.sina.com.cn/money/forex/hq/USDJPY.shtml"},
		Info{"XAU", "伦敦金", "https://finance.sina.com.cn/futures/quotes/XAU.shtml"},
		Info{"XAG", "伦敦银", "https://finance.sina.com.cn/futures/quotes/XAG.shtml"},
	)
	for {
		for _, item := range CodeData {
			go collection(item)
		}
		time.Sleep(time.Second)
		fmt.Println(time.Second)
	}
}

func collection(queryData Info) {
	var err error
	var lineLogRes models.KlineLog
	var KlineData models.KlineLog
	Symbol := queryData.Symbol
	url := utils.UrlGet(Symbol)
	str := utils.HttpGet(url, queryData.Referer)
	if len(str) <= 0 {
		return
	}
	if Symbol == "XAU" || Symbol == "XAG" {
		// 黄金白银
		lineLogRes = utils.Metal(str, Symbol)
	} else {
		// 外汇数据
		lineLogRes = utils.ForeignExchange(str, Symbol)
	}
	err = mysql.DB.Debug().
		Model(models.KlineLog{}).
		Select("log_id").
		Where("contract_code", lineLogRes.ContractCode).
		Where("upstream_ts", lineLogRes.UpstreamTs).
		Where("kline_source", "sina").
		Find(&KlineData).Error
	if err != nil {
		logger.Error(err)
	}
	if KlineData.LogId > 0 {
		return
	}
	cerr1 := mysql.DB.Debug().Model(models.KlineLog{}).Create(&lineLogRes).Error
	if cerr1 != nil {
		logger.Error(cerr1)
		return
	}
	// 收集数据缓存到redis
	var data response.SubIndexKLineResponse
	data.Ch = lineLogRes.Ch
	data.Ts = lineLogRes.Ts
	data.Tick.Id = lineLogRes.ID
	data.Tick.Vol = fmt.Sprintf("%v", lineLogRes.Vol)
	data.Tick.Count = lineLogRes.CCount
	data.Tick.Open = fmt.Sprintf("%v", lineLogRes.COpen)
	data.Tick.Close = fmt.Sprintf("%v", lineLogRes.CClose)
	data.Tick.Low = fmt.Sprintf("%v", lineLogRes.Low)
	data.Tick.High = fmt.Sprintf("%v", lineLogRes.High)
	data.Tick.Amount = fmt.Sprintf("%v", lineLogRes.Amount)
	tips := fmt.Sprintf(
		"新浪时差：系统时间：%v(%v)<br>新浪的 ts 时间：%v(%v)",
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
			Einfo["title"] = fmt.Sprintf("环境【%v】%v新浪服务通知：当前时间差为：%v秒大于 20秒", conf.GetString("app.env"), lineLogRes.Ch, shicha/1000)
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
	_, redisErr := redis.Add(data.Ch, msgData, 0)
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
}
