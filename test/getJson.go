package main

import (
	"fmt"
	"goapi/config"
	"goapi/pkg/logger"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("test")
}

func main() {
	client := &http.Client{}
	request, err := http.NewRequest("POST", "https://api.adjust.com/public/v1/apps/majbv8e0m8e8/trackers", strings.NewReader(`{
	  "name": "11588@qq.com"
	}`))
	if err != nil {
		panic(err)
	}
	//request.Header.Add("Authorization", "Token token=ygMaBirkbgfB97ADsFMn")
	request.Header.Add("Content-Type", "application/json")
	response, _ := client.Do(request)
	defer response.Body.Close()
	result, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(result))
}
