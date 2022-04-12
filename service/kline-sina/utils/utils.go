package utils

import (
	"errors"
	"fmt"
	"github.com/axgle/mahonia"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func UrlGet(Symbol string) string {
	contractCode := strings.ReplaceAll(strings.ToLower(Symbol), "-", "")
	var url string
	// 黃金、白银
	if strings.ToUpper(Symbol) == "XAU" || strings.ToUpper(Symbol) == "XAG" {
		url = fmt.Sprintf(
			"https://hq.sinajs.cn/?_=%v/&list=hf_%v",
			time.Now().Unix(),             // 时间戳
			strings.ToUpper(contractCode), // 大写
		)
	} else {
		url = fmt.Sprintf(
			"https://hq.sinajs.cn/rn=%vlist=fx_s%v",
			time.Now().Unix(),             // 时间戳
			strings.ToLower(contractCode), // 小写
		)
	}
	return url
}

func HttpGet(url, Referer string) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Error(err)
		return ""
	}
	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36",
	)
	req.Header.Set("Referer", Referer)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		logger.Error(err)
		return ""
	}
	if response.StatusCode != http.StatusOK {
		logger.Error(errors.New(fmt.Sprintf("Response failed with status code: %d and\n", response.StatusCode)))
		return ""
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
		return ""
	}
	return mahonia.NewDecoder("gbk").ConvertString(string(body))
}

// GetTs 格式化时间戳
func GetTs(a, b string) int64 {
	if len(a) <= 0 || len(b) <= 0 {
		return 0
	}
	// 时间 to 时间戳
	// loc, _ := time.LoadLocation("Asia/Shanghai")                     // 设置时区
	tt, err := time.ParseInLocation("2006-01-02 15:04:05", a+" "+b, new(time.Location)) // 2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
	if err != nil {
		return 0
	}
	//logger.Info(a + b)
	//logger.Info(tt.UnixMilli())
	return tt.UnixMilli()
}

func CutStr(str string, a, b int) string {
	s := regexp.MustCompile(`/([^\x00-\xff])/g`).ReplaceAllString(str, "\\x00$1")
	if len(s) < b {
		return str
	} else {
		return regexp.MustCompile(`/\x00/g`).ReplaceAllString(s[a:b-3], "") + "..."
	}
}

func GetPercent(a, b float64, t string) string {
	if t == "percent" {
		if a > 0.00 || a < 0.00 {
			return fmt.Sprintf("%v", Decimal(Decimal(10000*(a-b)/b)/100)) + "%"
		} else {
			return "0.00%"
		}
	}
	if t == "updown" {
		if a > b {
			return "up"
		}
		if a < b {
			return "down"
		}
	}
	return ""
}

func GetPercent2(a, b float64) string {
	if a > 0.00 || a < 0.00 {
		return fmt.Sprintf("%v", Decimal(Decimal(10000*a/b)/100)) + "%"
	} else {
		return "0.00%"
	}
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", value), 64)
	return value
}

// 外汇数据组装

func ForeignExchange(str, Symbol string) map[string]interface{} {
	//fmt.Println(str)
	result := strings.Replace(str, fmt.Sprintf("var hq_str_fx_s%v=\"", strings.ReplaceAll(strings.ToLower(Symbol), "-", "")), "", -1)
	//fmt.Println(result)
	result = strings.Replace(result, "\";", "", -1)
	//fmt.Println(result)
	result = strings.Replace(result, "\n", "", -1)
	//fmt.Println(result)
	resArr := strings.Split(result, ",")
	//logger.Info(resArr)
	//fmt.Println(resArr)
	//sinaTest(resArr, Symbol)
	// 对应数据
	lineLog := cmap.New().Items()
	lineLog["contract_code"] = strings.ToUpper(Symbol)                           // 合约代码 BTC-USD
	lineLog["ch"] = fmt.Sprintf("market.%v.index.1min", strings.ToUpper(Symbol)) // kilne返回日期时间戳
	lineLog["ts"] = GetTs(resArr[17], resArr[0])                                 // kilne返回日期时间戳 毫秒
	lineLog["id"] = GetTs(resArr[17], resArr[0]) / 1000                          // kilne返回日期时间戳 秒
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
	return lineLog
}

// 金属数据组装

