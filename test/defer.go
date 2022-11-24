package main

import (
	"fmt"
	"goapi/config"
	"goapi/pkg/logger"
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
	for i := 1; i <= 5; i++ {
		ws := new(WebSocketOp)
		ws.closeChan = make(chan byte, 1)
		ws.readChan = make(chan time.Time)
		go ws.readLoop(i)
		go ws.Close(i)
		time.Sleep(time.Second * time.Duration(i))
	}
	time.Sleep(time.Minute * 10)
}

type WebSocketOp struct {
	closeChan chan byte      // 关闭chan
	readChan  chan time.Time //读消息管道
}

func (ws *WebSocketOp) readLoop(num int) {
	defer func() {
		fmt.Println(fmt.Sprintf("已经退出协诚：%v", num))
	}()
	defer func() {
		fmt.Println(123456)
	}()
	for {
		select {
		case ws.readChan <- time.Now():
			logger.Info(ws.readChan)
			break
		case <-ws.closeChan:
			logger.Info("读协程关闭chan")
			logger.Info("退出循环")
			return
		}
	}
}

func (ws *WebSocketOp) Close(num int) {
	var ticker = time.NewTicker(2 * time.Second)
	i := 0
	for t := range ticker.C {
		i++
		logger.Info(fmt.Sprint(fmt.Sprintf("当前协诚：%v", num), t.Format("2006-01-02 15:04")))
		if i >= num*3 {
			close(ws.closeChan)
			return
		}
	}
}
