package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/logger"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("kline-raceoption")
	// 初始化 SQL
	bootstrap.SetupDB()
}

var addr = flag.String("addr", "trading.raceoption.com", "服务器地址")
var subMsg = flag.String("subMsg", `{"operation":"SUBSCRIBE.TICK","id":"EUR/USD","param":""}`, "订阅主题")
//var subMsg1 = flag.String("subMsg1", `{"sub":"market.ETH-USD.index.1min","id":"EUR/USD","param":""}`, "订阅主题")

func main() {
	bootstrap.SetupRedis("15")
	defer bootstrap.RedisClose()
	flag.Parse()
	//log.SetFlags(0)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: *addr, Path: "/api/ws/price"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial("wss://trading.raceoption.com/api/ws/price", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer func(c *websocket.Conn) {
		err := c.Close()
		if err != nil {
			logger.Info(err)
		}
	}(c)

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	err = c.WriteMessage(websocket.TextMessage, []byte(*subMsg))
	if err != nil {
		log.Println("write:", err)
		return
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		//case t := <-ticker.C:
			// 订阅消息
			//logger.Info(t.String())
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
