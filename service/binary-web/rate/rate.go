package rate

import (
	"fmt"
	"github.com/opesun/goquery"
	"goapi/app/models"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"strings"
	"time"
)

func CollectionRateOfExchange() {
	logger.Info("采集汇率")
	for true {
		var list []models.AppCurrencyRate
		mysql.DB.Debug().Model(models.AppCurrencyRate{}).Where(map[string]interface{}{
			"is_valid": 1,
		}).Find(&list)
		// 查询系统所有币种，抓取汇率
		for index, item := range list {
			go updateRate(item.CurrencyB)
			time.Sleep(time.Duration(index) * time.Second)
		}
		time.Sleep(60 * time.Second)
	}
}

func Get(currency string) float64 {
	var rate float64
	x, err := goquery.ParseUrl(fmt.Sprintf("https://www.xe.com/zh-CN/currencyconverter/convert/?Amount=1&From=USD&To=%s", currency))
	if err != nil {
		logger.Error(err)
		return 0
	}
	result := x.Find("p.result__BigRate-sc-1bsijpp-1").Text()
	if len(result) > 0 {
		arr := strings.Split(result, " ")
		if len(arr[0]) == 0 {
			// 未抓取到
			return 0
		}
		rate = helpers.StrToFloat64(strings.ReplaceAll(arr[0], ",", ""))
	}
	return rate
}

func updateRate(currency string) {
	num := Get(currency)
	logger.Info(num)
	if num > 0 {
		mysql.DB.Model(models.AppCurrencyRate{}).Where(map[string]interface{}{
			"currency_a": "USDT",
			"currency_b": currency,
		}).Updates(map[string]interface{}{
			"b_num":       num,
			"update_date": time.Now(),
		})
	}
}
