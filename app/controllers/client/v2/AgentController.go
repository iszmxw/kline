package v2

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/app/requests"
	"goapi/app/response"
	"goapi/pkg/address"
	"goapi/pkg/echo"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
	"goapi/pkg/validate"
	wss "goapi/pkg/websocket"
	wss_heart_beating "goapi/pkg/wss-heart-beating"
	"math/rand"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 通知公告服务

type AgentController struct {
	BaseController
}

func (h *AgentController) RegisterHandler(c *gin.Context) {
	res := address.GetMessageByIp(c.ClientIP())
	var params requests.RegisterParams
	_ = c.Bind(&params)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	// 检测该用户是否存在系统中
	var queryAppUser response.AppUser
	var CreateAppUser models.AppUser
	var queryAgentUser models.AgentUser
	var AppUserCredit models.AppUserCredit
	// 查询代理
	mysql.DB.Debug().Model(models.AgentUser{}).Where(map[string]interface{}{
		"item_code":     "SMO",
		"referral_code": params.Pid,
	}).Find(&queryAgentUser)
	if len(queryAgentUser.AgentUserId) == 0 {
		echo.Error(c, "FAILD", "：代理不存在")
		return
	}
	mysql.DB.Debug().Model(models.AppUser{}).Where(map[string]interface{}{
		"item_code": "SMO",
		"email":     params.Email,
	}).Find(&queryAppUser)
	if len(queryAppUser.AppUserId) > 0 {
		echo.Error(c, "FAILD", "：用户已存在")
		return
	}
	CreateAppUser.ItemCode = "SMO"
	CreateAppUser.Nation = queryAgentUser.Nation // 国家
	CreateAppUser.RegistNation = res.Country     // 注册时候国家
	CreateAppUser.City = res.City                // 城市
	CreateAppUser.AppUserId = helpers.GetUUID11()
	CreateAppUser.AgentUserId = queryAgentUser.AgentUserId
	CreateAppUser.Mobile = ""
	CreateAppUser.Email = params.Email
	CreateAppUser.Password = params.Password
	CreateAppUser.CreateDate = time.Now()
	CreateAppUser.UpdateDate = time.Now()
	CreateAppUser.BlackList = 0
	CreateAppUser.Balance = 0
	CreateAppUser.BonusBalance = 0
	if strings.EqualFold(queryAgentUser.Nation, "Indonesia") {
		CreateAppUser.VirtualBalance = 1000000.00
		CreateAppUser.CurrencyType = "IDR"
	}
	if strings.EqualFold(queryAgentUser.Nation, "India") {
		CreateAppUser.VirtualBalance = 10000.00
		CreateAppUser.CurrencyType = "INR"
	}
	CreateAppUser.UserLevel = 0
	CreateAppUser.IsNew = 1
	CreateAppUser.Avatar = "https://fileserver.iqoptionproo.com/file/download/0a3d5de4-2eb2-4001-93ea-f7c051ec3e5e.png"
	CreateAppUser.ReferralCode = string([]rune(helpers.GetUUID11())[:6])
	CreateAppUser.RegistIp = c.ClientIP()
	CreateAppUser.HadChangeDecive = 0
	CreateAppUser.UserName = "user"
	CreateAppUser.IsLogout = 0
	CreateAppUser.DeciveBrand = "HTML5" // 设置设备品牌 多个用,号隔开 当前为最后一个
	CreateAppUser.DeviceModel = ""      // 设置设备型号
	DB := mysql.DB
	DB.Begin()
	err := DB.Debug().Model(models.AppUser{}).Create(CreateAppUser).Error
	if err != nil {
		DB.Rollback()
		logger.Error(err)
		echo.Error(c, "FAILD", "：注册失败请稍后再试")
		return
	}
	// 创建关联认证记录
	AppUserCredit.AppUserId = CreateAppUser.AppUserId
	AppUserCredit.CreateDate = time.Now()
	AppUserCredit.TotalStep = 3
	AppUserCredit.ItemCode = "SMO"
	AppUserCredit.IsValid = 1
	AppUserCredit.Nation = CreateAppUser.Nation
	AppUserCredit.CreditStep = 0
	AppUserCredit.EmailCredit = 0
	AppUserCredit.IdentityCredit = 0
	AppUserCredit.IsFinishCredit = 0
	AppUserCredit.MobileCredit = 0
	AppUserCredit.BankCredit = 0
	AppUserCredit.ErrorCreditStatus = 0
	err = DB.Debug().Model(models.AppUserCredit{}).Create(AppUserCredit).Error
	if err != nil {
		DB.Rollback()
		logger.Error(err)
		echo.Error(c, "FAILD", "：注册失败请稍后再试")
		return
	}
	DB.Commit()
	echo.Success(c, nil, "SUCCESS")
}