func Metal(str, Symbol string) map[string]interface{} {
	var statusVal string
	//fmt.Println(str)
	result := strings.Replace(str, fmt.Sprintf("var hq_str_hf_%v=\"", strings.ReplaceAll(Symbol, "-", "")), "", -1)
	result = strings.Replace(result, "\";", "", -1)
	//fmt.Println(result)
	result = strings.Replace(result, "\n", "", -1)
	//fmt.Println(result)
	resArr := strings.Split(result, ",")
	resArr = append([]string{"hf_" + Symbol}, resArr...)
	logger.Info(resArr)
	var _curPrice float64
	if helpers.StrToFloat64(resArr[1]) > 0 {
		_curPrice = helpers.StrToFloat64(resArr[1])
	} else {
		_curPrice = helpers.StrToFloat64(resArr[8])
	}

	MetalData := cmap.New().Items()
	MetalData["price"] = _curPrice
	if helpers.StrToFloat64(resArr[8]) > 0 {
		statusVal = fmt.Sprintf("%v", (_curPrice-helpers.StrToFloat64(resArr[8]))/helpers.StrToFloat64(resArr[8])*100)
	} else {
		statusVal = "0"
	}
	MetalData["cls"] = GetStatus(helpers.StrToFloat64(statusVal), 0)
	MetalData["time"] = resArr[13] + " " + resArr[7]
	MetalData["open-price"] = resArr[9]
	MetalData["max-price"] = resArr[5]
	MetalData["min-price"] = resArr[6]
	MetalData["close-price"] = resArr[6]
	MetalData["yestoday-price"] = resArr[8]
	MetalData["amount"] = resArr[10]
	MetalData["volume"] = "--"
	MetalData["buy"] = resArr[3]
	MetalData["sell"] = resArr[4]
	MetalData["buy-amount"] = helpers.StrToFloat64(resArr[11])
	MetalData["sell-amount"] = helpers.StrToFloat64(resArr[12])
	MetalData["amt"] = statusVal + "%"
	MetalData["amt-value"] = _curPrice - helpers.StrToFloat64(resArr[8])
	logger.Info(MetalData)

	// 金属
	// 对应数据
	lineLog := cmap.New().Items()
	lineLog["contract_code"] = getCode(Symbol)                           // 合约代码 BTC-USD
	lineLog["ch"] = fmt.Sprintf("market.%v.index.1min", getCode(Symbol)) // kilne返回日期时间戳
	lineLog["ts"] = GetTs(resArr[13], resArr[7])                         // kilne返回日期时间戳 毫秒
	lineLog["id"] = GetTs(resArr[13], resArr[7]) / 1000                  // kilne返回日期时间戳 秒
	lineLog["mrid"] = 0.00                                               // ID
	lineLog["c_open"] = resArr[9]                                        // 起始峰值
	lineLog["c_close"] = _curPrice                                       // 结束峰值
	lineLog["high"] = resArr[5]                                          // 最高峰值
	lineLog["low"] = resArr[6]                                           // 最低峰值
	lineLog["amount"] = 0                                                // 成交金额 TODO
	lineLog["vol"] = 0                                                   // 成交手 TODO
	lineLog["c_count"] = 0                                               // 交易笔数 TODO
	lineLog["ts_date"] = resArr[13] + " " + resArr[7]                    // Kline返回日期
	lineLog["create_date"] = time.Now()                                  // 创建日期
	lineLog["kline_source"] = "sina"                                     // K线来源新浪
	return lineLog
}

func getCode(Symbol string) string {
	switch Symbol {
	case "XAU":
		Symbol = "GOLD"
		break
	case "XAG":
		Symbol = "SILVER"
		break
	default:
		break
	}
	return Symbol
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
	_data["updownCount"] = Decimal(helpers.StrToFloat64(resArr[8]) - helpers.StrToFloat64(resArr[3]))
	_data["fullcompany"] = resArr[13]
	_data["company"] = CutStr(resArr[13], 0, 40)
	_data["percent"] = GetPercent(helpers.StrToFloat64(resArr[8]), helpers.StrToFloat64(resArr[3]), "percent")
	_data["updown"] = GetPercent(helpers.StrToFloat64(resArr[8]), helpers.StrToFloat64(resArr[3]), "updown")
	_data["zf"] = GetPercent2(helpers.StrToFloat64(resArr[6])-helpers.StrToFloat64(resArr[7]), helpers.StrToFloat64(resArr[3]))
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

// GetStatus 比较数值
// ['gray', 'red', 'green']
func GetStatus(amt, base float64) string {
	if amt == base {
		return "gray" // 0
	} else if amt > base {
		return "red" // 1
	} else {
		return "green" // 2
	}
}

func GetBuyAmount() {
	//arr[11] === 0 ? "--" : tool.plus(tool.num_to_str(arr[11], 2)),
}
