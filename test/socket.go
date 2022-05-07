package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"goapi/config"
	"goapi/pkg/logger"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("kline-pocketoption")
}

type Url struct {
	Server string `json:"server"`
	Name   string `json:"name"`
}

type Message struct {
	ser chan []byte // 得到地址通知
}

type Message1 struct {
	ser chan []byte // 得到地址通知
}

func main() {
	message := []byte("")
	var M Message
	var M1 Message1

	M.ser = make(chan []byte, 30)
	M1.ser = make(chan []byte, 30)
	//go func() {
	//
	//}()
	for {
		select {
		case M.ser <- message:
		case <-M1.ser:
			fmt.Println(123456)
		}
	}
	//M.Init()
}

func (S *Message) Init() {
	var urls []Url
	str := `[{"server":"wss://api-asia.po.market:8030","name":"Singapore"},{"server":"wss://api-c.po.market:8030","name":"Switzerland"},{"server":"wss://api-fin.po.market:8030","name":"Finland"},{"server":"wss://api-fr.po.market:8030","name":"Germany"},{"server":"wss://api-hk.po.market:8030","name":"Hong Kong"},{"server":"wss://api-in.po.market:8030","name":"India"},{"server":"wss://api-l.po.market:8030","name":"Great Britain"},{"server":"wss://api-msk.po.market:8030","name":"Russia"},{"server":"wss://api-sc.po.market","name":"USA"},{"server":"wss://api-sg2.po.market:8030","name":"Singapore 2"},{"server":"wss://api-us2.po.market","name":"USA 2"},{"server":"wss://api-us3.po.market:8030","name":"USA 3"},{"server":"wss://api-us4.po.market:8030","name":"USA 4"},{"server":"wss://api-fr2.po.market","name":"Germany 2"}]`
	err := json.Unmarshal([]byte(str), &urls)
	if err != nil {
		return
	}
	start := true
	for {
		if start {
			for i := 0; i < len(urls); i++ {
				item := urls[i]
				go S.wsCheck(item)
			}
			start = false
		}
		select {
		case msg, ok := <-S.ser:
			if ok {
				logger.Info(msg)
				goto SUCCESS
			}
		}
	}
SUCCESS:
	fmt.Println("找到服务")
}

func (S *Message) wsCheck(url Url) {
	wssUrl := fmt.Sprintf("%s%s", url.Server, "/socket.io/?EIO=4&transport=websocket")
	dialer := websocket.Dialer{
		//Proxy: http.ProxyURL(ws.proxy),
		HandshakeTimeout: 10000 * time.Millisecond,
	}
	logger.Info(wssUrl)
	conn, _, err := dialer.Dial(wssUrl, nil)
	if err != nil {
		logger.Info(url.Name)
		return
	}
	if conn == nil {
		logger.Info("服务异常")
	}
	for {
		_, buf, err := conn.ReadMessage()
		if err != nil {
			logger.Info(err.Error())
		}
		if len(buf) >= 0 {
			logger.Info(string(buf))
		} else {
			logger.Info("45545")
			//S.ser <- []byte(`123456`)
		}
	}
}
