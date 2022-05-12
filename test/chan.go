package main

import (
	"fmt"
	"goapi/config"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"time"
)

func init() {
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("chan")
}

func main() {
	wsmkClient := new(WebSocketOp)
	wsmkClient.readChan = make(chan []byte)
	// 子协程读数据
	go wsmkClient.DoHand()
	// 消息写入
	go func(ws *WebSocketOp) {
		for i := 1; i < 100; i++ {
			logger.Info(i)
			ws.readChan <- []byte(helpers.IntToString(i))
		}
	}(wsmkClient)
	time.Sleep(time.Second * 60 * 60)
}

type WebSocketOp struct {
	readChan chan []byte //读消息管道
}

func (ws *WebSocketOp) DoHand() {
	for {
		select {
		case msg, ok := <-ws.readChan:
			if !ok {
				break
			}
			go ws.msgHand(msg)
			break
		}
	}
}

func (ws *WebSocketOp) msgHand(message []byte) {
	logger.Info(fmt.Sprint("msgHand：", string(message)))
}
