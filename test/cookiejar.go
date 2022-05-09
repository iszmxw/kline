package main

import (
	"fmt"
	"goapi/config"
	"goapi/pkg/logger"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("session")
}

func Get() string {
	var token string
	jar, err := cookiejar.New(nil)
	if err != nil {
		logger.Error(err)
		return token
	}
	req, err := http.NewRequest("GET", "https://qxbroker.com/zh/sign-up/fast/", nil)
	if err != nil {
		logger.Error(err)
		return token
	}
	c := http.Client{
		Transport: &http.Transport{
			// 设置代理
			Proxy: http.ProxyURL(&url.URL{
				Scheme: "http",
				Host:   "127.0.0.1:10809",
				//Host: "42.101.43.106:16790",
				//User: url.UserPassword("8319", "9742"),
			}),
		},
		Jar: jar,
	}
	response, err := c.Do(req)
	if err != nil {
		logger.Error(err)
		return token
	}
	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err == nil {
			html := string(body)
			if len(html) > 0 {
				js := html[strings.Index(html, `"token"`):]
				js = js[:strings.Index(js, `",`)]
				token = strings.ReplaceAll(js, `"token":"`, "")
			}
		}
	}
	return token
}

func main() {
	token := Get()
	fmt.Println(token)
}
