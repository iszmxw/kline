package main

import (
	"encoding/json"
	"fmt"
	"goapi/pkg/logger"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	for true {
		go getCheckDataCount(0, "/web/v1/main/getCheckDataCount.json", `{
	"languageCode": "CH",
	"itemCode": "IQO",
	"checkDate": "2022-07-21"
}`)
		go getCheckDataCount(1, "/web/v1/main/getCheckDataCountByType.json", `{
    "languageCode": "CH",
    "itemCode": "IQO",
    "checkDate": "2022-07-21",
    "checkType": 1
}`)
		go getCheckDataCount(2, "/web/v1/main/getCheckDataCountByType.json", `{
    "languageCode": "CH",
    "itemCode": "IQO",
    "checkDate": "2022-07-21",
    "checkType": 2
}`)
		go getCheckDataCount(3, "/web/v1/main/getCheckDataCountByType.json", `{
    "languageCode": "CH",
    "itemCode": "IQO",
    "checkDate": "2022-07-21",
    "checkType": 3
}`)
		fmt.Println("========================================")
		time.Sleep(time.Second * 5)
	}
}

type All struct {
	Code int `json:"code"`
	Data struct {
		RegistUser []struct {
			Period      string `json:"period"`
			PeriodValue string `json:"periodValue"`
		} `json:"registUser"`
		RechargeOrder []struct {
			Period      string `json:"period"`
			PeriodValue string `json:"periodValue"`
		} `json:"rechargeOrder"`
		TradeOrder []struct {
			Period      string `json:"period"`
			PeriodValue string `json:"periodValue"`
		} `json:"tradeOrder"`
	} `json:"data"`
	Message string `json:"message"`
}

type T1 struct {
	Code int `json:"code"`
	Data struct {
		RegistUserBefore []struct {
			Period      string `json:"period"`
			PeriodValue string `json:"periodValue"`
		} `json:"registUserBefore"`
		RegistUserThat []struct {
			Period      string `json:"period"`
			PeriodValue string `json:"periodValue"`
		} `json:"registUserThat"`
		RegistUserBeforetomo []struct {
			Period      string `json:"period"`
			PeriodValue string `json:"periodValue"`
		} `json:"registUserBeforetomo"`
	} `json:"data"`
	Message string `json:"message"`
}

type T2 struct {
	Code int `json:"code"`
	Data struct {
		RechargeOrderThat []struct {
			Period      string `json:"period"`
			PeriodValue string `json:"periodValue"`
		} `json:"rechargeOrderThat"`
		RechargeOrderBefore []struct {
			Period      string `json:"period"`
			PeriodValue string `json:"periodValue"`
		} `json:"rechargeOrderBefore"`
		RechargeOrderBeforetomo []struct {
			Period      string `json:"period"`
			PeriodValue string `json:"periodValue"`
		} `json:"rechargeOrderBeforetomo"`
	} `json:"data"`
	Message string `json:"message"`
}

type T3 struct {
	Code int `json:"code"`
	Data struct {
		TradeOrderBefore []struct {
			Period      string `json:"period"`
			PeriodValue string `json:"periodValue"`
		} `json:"tradeOrderBefore"`
		TradeOrderThat []struct {
			Period      string `json:"period"`
			PeriodValue string `json:"periodValue"`
		} `json:"tradeOrderThat"`
		TradeOrderBeforetomo []struct {
			Period      string `json:"period"`
			PeriodValue string `json:"periodValue"`
		} `json:"tradeOrderBeforetomo"`
	} `json:"data"`
	Message string `json:"message"`
}

func getCheckDataCount(num int, url string, payloadBody string) {
	method := "POST"
	payload := strings.NewReader(payloadBody)
	client := &http.Client{}
	req, err := http.NewRequest(method, "http://127.0.0.1:403"+url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("Origin", "http://test.iqoptionpro.com")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Proxy-Connection", "keep-alive")
	req.Header.Add("Referer", "http://test.iqoptionpro.com/")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	req.Header.Add("token", "eae5c6d57a9649fea7a64c36e760dab4")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Error(err)
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if num == 0 {
		var data All
		err := json.Unmarshal(body, &data)
		if err != nil {
			fmt.Println("错误---", err)
			return
		}
		fmt.Println("00000-", len(data.Data.RegistUser), len(data.Data.RechargeOrder), len(data.Data.TradeOrder))
	}
	if num == 1 {
		var data T1
		err := json.Unmarshal(body, &data)
		if err != nil {
			fmt.Println("错误---", err)
			return
		}
		fmt.Println("11111-", len(data.Data.RegistUserBefore), len(data.Data.RegistUserThat), len(data.Data.RegistUserBeforetomo))

	}
	if num == 2 {
		var data T2
		err := json.Unmarshal(body, &data)
		if err != nil {
			fmt.Println("错误---", err)
			return
		}
		fmt.Println("22222-", len(data.Data.RechargeOrderThat), len(data.Data.RechargeOrderBefore), len(data.Data.RechargeOrderBeforetomo))

	}
	if num == 3 {
		var data T3
		err := json.Unmarshal(body, &data)
		if err != nil {
			fmt.Println("错误---", err)
			return
		}
		fmt.Println("33333-", len(data.Data.TradeOrderBefore), len(data.Data.TradeOrderThat), len(data.Data.TradeOrderBeforetomo))

	}
}
