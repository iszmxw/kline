package v2

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/app/requests"
	"goapi/app/response"
	"goapi/pkg/config"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
	wss "goapi/pkg/websocket"
	wssHeartBeating "goapi/pkg/wss-heart-beating"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// 通知公告服务

type NoticeController struct {
	BaseController
}

func (h *NoticeController) StationLetterHandler(c *gin.Context) {
	ws, err := wss.GetSocket(c)
	if err != nil {
		logger.Error(err)
		return
	}
	socketId := c.GetHeader("Sec-WebSocket-Key")
	logger.RequestId = socketId
	HeartBeatingKey := fmt.Sprintf("NoticeHeartBeatingKey:%v", socketId)
	_, _ = redis.Client.Add(HeartBeatingKey, time.Now().Format("2006-01-02 15:04:05"), 300) // 5分钟过期
	// 为了释放for死循环的资源
	for !ws.IsClose {
		// 读取用户客户端ws中订阅的 topic
		mt, message, ReadMessageErr := ws.Conn.ReadMessage()
		if ReadMessageErr != nil {
			logger.Error(ReadMessageErr)
			// 跳出for关闭后面的读取
			goto END
		}
		//对数据进行切割，读取参数
		//如果请求的是 "market.ETH-USD.index.1min" ,订阅这条信息，然后再返回
		var SubStruct requests.SubStationLetter
		SubErr := json.Unmarshal(message, &SubStruct)
		if SubErr != nil {
			logger.Error(SubErr)
			goto END
		}
		switch true {
		case ws.IsClose:
			{
				goto END
			}
		case strings.EqualFold(SubStruct.Id, "02"): // 推送小红点消息
			{
				// 开启协诚处理 小红点消息通知
				go h.PushNotice(socketId, mt, ws, SubStruct.Sub)
			}
			break
		case strings.EqualFold(SubStruct.Id, "03"):
			{
				// 处理心跳数据
				currentTime := time.Now().Format("2006-01-02 15:04:05")
				_, _ = redis.Client.Add(HeartBeatingKey, currentTime, 300) // 5分钟过期
				ws.Mux.Lock()
				WriteMessageErr := ws.Conn.WriteMessage(websocket.TextMessage, []byte(`{"id":"03","data":"pong"}`))
				ws.Mux.Unlock()
				if WriteMessageErr != nil {
					logger.Error(WriteMessageErr)
					wss.WsClose(ws)
				}
				logger.Info(fmt.Sprintf("当前心跳时间：%v", currentTime))
			}
			break
		default:
			break
		}
		time.Sleep(time.Millisecond * 10) // 在接受的white循环里面停顿10毫秒
	}
END:
	wss.WsClose(ws)
	logger.Info("socket被关闭，本次链接结束")
}

// 小红点消息通知

func (h *NoticeController) PushNotice(socketId string, messageType int, ws *wss.WsConn, SubToken requests.SubToken) {
	logger.RequestId = socketId
	HeartBeatingKey := fmt.Sprintf("NoticeHeartBeatingKey:%v", socketId)
	// {"token":"123456","decive_brand":"apple"}
	if len(SubToken.Token) == 0 {
		wss.WsClose(ws)
		return
	}
	// 为了释放for死循环的资源
	for !ws.IsClose {
		if wssHeartBeating.IsDie(HeartBeatingKey) {
			goto STOP
		}
		strInfo, err := redis.Client.DefaultGet(SubToken.Token)
		if err != nil {
			logger.Error(err)
			goto STOP
		}
		// json字符串数组,转换成切片
		var user response.ClientUserBeans
		var NoticeInfo response.NoticeInfo
		multiErr := json.Unmarshal([]byte(strInfo), &user)
		if multiErr != nil {
			logger.Error(errors.New("转换出错"))
			logger.Error(multiErr)
			goto STOP
		}
		logger.Info(user)
		result := cmap.New().Items()
		// 查询用户未读消息，展示小红点
		err = mysql.DB.Debug().Model(models.OffMailSendLog{}).Where(
			map[string]interface{}{
				"item_code":    config.GetString("app.item_code"),
				"nation":       user.Nation,
				"decive_brand": SubToken.DeciveBrand,
			}).Where("app_user_id = 'ALL' OR app_user_id = ?", user.AppUserId).
			Select("count(1) AS total").
			Find(&result).Error
		if err != nil {
			_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10409","data":"no data"}`))
			logger.Error(err)
			return
		}
		// map转struct
		err = mapstructure.Decode(result, &NoticeInfo)
		if err != nil {
			ws.Mux.Lock()
			_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10409","data":"no data"}`))
			ws.Mux.Unlock()
			logger.Error(err)
			return
		}
		marshal, err := json.Marshal(gin.H{
			"id":   "02",
			"data": NoticeInfo,
		})
		if err != nil {
			_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10409","data":"no data"}`))
			logger.Error(err)
			return
		}
		ws.Mux.Lock()
		WriteMessageErr := ws.Conn.WriteMessage(messageType, marshal)
		ws.Mux.Unlock()
		if WriteMessageErr != nil {
			wss.WsClose(ws)
			// 跳出for关闭后面的推送
			goto END
		}
		// 5秒推送一次消息
		time.Sleep(time.Second * 5)
		// 返回数据给用户
	}
STOP:
	logger.Error(errors.New("PushNotice关闭Socket"))
	wss.WsClose(ws)
END:
	logger.Info("PushNotice已经跳出协诚的循环推送")
}
