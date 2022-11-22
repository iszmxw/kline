package rate

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"goapi/app/models"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func CollectionRateOfExchange() {
	logger.Info("采集汇率")
	for true {
		Get()
		time.Sleep(60 * time.Second)
	}
}

func Get() {
	req, _ := http.NewRequest("GET", "https://www.xe.com/api/protected/midmarket-converter/", nil)
	// 权限
	sEnc := base64.StdEncoding.EncodeToString([]byte("lodestar:5RTkubKoMT34MOytElLnZ90srV0VmcFM"))
	fmt.Println(sEnc)
	logger.Info(sEnc)
	req.Header.Set("authorization", fmt.Sprintf("Basic %s", sEnc))
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		logger.Error(err)
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logger.Error(err)
			return
		}
	}(resp.Body)
	respByte, _ := ioutil.ReadAll(resp.Body)
	var data ReteData
	err = json.Unmarshal(respByte, &data)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(data)
	if data.Rates.IDR > 0 {
		mysql.DB.Model(models.AppCurrencyRate{}).Where(map[string]interface{}{
			"currency_a": "USDT",
			"currency_b": "IDR",
		}).Updates(map[string]interface{}{
			"b_num":       data.Rates.IDR,
			"update_date": time.Now(),
		})
	}
	if data.Rates.INR > 0 {
		mysql.DB.Model(models.AppCurrencyRate{}).Where(map[string]interface{}{
			"currency_a": "USDT",
			"currency_b": "INR",
		}).Updates(map[string]interface{}{
			"b_num":       data.Rates.INR,
			"update_date": time.Now(),
		})
	}

}

type ReteData struct {
	Timestamp int64 `json:"timestamp"`
	Rates     struct {
		IDR float64 `json:"IDR"`
		INR float64 `json:"INR"`
	} `json:"rates"`
}
