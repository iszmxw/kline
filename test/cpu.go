package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	type topics struct {
		Topic string `json:"topic"` // K线图代码
	}
	var Topics []topics
	// 收集所有要订阅处理的 topic
	Topics = append(Topics,
		// BTC-USD
		// 1分钟k线图
		topics{
			Topic: "market.BTC-USD.index.1min",
		},
		// ETH-USD
		// 1分钟k线图
		topics{
			Topic: "market.ETH-USD.index.1min",
		})
	start := "ok"
	for {
		if start == "ok" {
			for _, obj := range Topics {
				go SocketHuoBi(obj.Topic)
			}
			start = "no"
		}
		// 消费消息队列，数据存入es里面
		//go MsgQueue()
		time.Sleep(time.Second)
	}
}

func SocketHuoBi(parentTopic string) {
	callback := make(chan struct{}, 1)
	callback <- struct{}{}
	select {
	case <-callback:
		fmt.Println(parentTopic + "：已关闭")
		fmt.Println(parentTopic + "：正在重新启动 Goroutine")
		go SocketHuoBi(parentTopic)
	}
	fmt.Println(123456)
	time.Sleep(time.Second * 5)
	runtime.Goexit()
	// 退出
}
