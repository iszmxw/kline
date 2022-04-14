package raceoption

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/gostudys/huobi_contract_golang/src/ws/response"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	conf "goapi/pkg/config"
	"goapi/pkg/email"
	"goapi/pkg/esdk"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
	"log"
	"strings"
	"time"
)

func Client(subMsg string) {
	//interrupt := make(chan os.Signal, 1)
	//signal.Notify(interrupt, os.Interrupt)
	//u := url.URL{Scheme: "wss", Host: "trading.raceoption.com", Path: "/api/ws/price"}
	//logger.Info(fmt.Sprintf("connecting to %s", u.String()))

	c, _, err := websocket.DefaultDialer.Dial("wss://trading.raceoption.com/api/ws/price", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	// 关闭 socket
	defer Close(c)
	done := make(chan struct{})
	// 读取消息
	go ReadMsg(done, c)

	// 订阅消息
	err = c.WriteMessage(websocket.TextMessage, []byte(subMsg))
	if err != nil {
		logger.Error(err)
		return
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-done:
			return
			//case t := <-ticker.C:
			// 订阅消息
			//logger.Info(t.String())
			//case <-interrupt:
			//	logger.Info("interrupt")
			//	// Cleanly close the connection by sending a close message and then
			//	// waiting (with timeout) for the server to close the connection.
			//	err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			//	if err != nil {
			//		logger.Info("write close:")
			//		logger.Error(err)
			//		return
			//	}
			//	select {
			//	case <-done:
			//	case <-time.After(time.Second):
			//	}
			//	return
		}
	}
}

func Close(c *websocket.Conn) {
	err := c.Close()
	if err != nil {
		logger.Info(err)
	}
}

func ReadMsg(done chan struct{}, c *websocket.Conn) {
	defer close(done)
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			logger.Error(err)
			return
		}
		var resp Response
		err = json.Unmarshal(message, &resp)
		if err != nil {
			return
		}
		logger.Info(resp)

		ContractCode := strings.ReplaceAll(strings.ToUpper(resp.Symbol), "/", "-")

		var data response.SubIndexKLineResponse
		data.Ch = fmt.Sprintf("market.%v.index.1min", ContractCode)
		data.Ts = helpers.StringToInt64(resp.TimeId)
		data.Tick.Id = helpers.StringToInt64(resp.TimeId)
		data.Tick.Vol = fmt.Sprintf("%v", resp.AskVolume)
		data.Tick.Count = 0
		data.Tick.Open = "0"
		data.Tick.Close = fmt.Sprintf("%v", resp.Strike)
		data.Tick.Low = fmt.Sprintf("%v", resp.Ask)
		data.Tick.High = fmt.Sprintf("%v", resp.Bid)
		data.Tick.Amount = "0"

		var KlineData models.KlineLog
		err = mysql.DB.Debug().
			Model(models.KlineLog{}).
			Select("log_id").
			Where("contract_code", ContractCode).
			Where("upstream_ts", data.Ts).
			Where("kline_source", "raceoption").
			Find(&KlineData).Error
		if err != nil {
			logger.Error(err)
		}
		if KlineData.LogId > 0 {
			return
		}
		lineLogRes := LogData(resp)
		cerr1 := mysql.DB.Debug().Model(models.KlineLog{}).Create(&lineLogRes).Error
		if cerr1 != nil {
			logger.Error(cerr1)
			return
		}

		tips := fmt.Sprintf(
			"新浪时差：系统时间：%v(%v)<br>新浪的 ts 时间：%v(%v)",
			time.Now().UnixMilli(),
			time.Now().Format("2006-01-02 15:04:05"),
			data.Ts,
			time.Unix(data.Ts/1000, 0).Format("2006-01-02 15:04:05"),
		)
		logger.Info(tips)
		shicha := time.Now().UnixMilli() - data.Ts
		logger.Info(fmt.Sprintf("shicha=%v", shicha))
		// 时间差大于 10 秒发送邮件通知
		if shicha/1000 >= 20 {
			go func() {
				Einfo := cmap.New().Items()
				Einfo["title"] = fmt.Sprintf("环境【%v】%v raceoption 服务通知：当前时间差为：%v秒大于 20秒", conf.GetString("app.env"), data.Ch, shicha/1000)
				Einfo["content"] = tips
				err = email.SendEmail(Einfo, conf.GetString("notify.email"))
				if err != nil {
					logger.Error(err)
				}
			}()
		}
		msgData, MarshalJSONErr := json.Marshal(data) //转换成JSON返回的是byte[]
		if MarshalJSONErr != nil {
			logger.Error(MarshalJSONErr)
			return
		}
		_, redisErr := redis.Add(data.Ch, msgData, 0)
		if redisErr != nil {
			logger.Error(redisErr)
			return
		}
		// 收集数据缓存到redis

		// 历史存入 elastic
		putRes := esdk.IndexPut("kline_log", helpers.Int64ToString(KlineData.LogId), KlineData)
		if !putRes {
			logger.Error(errors.New("历史存入 elastic 失败"))
		} else {
			logger.Info("elastic-success")
		}

	}

}

func LogData(resp Response) models.KlineLog {
	var klineLog models.KlineLog
	klineLog.ContractCode = strings.ReplaceAll(strings.ToUpper(resp.Symbol), "/", "-") // 合约代码 BTC-USD
	klineLog.Ch = fmt.Sprintf("market.%v.index.1min", klineLog.ContractCode)           // kilne返回日期时间戳
	klineLog.Ts = time.Now().UnixMilli()                                               // 自己生成的ts
	klineLog.UpstreamTs = helpers.StringToInt64(resp.TimeId)                           // 上游的ts
	klineLog.ID = helpers.StringToInt64(resp.TimeId)                                   // 自己生成的ts/1000
	klineLog.Mrid = 0.00                                                               // ID
	klineLog.COpen = 0                                                                 // 起始峰值
	klineLog.CClose = resp.Strike                                                      // 结束峰值
	klineLog.High = resp.Bid                                                           // 最高峰值
	klineLog.Low = resp.Ask                                                            // 最低峰值
	klineLog.Amount = 0                                                                // 成交金额 TODO
	klineLog.Vol = helpers.StrToFloat64(resp.AskVolume)                                // 成交手 TODO
	klineLog.CCount = 0                                                                // 交易笔数 TODO
	klineLog.TsDate = time.Unix(0, time.Now().UnixMilli()*int64(time.Millisecond))     // Kline返回日期
	klineLog.CreateDate = time.Now()                                                   // 创建日期
	klineLog.KlineSource = "raceoption"                                                // K线来源新浪
	return klineLog
}
