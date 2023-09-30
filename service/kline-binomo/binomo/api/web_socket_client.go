package api

import (
	"container/list"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"goapi/pkg/logger"
	"goapi/service/kline-binomo/binomo/wsdata"
	KlineCommon "goapi/service/kline-common"
	"net/http"
	"net/url"
	"reflect"
)

func WsInit() *WsClient {
	wc := new(WsClient)
	wc.DeviceId = "ba96afbe95990cefec98a94200bb3545"
	wc.AuthToken = "d98d908d-744a-4b4c-b3a0-1fbe4f2ed993"
	wc.SubTopicMap = make(map[string]*CallbackFunc)
	wc.List = list.New()
	return wc
}

func (wsIx *WSPriceClient) SubPriceKLine(info KlineCommon.Info, callbackFun OnSubPriceKLineResponse) {
	wsIx.sub(info, callbackFun, reflect.TypeOf(wsdata.SubPriceKLineResponse{}))
}

func (wc *WsClient) WsConnect(proxyURL *url.URL) {
	// WebSocket服务器的URL
	serverURL := "wss://as.binomo.com"
	// 创建WebSocket连接
	u, _ := url.Parse(serverURL)
	headers := http.Header{}
	// 您的cookies信息，将其替换为实际的cookies
	cookies := []*http.Cookie{
		{Name: "device_type", Value: "web"},
		{Name: "device_id", Value: wc.DeviceId},
		{Name: "authtoken", Value: wc.AuthToken},
	}
	for _, cookie := range cookies {
		headers.Add("Cookie", cookie.String())
	}
	// 创建一个HTTP Transport，并设置代理
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	// 使用自定义Transport创建WebSocket连接
	dialer := websocket.DefaultDialer
	dialer.NetDial = transport.Dial
	conn, _, err := dialer.Dial(u.String(), headers)
	if err != nil {
		logger.Error(fmt.Errorf("WebSocket连接失败:%v", err))
		wc.Close()
		return
	}
	wc.Conn = conn
	go wc.msgHand()
}

// 发送数据

func (wc *WsClient) Send(Symbol string) {
	message := fmt.Sprintf(`{"action":"subscribe","rics":["%s"]}`, Symbol)
	fmt.Println("开始发送")
	if wc.Conn != nil {
		defer wc.Mux.Unlock()
		wc.Mux.Lock()
		err := wc.Conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			err = fmt.Errorf("发送消息失败:%v", err)
			logger.Error(err)
			wc.Close()
			return
		}
	}
}

// 读取消息
func (wc *WsClient) msgHand() {
	for wc.Conn != nil {
		// 接收服务器的响应
		_, response, err := wc.Conn.ReadMessage()
		if err != nil {
			logger.Error(fmt.Errorf("读取响应失败:%v", err))
			wc.Close()
			return
		}
		var Res wsdata.Res
		err = json.Unmarshal(response, &Res)
		if err != nil {
			logger.Error(fmt.Errorf("解析数据失败:%v", err))
			wc.Close()
			return
		}
		if !Res.Success {
			logger.Error(fmt.Errorf("返回错误信息:%v", Res.Errors))
			wc.Close()
			return
		}
		if len(Res.Errors) > 0 {
			logger.Error(fmt.Errorf("返回错误信息:%v", Res.Errors))
			wc.Close()
			return
		}
		if len(Res.Data) > 1 {
			fmt.Println("超过一个元素")
		}
		if Res.Data[0].Action != "assets" {
			fmt.Println(Res.Data[0].Action)
		}
		Assets := Res.Data[0].Assets
		for _, item := range Assets {
			byteItem, err := json.Marshal(item)
			if err != nil {
				logger.Error(fmt.Errorf("返回错误信息:%v", Res.Errors))
				wc.Close()
				return
			}
			wc.handleSubCallbackFun(item.Ric, byteItem, item)
		}
	}
}

// 关闭

func (wc *WsClient) Close() {
	if wc.Conn != nil {
		err := wc.Conn.Close()
		if err != nil {
			logger.Error(fmt.Errorf("关闭长连接失败:%v", err))
			return
		}
	} else {
		fmt.Println("已关闭")
	}
	wc.Callback <- 1
	fmt.Println("正在关闭")
}
