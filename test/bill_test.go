package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestBill(t *testing.T) {
	Bill()
	smo := 21971206.15
	uad := 3331087.5
	all := smo + uad
	fmt.Println(all)
	// 25302293.65 - 4969979 = 20332314.65
	time.Sleep(time.Minute * 2)
}

func Bill() {
	url := "https://merchant.idcashypay.com/prod-api/mchManage/api/mchBillList?pageNum=1&pageSize=14000&createStartTime=2014-05-01%252000%253A00%253A00&createEndTime=2024-05-31%252023%253A59%253A59"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "zh-CN")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzUxMiJ9.eyJsb2dpbl91c2VyX2tleSI6Im1jaGZRMzZOeHZ3SE06N2U2ZDY1YmQtNmMxMy00YWJlLTgzNzYtMGFhNDc3ZWZkNDg4In0.pLn3x1MfLPS_01-NDnjmIICvWQnHJIRGlfnRGGsKabfRkKfwhi-nA_2vO4vc0VsgK-OowDEgheOKnB8o1KaTDQ")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Referer", "https://merchant.idcashypay.com/mchManage/financeManage/mchBillList")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"125\", \"Chromium\";v=\"125\", \"Not.A/Brand\";v=\"24\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("Cookie", "username=mchfQ36NxvwHM; rememberMe=true; password=ZtJ2jWQKdlNiy4ZFle783TQyPP5zODAaHKAOkf/+3pctCcLzNKT8/Xk4uXXj7zU9AIAtSr2zB8MB/ko40URoag; acw_tc=0bc1a05317171452168357709e4e8d27ffee3e061d81ac8c7bc11195335b3c")
	req.Header.Add("Host", "merchant.idcashypay.com")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var Resp TResp
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	Fee := 0.0
	OrderAmount := 0.0
	ChangeAmount := 0.0
	for _, item := range Resp.Rows {
		Fee += item.Fee
		OrderAmount += item.OrderAmount
		if strings.EqualFold("payIn", item.BillType) || strings.EqualFold("payOut_rev", item.BillType) {
			// 加法
			ChangeAmount = ChangeAmount + item.ChangeAmount
		}
		if strings.EqualFold("withdrawal", item.BillType) || strings.EqualFold("payOut", item.BillType) {
			// 减法
			ChangeAmount = ChangeAmount - item.ChangeAmount
		}
	}
	fmt.Println(fmt.Sprintf("Fee:%v", Fee))
	fmt.Println(fmt.Sprintf("OrderAmount:%v", OrderAmount))
	fmt.Println(fmt.Sprintf("ChangeAmount:%v", ChangeAmount))
}

type TResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Rows []struct {
		Balance      float64 `json:"balance"`      // 账户总金额
		BillType     string  `json:"billType"`     // 业务类型
		ChangeAmount float64 `json:"changeAmount"` // 变更金额
		CreateTime   string  `json:"createTime"`   // 创建时间
		Fee          float64 `json:"fee"`          // 手续费
		Id           int     `json:"id"`           // ID
		MchId        int     `json:"mchId"`        // 商户ID
		MchOrderNo   string  `json:"mchOrderNo"`   // 商户订单号
		OrderAmount  float64 `json:"orderAmount"`  // 订单金额
		OrderNo      string  `json:"orderNo"`      // 平台订单号
	} `json:"rows"`
	Total int `json:"total"`
}
