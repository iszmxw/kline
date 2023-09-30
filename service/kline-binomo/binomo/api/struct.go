package api

import (
	"container/list"
	"github.com/gorilla/websocket"
	"goapi/service/kline-binomo/binomo/wsdata"
	KlineCommon "goapi/service/kline-common"
	"reflect"
	"sync"
)

var App *WSPriceClient

var Once sync.Once

type CallbackFunc struct {
	fun   interface{}
	param reflect.Type
}

type closeChanCallback chan int

type WsClient struct {
	DeviceId     string       // 设备 ID
	AuthToken    string       // 授权 token
	Mux          sync.RWMutex // 锁
	Conn         *websocket.Conn
	Info         KlineCommon.Info
	SubTopicMap  map[string]*CallbackFunc // 订阅历史
	ContractCode string                   // 订阅的币种
	Symbol       string                   // 订阅的币种
	List         *list.List
	Callback     closeChanCallback
	Once         sync.Once
}

type WSPriceClient struct {
	*WsClient
}

type OnSubPriceKLineResponse func(*wsdata.SubPriceKLineResponse)