func (h *AgentController) StatisticsHandler(c *gin.Context) {
	ws, err := wss.GetSocket(c)
	if err != nil {
		logger.Error(err)
		return
	}
	socketId := c.GetHeader("Sec-WebSocket-Key")
	logger.RequestId = socketId
	HeartBeatingKey := fmt.Sprintf("StatisticsHeartBeatingKey:%v", socketId)
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
		var SubStruct requests.SubStatistics
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
		case strings.EqualFold(SubStruct.Id, "01"): // 推送客户总量，今日交易总量(Rp)
			{
				// 开启协诚处理 推送客户总量，今日交易总量(Rp)
				go h.PushStatistics(socketId, mt, ws, SubStruct.Sub)
			}
			break
		case strings.EqualFold(SubStruct.Id, "03"): // 维持心跳
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

// 客户总量在 1524788 开始递增，每20秒随机增加1~2个用户数；

// 印尼：交易总量：每天从0点开始，每秒递增4~7K；（最高一天604,800K大约28万RMB）

// 印度：交易总量：每天从0点开始，每秒递增0.5~0.8k；保留一位小数（最高一天63,675K大约56万RMB）

// 交互规则：

// 1.客户总量刷新时更新；

// 2.今日交易总量实时更新，并有数字滚动效果；
const initTime = 1665504000 // 初始时间戳
const costumer = 1524788    // 初始客户数据

func (h *AgentController) PushStatistics(socketId string, messageType int, ws *wss.WsConn, SubBody requests.SubBody) {
	logger.RequestId = socketId
	currentTime := time.Now()
	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location()).Unix()
	HeartBeatingKey := fmt.Sprintf("StatisticsHeartBeatingKey:%v", socketId)
	// 为了释放for死循环的资源
	for !ws.IsClose {
		if wss_heart_beating.IsDie(HeartBeatingKey) {
			goto STOP
		}
		if len(SubBody.Nation) == 0 {
			goto STOP
		}
		result := cmap.New().Items()
		timeIncr := (time.Now().Unix() - initTime) / 20
		result["endCostumer"] = costumer + timeIncr // 最终客户数
		rand.Seed(time.Now().UnixNano())
		if strings.EqualFold(SubBody.Nation, "Indonesia") {
			min := 4000
			max := 7000
			tradeAmount := (time.Now().Unix() - startTime) * int64(rand.Intn(max-min)+min)
			result["tradeAmount"] = tradeAmount // 最终客户数
		}
		if strings.EqualFold(SubBody.Nation, "India") {
			min := 500
			max := 8000
			tradeAmount := (time.Now().Unix() - startTime) * int64(rand.Intn(max-min)+min)
			result["tradeAmount"] = tradeAmount // 最终客户数
		}
		// 查询用户未读消息，展示小红点
		marshal, err := json.Marshal(gin.H{
			"id":   "02",
			"data": result,
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
		// 2秒推送一次消息
		time.Sleep(time.Second * 2)
		// 返回数据给用户
	}
STOP:
	logger.Error(errors.New("PushStatistics关闭Socket"))
	wss.WsClose(ws)
END:
	logger.Info("PushStatistics已经跳出协诚的循环推送")
}
