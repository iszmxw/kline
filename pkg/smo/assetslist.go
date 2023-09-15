package smo

import (
	"encoding/json"
	"goapi/app/response"
	"goapi/pkg/logger"
	"io/ioutil"
	"net/http"
	"strings"
)

type QueryData struct {
	ItemCode string `json:"itemCode"`
	Nation   string `json:"nation"`
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
}

func GetPushRateData(ItemCode, Nation string) []response.PushRateData {
	url := "https://api.ssmartoption.com/app/v1/trade/getAssetsList.json"
	method := "POST"
	data := QueryData{
		ItemCode: ItemCode,
		Nation:   Nation,
		PageNum:  1,
		PageSize: 100,
	}
	bytes, _ := json.Marshal(data)
	payload := strings.NewReader(string(bytes))
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		logger.Error(err)
		return nil
	}
	req.Header.Add("Content-Type", "application/json ")
	res, err := client.Do(req)
	if err != nil {
		logger.Error(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error(err)
		return nil
	}
	var smoData SMOres
	err = json.Unmarshal(body, &smoData)
	if err != nil {
		logger.Error(err)
		return nil
	}
	var PushRateData []response.PushRateData
	if smoData.Code == 1 {
		TradingAssetsList := smoData.Data.TradingAssetsList
		for _, item := range TradingAssetsList {
			PushRateData = append(PushRateData, response.PushRateData{
				ItemCode:     item.ItemCode,
				AssetsName:   item.AssetsName,
				AssetsCode:   item.AssetsCode,
				IsValid:      item.IsValid,
				IsShow:       item.IsShow,
				ReturnRate:   item.ReturnRate,
				PointNumber:  item.PointNumber,
				AssetsStatus: item.AssetsStatus,
				SleepTime:    "",
			})
		}
	}
	return PushRateData
}

type SMOres struct {
	Code int `json:"code"`
	Data struct {
		TradingAssetsList []struct {
			Id             int     `json:"id"`
			ItemCode       string  `json:"itemCode"`
			AssetsName     string  `json:"assetsName"`
			AssetsCode     string  `json:"assetsCode"`
			AssetsLogoUrl  string  `json:"assetsLogoUrl"`
			ReturnRate     float64 `json:"returnRate"`
			AssetsStatus   int     `json:"assetsStatus"`
			IsValid        int     `json:"isValid"`
			IsShow         int     `json:"isShow"`
			PointNumber    int     `json:"pointNumber"`
			AssetsTypeCode string  `json:"assetsTypeCode"`
			AssetsTypeName string  `json:"assetsTypeName"`
		} `json:"tradingAssetsList"`
		Total int `json:"total"`
	} `json:"data"`
	Message string `json:"message"`
}
