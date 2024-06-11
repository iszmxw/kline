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
	url := "https://merchant.idcashypay.com/prod-api/mchManage/api/mchBillList?pageNum=1&pageSize=10000&billType=payOut"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "zh-CN")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzUxMiJ9.eyJsb2dpbl91c2VyX2tleSI6Im1jaGZRMzZOeHZ3SE06ODcxZTg0Y2YtZjkyMi00NmM0LWJmZDgtMzEwZGMzMjZjM2Q1In0.YEXsmln6aJC4md-B5r3OIah3I3tmReN6MA4GpTLTrTfxo-LZcfxI3Quo8BgKe6kfNKDNdzEdRCfLtTQBmNmcoA")
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
	req.Header.Add("Cookie", "username=mchfQ36NxvwHM; rememberMe=true; acw_tc=0bc159c417177487851791296e9bd40b15e4ee9e31212a3f906eb3416bf04e; password=Ke2zl51o5Jv9Cdqg0vMUzdyL/O7m92UQ+5A59AgnS7WGGPjJona8zUgbkSW3G95WnAeEvHNFPVtMSNreK5IDTQ")

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
