package main

import (
	"encoding/json"
	"fmt"
	"goapi/pkg/logger"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	body := queryAdId("5d7318768da6704ec01fcec4bac445ce")
	fmt.Println(body.TrackerName)
}

type RespQueryAdId struct {
	Adid                        string      `json:"Adid"`
	AdvertisingId               string      `json:"AdvertisingId"`
	Tracker                     string      `json:"Tracker"`
	TrackerName                 string      `json:"TrackerName"`
	FirstTracker                string      `json:"FirstTracker"`
	FirstTrackerName            string      `json:"FirstTrackerName"`
	ClickTime                   time.Time   `json:"ClickTime"`
	InstallTime                 time.Time   `json:"InstallTime"`
	LastAppVersion              string      `json:"LastAppVersion"`
	LastAppVersionShort         string      `json:"LastAppVersionShort"`
	LastSessionTime             time.Time   `json:"LastSessionTime"`
	LastEventTimes              interface{} `json:"LastEventTimes"`
	PushToken                   string      `json:"PushToken"`
	State                       string      `json:"State"`
	InstallState                string      `json:"InstallState"`
	SignatureAcceptanceStatus   string      `json:"SignatureAcceptanceStatus"`
	SignatureVerificationResult string      `json:"SignatureVerificationResult"`
}

// 查旬 adId 信息
func queryAdId(adId string) *RespQueryAdId {
	var respQueryAdId *RespQueryAdId
	url := "https://api.adjust.com/device_service/api/v1/inspect_device?app_token=tiwyto5q38jk&advertising_id=" + adId
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		logger.Error(err)
		return nil
	}
	req.Header.Add("Authorization", "Bearer ygMaBirkbgfB97ADsFMn")
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
	err = json.Unmarshal(body, &respQueryAdId)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return respQueryAdId
}
