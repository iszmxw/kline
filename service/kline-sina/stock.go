package main

import (
	"errors"
	"fmt"
	"github.com/axgle/mahonia"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/service/kline-sina/utils"
	"io"
	"io/ioutil"
	"net/http"
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
	logger.Init("sina")
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
		//Info{"XAU", "伦敦金", "https://finance.sina.com.cn/money/forex/hq/XAU.shtml"},
		//Info{"XAG", "伦敦银", "https://finance.sina.com.cn/money/forex/hq/XAG.shtml"},
	)
	for {
		for _, item := range CodeData {
			go collection(item.Symbol)
		}
		time.Sleep(time.Second)
		fmt.Println(time.Second)
	}
}

func collection(Symbol string) {
	Symbol = strings.ToLower(Symbol)
	contractCode := strings.ReplaceAll(Symbol, "-", "")
	url := fmt.Sprintf(
		"https://hq.sinajs.cn/rn=%vlist=fx_s%v",
		time.Now().Unix(),
		contractCode,
	)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Error(err)
		return
	}
	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36",
	)
	req.Header.Set("Referer", BaseUrl)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		logger.Error(err)
		return
	}
	if response.StatusCode != http.StatusOK {
		logger.Error(errors.New(fmt.Sprintf("Response failed with status code: %d and\n", response.StatusCode)))
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Error(err)
		}
	}(response.Body)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Error(err)
		return
	}
	dec := mahonia.NewDecoder("gbk")
	str := dec.ConvertString(string(body))
	//fmt.Println(str)
	result := strings.Replace(str, fmt.Sprintf("var hq_str_fx_s%v=\"", contractCode), "", -1)
	//fmt.Println(result)
	result = strings.Replace(result, "\";", "", -1)
	//fmt.Println(result)
	result = strings.Replace(result, "\n", "", -1)
	//fmt.Println(result)
	resArr := strings.Split(result, ",")
	//logger.Info(resArr)
	//fmt.Println(resArr)
	sinaTest(resArr, Symbol)
	// 对应数据
	lineLog := cmap.New().Items()
	lineLog["contract_code"] = strings.ToUpper(Symbol)                           // 合约代码 BTC-USD
	lineLog["ch"] = fmt.Sprintf("market.%v.kline.1min", strings.ToUpper(Symbol)) // kilne返回日期时间戳
	lineLog["ts"] = utils.GetTs(resArr[17], resArr[0])                           // kilne返回日期时间戳 毫秒
	lineLog["id"] = utils.GetTs(resArr[17], resArr[0]) / 1000                    // kilne返回日期时间戳 秒
	lineLog["mrid"] = 0.00                                                       // ID
	lineLog["c_open"] = resArr[5]                                                // 起始峰值
	lineLog["c_close"] = resArr[8]                                               // 结束峰值
	lineLog["high"] = resArr[6]                                                  // 最高峰值
	lineLog["low"] = resArr[7]                                                   // 最低峰值
	lineLog["amount"] = 0                                                        // 成交金额 TODO
	lineLog["vol"] = 0                                                           // 成交手 TODO
	lineLog["c_count"] = 0                                                       // 交易笔数 TODO
	lineLog["ts_date"] = resArr[17] + " " + resArr[0]                            // Kline返回日期
	lineLog["create_date"] = time.Now()                                          // 创建日期
	lineLog["kline_source"] = "sina"                                             // K线来源新浪
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

// 测试新浪数据
func sinaTest(resArr []string, Symbol string) {
	_data := cmap.New().Items()
	_data["name"] = strings.Replace(resArr[9], "\\u5373\\u671F\\u6C47\\u7387", "", -1)
	_data["price"] = resArr[8]
	_data["spj"] = resArr[3]
	_data["open"] = resArr[5]
	_data["higner"] = resArr[6]
	_data["lower"] = resArr[7]
	_data["updownCount"] = utils.Decimal(helpers.StrToFloat64(resArr[8]) - helpers.StrToFloat64(resArr[3]))
	_data["fullcompany"] = resArr[13]
	_data["company"] = utils.CutStr(resArr[13], 0, 40)
	_data["percent"] = utils.GetPercent(helpers.StrToFloat64(resArr[8]), helpers.StrToFloat64(resArr[3]), "percent")
	_data["updown"] = utils.GetPercent(helpers.StrToFloat64(resArr[8]), helpers.StrToFloat64(resArr[3]), "updown")
	_data["zf"] = utils.GetPercent2(helpers.StrToFloat64(resArr[6])-helpers.StrToFloat64(resArr[7]), helpers.StrToFloat64(resArr[3]))
	_data["bofu"] = helpers.StrToFloat64(resArr[6]) - helpers.StrToFloat64(resArr[7])
	_data["date"] = resArr[17] + " " + resArr[0]
	_data["created_at"] = time.Now()
	_data["updated_at"] = time.Now()
	_data["symbol"] = strings.ToUpper(Symbol)
	var Sina models.Sina
	err := mysql.DB.Debug().
		Model(models.Sina{}).
		Select("id").
		Where("symbol", _data["symbol"]).
		Where("date", _data["date"]).
		Find(&Sina).Error
	if err != nil {
		logger.Error(err)
	}
	if Sina.Id > 0 {
		return
	}
	cerr := mysql.DB.Debug().Model(models.Sina{}).Create(&_data).Error
	if cerr != nil {
		logger.Error(cerr)
	}
}
