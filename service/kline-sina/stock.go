package main

import (
	"fmt"
	"goapi/app/models"
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
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
	}
}
