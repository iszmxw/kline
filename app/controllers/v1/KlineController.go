package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/app/models"
	"goapi/app/requests"
	"goapi/app/response"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/output"
	"goapi/pkg/redis"
	"goapi/pkg/redis_socket"
	"goapi/pkg/websocket"
	"strings"
	"time"
)

// k线图服务

type KlineController struct {
	BaseController
}

// WsIndexHandler socket 负责转发redis缓存的socket数据
func (h *KlineController) WsIndexHandler(c *gin.Context) {
	RequestId := logger.RequestId
	ws, err := wss.GetSocket(c)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("客户端建立socket成功")
	// 为了释放for死循环的资源
	for {
		// 读取用户客户端ws中订阅的 topic
		mt, message, ReadMessageErr := ws.Conn.ReadMessage()
		if ReadMessageErr != nil {
			logger.Error(ReadMessageErr)
			marketErr := ws.Conn.Close()
			if marketErr != nil {
				logger.Error(marketErr)
			} else {
				logger.Info("关闭成功")
			}
			// 跳出for关闭后面的读取
			break
		}
		//对数据进行切割，读取参数
		//如果请求的是 "market.ETH-USD.index.1min" ,订阅这条信息，然后再返回
		var SubStruct requests.Sub
		SubErr := json.Unmarshal(message, &SubStruct)
		if SubErr != nil {
			logger.Error(SubErr)
			ws.Conn.Close()
			break
		}
		msg := SubStruct.Sub
		logger.Info(SubStruct)
		// 2 小时内一个长连接禁止重复订阅相同的 topic
		checkMsg := "socket_request:" + RequestId + ":" + msg
		if redis.CheckExist(checkMsg) {
			logger.Info("该订阅已存在，请勿重新订阅")
			continue
		}
		_, redisAddErr := redis.Add(checkMsg, msg, 30*60) // 半个小时内不允许重复请求
		if redisAddErr != nil {
			logger.Error(redisAddErr)
			break
		}
		// 24 小时内一个长连接禁止重复订阅相同的 topic

		// 开启协诚处理推送消息
		go func(ws *wss.WsConn, msg string) {
			// 为了释放for死循环的资源
			for {
				push, getErr := redis_socket.Get(msg)
				if getErr != nil {
					logger.Error(getErr)
					return
				}
				// 空消息跳过
				if len(push) <= 0 {
					continue
				}
				var WsIndex *response.WsIndex
				WsIndexErr := json.Unmarshal([]byte(push), &WsIndex)
				if WsIndexErr != nil {
					logger.Error(errors.New("Umarshal failed:" + WsIndexErr.Error()))
					logger.Error(WsIndexErr)
					continue
				}
				WsIndexJson, WsIndexJsonErr := json.Marshal(WsIndex)
				if WsIndexJsonErr != nil {
					logger.Error(WsIndexJsonErr)
					continue
				}
				// 1秒推送一次消息
				time.Sleep(time.Second * 1)
				ws.Mux.Lock()
				WriteMessageErr := ws.Conn.WriteMessage(mt, WsIndexJson)
				ws.Mux.Unlock()
				if WriteMessageErr != nil {
					ws.Conn.Close()
					// 跳出for关闭后面的推送
					break
				}
				// 返回数据给用户
			}
			logger.Info("已经跳出协诚的循环推送")
		}(ws, msg)
	}
	logger.Info("socket被关闭，本次链接结束")
}

// HistoryHandler 历史行情
func (h *KlineController) HistoryHandler(c *gin.Context) {
	//beginDate := time.Now().Format("2006-01-02 15:04:05")
	symbol := c.Query("symbol")
	period := c.Query("period")
	beginDate := c.Query("beginDate")
	endDate := c.Query("endDate")
	if len(beginDate) <= 0 || helpers.StringToInt(beginDate) <= 0 {
		output.Rjson(c, nil, "beginDate 起始交易日期必填", 0)
		return
	}
	if len(endDate) <= 0 || helpers.StringToInt(endDate) <= 0 {
		output.Rjson(c, nil, "endDate 截止交易日期必填", 0)
		return
	}
	if len(symbol) <= 0 || len(period) <= 0 {
		symbol = "BTC-USD"
		period = "1min"
	}
	var data []response.KlineLog
	ch := fmt.Sprintf("market.%s.index.%s", strings.ToUpper(symbol), period)
	DB := mysql.DB.Debug()
	// 一次取出历史记录最新三百条
	KlineDB := DB.Model(models.KlineLog{}).
		Where("ch", ch).
		Where("ts >= ?", beginDate).
		Where("ts <= ?", endDate).
		Order("log_id desc")
	var Count int64 // 查询总条数
	KlineDB.Count(&Count)
	if Count > 2000 {
		Count = 2000
	}
	err := KlineDB.Limit(int(Count)).Find(&data).Error
	if err != nil {
		logger.Error(err)
		output.Rjson(c, nil, "no data", 0)
		return
	}
	output.Rjson(c, gin.H{
		"HistKlineLogs": data,
		"beginDate":     time.Unix(helpers.StringToInt64(beginDate)/1000, 0).Format("2006-01-02 15:04:05"),
		"endDate":       time.Unix(helpers.StringToInt64(endDate)/1000, 0).Format("2006-01-02 15:04:05"),
		"klineSource":   "huobi",
		"totals":        Count,
		"contractCode":  symbol,
	}, "success", 1)
}
