package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"goapi/config"
	"goapi/pkg/logger"
	"net/http"
	"net/url"
)

func init() {
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("test")
}

func main() {
	wssUrl := "wss://stream.binance.com:9443/ws/ethusd@kline_1m"
	dialer := websocket.Dialer{
		Proxy: http.ProxyURL(&url.URL{
			Scheme: "http",
			//Host:   "47.104.184.7:888",
			//Host:   "47.242.148.57:18888",
			Host: "127.0.0.1:10809",
			//Host: "42.101.43.106:16790",
			//User: url.UserPassword("8319", "9742"),
		}),
	}
	conn, response, err := dialer.Dial(wssUrl, nil)
	fmt.Println(conn, response, err)
}
