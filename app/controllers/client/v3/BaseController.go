package v3

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/pkg/logger"
	"goapi/pkg/utils"
	wss "goapi/pkg/websocket"
	"strings"
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

type Header struct {
	// 应用信息
	DeviceIdentity string `json:"deviceidentity"`
	IsSecret       string `json:"issecert"`
	Sign           string `json:"sign"`
	Token          string `json:"token"`
	Version        string `json:"version"`
	NationLanguage string `json:"nationlanguage"`
	UserAgent      string `json:"User-Agent"`
	ContentType    string `json:"Content-Type"`
}

func getHeaderParams(c *gin.Context) (Header Header) {
	Header.DeviceIdentity = c.GetHeader("deviceidentity")
	Header.IsSecret = c.GetHeader("issecert")
	Header.Sign = c.GetHeader("sign")
	Header.Token = c.GetHeader("token")
	Header.Version = c.GetHeader("version")
	Header.NationLanguage = c.GetHeader("nationlanguage")
	Header.UserAgent = c.GetHeader("User-Agent")
	Header.ContentType = c.GetHeader("Content-Type")
	return
}
