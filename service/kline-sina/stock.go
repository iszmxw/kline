package main

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/service/kline-sina/utils"
	"strings"
	"time"
)

const (
	BaseUrl = "https://finance.sina.com.cn/money/forex/hq/asas.shtml"
)

type Info struct {
	Symbol  string `json:"symbol"`
	Desc    string `json:"desc"`
	Referer string `json:"Referer"`
}

func init() {
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
		//Info{"EUR-USD", "欧元/美元", "https://finance.sina.com.cn/money/forex/hq/EURUSD.shtml"},
		//Info{"GBP-USD", "英镑/美元", "https://finance.sina.com.cn/money/forex/hq/GBPUSD.shtml"},
		//Info{"AUD-CAD", "澳大利亚元/加拿大元", "https://finance.sina.com.cn/money/forex/hq/AUDCAD.shtml"},
		//Info{"USD-JPY", "美元/日元", "https://finance.sina.com.cn/money/forex/hq/USDJPY.shtml"},
		Info{"XAU", "伦敦金", "https://finance.sina.com.cn/futures/quotes/XAU.shtml"},
		//Info{"XAG", "伦敦银", "https://finance.sina.com.cn/futures/quotes/XAG.shtml"},
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
	Symbol := queryData.Symbol
	url := utils.UrlGet(Symbol)
	str := utils.HttpGet(url, queryData.Referer)
	if len(str) <= 0 {
		return
	}
	// 黄金白银
	if Symbol == "XAU" || Symbol == "XAG" {
		//fmt.Println(str)
		result := strings.Replace(str, fmt.Sprintf("var hq_str_hf_%v=\"", strings.ReplaceAll(Symbol, "-", "")), "", -1)
		result = strings.Replace(result, "\";", "", -1)
		//fmt.Println(result)
		result = strings.Replace(result, "\n", "", -1)
		//fmt.Println(result)
		resArr := strings.Split(result, ",")

		logger.Info(resArr)
		var _curPrice float64
		if helpers.StrToFloat64(resArr[1]) > 0 {
			_curPrice = helpers.StrToFloat64(resArr[1])
		} else {
			_curPrice = helpers.StrToFloat64(resArr[8])
		}

		Metal := cmap.New().Items()
		Metal["price"] = _curPrice
		//Metal["cls"] = config.marks[tool.get_status(resArr[100])]
		Metal["time"] = resArr[13] + " " + resArr[7]
		Metal["open-price"] = resArr[9]
		Metal["max-price"] = resArr[5]
		Metal["min-price"] = resArr[6]
		Metal["close-price"] = resArr[6]
		Metal["yestoday-price"] = resArr[6]
		Metal["amount"] = resArr[6]
		Metal["volume"] = resArr[6]
		Metal["buy"] = resArr[6]
		Metal["sell"] = resArr[6]
		Metal["buy-amount"] = resArr[6]
		Metal["sell-amount"] = resArr[6]
		Metal["amt"] = resArr[6]
		Metal["amt-value"] = resArr[6]
		logger.Info(Metal)

		// 金属

		//{
		//	'price': _curPrice,
		//	'cls': config.marks[tool.get_status(resArr[100])],
		//	'time': resArr[13] + " " + resArr[7],
		//	'open-price': resArr[9],
		//	'max-price': resArr[5],
		//	'min-price': resArr[6],
		//	'close-price': "--",
		//	'yestoday-price': resArr[8],
		//	'amount': resArr[10],
		//	'volume': "--",
		//	'buy': resArr[3],
		//	'sell': resArr[4],
		//	'buy-amount': resArr[11] === 0 ? "--" : tool.plus(tool.num_to_str(resArr[11], 2)),
		//	'sell-amount': resArr[12] === 0 ? "--" : tool.plus(tool.num_to_str(resArr[12], 2)),
		//	'amt': resArr[100] === 0 ? "--" : tool.plus(tool.num_to_str(resArr[100], 2)) + "%",
		//	'amt-value': tool.plus(tool.num_to_str(_curPrice - resArr[8], toFixedNumber))
		//}

		//// 对应数据
		//lineLog := cmap.New().Items()
		//lineLog["contract_code"] = strings.ToUpper(Symbol)                           // 合约代码 BTC-USD
		//lineLog["ch"] = fmt.Sprintf("market.%v.kline.1min", strings.ToUpper(Symbol)) // kilne返回日期时间戳
		////lineLog["ts"] = GetTs(resArr[17], resArr[0])                                 // kilne返回日期时间戳 毫秒
		////lineLog["id"] = GetTs(resArr[17], resArr[0]) / 1000                          // kilne返回日期时间戳 秒
		//lineLog["mrid"] = 0.00                                                       // ID
		//lineLog["c_open"] = resArr[5]                                                // 起始峰值
		//lineLog["c_close"] = resArr[8]                                               // 结束峰值
		//lineLog["high"] = resArr[6]                                                  // 最高峰值
		//lineLog["low"] = resArr[7]                                                   // 最低峰值
		//lineLog["amount"] = 0                                                        // 成交金额 TODO
		//lineLog["vol"] = 0                                                           // 成交手 TODO
		//lineLog["c_count"] = 0                                                       // 交易笔数 TODO
		//lineLog["ts_date"] = resArr[17] + " " + resArr[0]                            // Kline返回日期
		//lineLog["create_date"] = time.Now()                                          // 创建日期
		//lineLog["kline_source"] = "sina"                                             // K线来源新浪

	} else {
		// 外汇数据
		lineLog := utils.ForeignExchange(str, Symbol)
		var KlineData models.KlineLog
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
}
