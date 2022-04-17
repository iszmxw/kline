package main

import (
	"fmt"
	"time"
)

type WebSocketOp struct {
	closeChan chan byte // 关闭chan
}

func isChanClose(ch <-chan byte) bool {
	select {
	case _, received := <-ch:
		return !received
	default:
	}
	return false
}

func main() {
	var ws = WebSocketOp{}
	ws.open()
	go ws.close()
	go ws.close()
	go ws.close()
	go ws.close()
	go ws.close()
	for i := 0; i < 10000; i++ {
		time.Sleep(time.Second)
		if i%4 == 0 {
			fmt.Println(i)
		}
		if isChanClose(ws.closeChan) {
			fmt.Println("已关闭无需重复关闭")
		} else {
			close(ws.closeChan)
		}
	}
}

func (ws *WebSocketOp) close() {
	for {
		select {
		case res := <-ws.closeChan:
			fmt.Println("res", res)
			goto ERR
		}
	}
ERR:
	fmt.Println(123456)
}

func (ws *WebSocketOp) open() {
	ws.closeChan = make(chan byte, 1)
}
