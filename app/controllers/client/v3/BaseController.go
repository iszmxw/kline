package v3

import (
	"errors"
	"fmt"
	"goapi/pkg/logger"
	"goapi/pkg/redis"
	wss "goapi/pkg/websocket"
	"goapi/pkg/xtime"
	"time"
)

type BaseController struct {
}

type Group struct {
	BaseController
	KlineController
	NoticeController
	AgentController
	ConfigController
	TradeController
}

// 响应消息
func socketResp(ws *wss.WsConn, messageType int, msg string) {
	ws.Mux.Lock()
	_ = ws.Conn.WriteMessage(messageType, []byte(msg))
	ws.Mux.Unlock()
}

func IsDie(HeartBeatingKey string) bool { // 检测是否死掉
	var isDie bool
	// 检测心跳活跃
	HeartBeatingTime, _ := redis.Client.Get(HeartBeatingKey)
	if len(HeartBeatingTime) > 0 {
		nowTime := time.Now().Unix()
		HeartBeatTime := xtime.Parse(HeartBeatingTime).Unix()
		interval := nowTime - HeartBeatTime // 心跳间隔
		//logger.Info(fmt.Sprintf("nowTime(%v) - HeartBeatTime(%v)=%v", nowTime, HeartBeatTime, interval))
		if interval > 30 {
			logger.Error(errors.New("30秒无心跳，主动断开"))
			isDie = true
		}
	} else {
		logger.Error(errors.New("心跳初始化失效，或者心跳已停止"))
		isDie = true
	}
	return isDie
}

// 心跳观测

func HeartBeatingWatch(ws *wss.WsConn, socketId string) {
	for !ws.IsClose {
		HeartBeatingKey := fmt.Sprintf("ws_index:%v:heart_beating_key", socketId)
		if IsDie(HeartBeatingKey) {
			goto STOP
		}
		time.Sleep(time.Millisecond * 1000) // 每隔1000
	}
STOP:
	logger.Error(errors.New("暂无心跳关闭链接"))
	wss.WsClose(ws)
}
