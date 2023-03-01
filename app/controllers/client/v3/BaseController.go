package v3

import wss "goapi/pkg/websocket"

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
