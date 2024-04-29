package v3

import (
	"encoding/json"
	"errors"
	"fmt"
	"goapi/pkg/logger"
	"goapi/pkg/redis"
	"goapi/pkg/utils"
	wss "goapi/pkg/websocket"
	"goapi/pkg/xtime"
	"strings"
	"time"
)

type BaseController struct {
}

type Group struct {
	AgentController
	BaseController
	BindController
	ConfigController
	KlineController
	OtherController
	PromoteController
	TradeController
	UserController
}

const NoData = `{"id":"10409","msg":"no data"}`
const NoSub = `{"id":"10407","msg":"no sub"}`
const Pong = `{"id":"03","msg":"pong"}`
const TokenErr = `{"id":"10409","msg":"token error"}`

// 响应消息

func SocketResp(ws *wss.WsConn, messageType int, msg string) {
	if strings.EqualFold("1", ws.Confuse) {
		tempMap := make(map[string]interface{})
		err := json.Unmarshal([]byte(msg), &tempMap)
		if err != nil {
			logger.Error(err)
			wss.WsClose(ws)
		}
		result := conversionData(tempMap)
		message, _ := json.Marshal(result)
		msg = string(message)
	}
	//logger.Info(ws.Uuid, "---", msg)
	ws.Mux.Lock()
	WriteMessageErr := ws.Conn.WriteMessage(messageType, []byte(msg))
	ws.Mux.Unlock()
	if WriteMessageErr != nil {
		logger.Error(WriteMessageErr)
		wss.WsClose(ws)
	}
}

func conversionData(data map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for key, val := range data {
		vType := fmt.Sprintf("%T", val)
		if vType == "map[string]interface {}" {
			tempRes := conversionData(val.(map[string]interface{}))
			result[utils.GetEncryptionString(key)] = tempRes
		} else {
			result[utils.GetEncryptionString(key)] = val
		}
	}
	return result
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
		// 处理心跳数据
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		_, _ = redis.Client.Add(HeartBeatingKey, currentTime, 300) // 5分钟过期
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
