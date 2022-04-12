package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gostudys/huobi_contract_golang/src/ws/response"
	"goapi/app/models"
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/esdk"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
	"goapi/service/kline-sina/utils"
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
	var lineLog map[string]interface{}
	var KlineData models.KlineLog
	Symbol := queryData.Symbol
	url := utils.UrlGet(Symbol)
	str := utils.HttpGet(url, queryData.Referer)
	if len(str) <= 0 {
		return
	}
	if Symbol == "XAU" || Symbol == "XAG" {
		// 黄金白银
		lineLog = utils.Metal(str, Symbol)
	} else {
		// 外汇数据
		lineLog = utils.ForeignExchange(str, Symbol)
	}

	// 收集数据缓存到redis
	var data response.SubIndexKLineResponse
	data.Ch = lineLog["ch"].(string)
	data.Ts = lineLog["ts"].(int64)
	data.Tick.Id = lineLog["id"].(int64)
	data.Tick.Vol = "0"
	data.Tick.Count = 0
	data.Tick.Open = lineLog["c_open"].(string)
	data.Tick.Close = lineLog["c_close"].(string)
	data.Tick.Low = lineLog["low"].(string)
	data.Tick.High = lineLog["high"].(string)
	data.Tick.Amount = lineLog["amount"].(string)
	data.Tick.TradeTurnover = ""
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

	err = mysql.DB.Debug().
		Model(models.KlineLog{}).
		Select("log_id").
		Where("contract_code", lineLog["contract_code"]).
		Where("ts", lineLog["ts"]).
		Find(&KlineData).Error
	if err != nil {
		logger.Error(err)
	}
	if KlineData.LogId > 0 {
		return
	}
	cerr1 := mysql.DB.Debug().Model(models.KlineLog{}).Create(&lineLog).Error
	if cerr1 != nil {
		logger.Error(cerr1)
		return
	}
	// 历史存入 elastic
	putRes := esdk.IndexPut("kline_log", helpers.Int64ToString(KlineData.LogId), KlineData)
	if !putRes {
		logger.Error(errors.New("历史存入 elastic 失败"))
	} else {
		logger.Info("elastic-success")
	}
}
