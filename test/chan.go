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
	wsmkClient.closeChan = make(chan byte, 1)
	// 子协程读数据
	go wsmkClient.DoHand()
	// 消息写入
	go func(ws *WebSocketOp) {
		for i := 1; i < 10; i++ {
			ws.readChan <- []byte(helpers.IntToString(i))
			time.Sleep(time.Second * 1)
		}
		close(ws.readChan)  // 关闭 chan
		close(ws.closeChan) // 关闭 chan
	}(wsmkClient)
	time.Sleep(time.Second * 60 * 60)
}

type WebSocketOp struct {
	readChan  chan []byte //读消息管道
	closeChan chan byte   // 关闭chan
}

func (ws *WebSocketOp) DoHand() {
	for {
		select {
		case msg, ok := <-ws.readChan:
			if !ok {
				logger.Info("不 ok 了")
				return
			}
			go ws.msgHand(msg)
			break
		case <-ws.closeChan:
			goto ERR
		}
	}
ERR:
	logger.Info("读协程已关闭")
}

func (ws *WebSocketOp) msgHand(message []byte) {
	logger.Info(fmt.Sprint("msgHand：", string(message)))
}
