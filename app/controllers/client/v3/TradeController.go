package v3

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"goapi/app/models"
	"goapi/app/requests"
	goadjust "goapi/pkg/adjust"
	"goapi/pkg/config"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	wss "goapi/pkg/websocket"
	"goapi/service/binary-monitoring/pkg"
	"strings"
	"time"
)

// 交易模块

type TradeController struct {
}

func getMessage(languageCode, nation, resultCode string) string {
	var AppSelectEnum models.AppSelectEnum
	var msg string
	if len(languageCode) == 0 {
		languageCode = "EN"
	}
	if len(nation) == 0 {
		nation = "India"
	}
	if len(resultCode) == 0 {
		resultCode = "ErrorRequest"
	}
	err := mysql.DB.Debug().Model(models.AppSelectEnum{}).
		Where("language_code", languageCode).
		Where("nation", nation).
		Where("result_code", resultCode).Find(&AppSelectEnum).Error
	if err != nil {
		msg = fmt.Sprintf(`{"id":"10409","msg":"%v"}`, "Unknown Error")
	}
	msg = fmt.Sprintf(`{"id":"10409","msg":"%v"}`, AppSelectEnum.Message)
	return msg
}

// 提交下单

func (h *TradeController) SubmitTrade(ws *wss.WsConn, params requests.SubmitTradeParams) {
	logger.RequestId = ws.SecWebSocketKey
	if params.Amounts <= 0 {
		SocketResp(ws, websocket.TextMessage, getMessage(params.LanguageCode, params.Nation, "ParamsCanNotBeNull"))
		return
	}
	if params.AmountType <= 0 {
		SocketResp(ws, websocket.TextMessage, getMessage(params.LanguageCode, params.Nation, "ParamsCanNotBeNull"))
		return
	}
	if len(params.AssetsId) == 0 {
		SocketResp(ws, websocket.TextMessage, getMessage(params.LanguageCode, params.Nation, "ParamsCanNotBeNull"))
		return
	}
	if len(params.TradingBeginDate) <= 0 {
		SocketResp(ws, websocket.TextMessage, getMessage(params.LanguageCode, params.Nation, "ParamsCanNotBeNull"))
		return
	}
	if len(params.TradeEndDate) <= 0 {
		SocketResp(ws, websocket.TextMessage, getMessage(params.LanguageCode, params.Nation, "ParamsCanNotBeNull"))
		return
	}
	if len(params.UpDownType) == 0 {
		SocketResp(ws, websocket.TextMessage, getMessage(params.LanguageCode, params.Nation, "ParamsCanNotBeNull"))
		return
	}
	if params.BeginAssetsPrice <= 0 {
		SocketResp(ws, websocket.TextMessage, getMessage(params.LanguageCode, params.Nation, "ParamsCanNotBeNull"))
		return
	}
	if params.KlineTime <= 0 {
		SocketResp(ws, websocket.TextMessage, getMessage(params.LanguageCode, params.Nation, "ParamsCanNotBeNull"))
		return
	}
	sysTime := time.Now()
	// 请求超时拦截
	if helpers.StringToInt64(params.TradingBeginDate) < sysTime.UnixMilli()-10000 && 3 != params.AmountType {
		SocketResp(ws, websocket.TextMessage, getMessage(params.LanguageCode, params.Nation, "RequestTimeout"))
		return
	}
	// 时间超前拦截
	if helpers.StringToInt64(params.TradingBeginDate) > sysTime.UnixMilli() && 3 != params.AmountType {
		SocketResp(ws, websocket.TextMessage, getMessage(params.LanguageCode, params.Nation, "PleaseCheckTheNetworkTime"))
		return
	}

	// 双重锁
	l := ws.List.Back()
	if l != nil {
		ws.List.Remove(l)
		up := l.Value.(int64)
		sysT := sysTime.UnixMilli()
		if sysT-up < 100 {
			msg := fmt.Sprintf("检测到用户【%s】频繁下单,klineTime=%v", params.AppUserId, params.KlineTime)
			logger.Info(msg)
			pkg.Notice(msg)
			SocketResp(ws, websocket.TextMessage, getMessage(params.LanguageCode, params.Nation, "FailOperate"))
			return
		}
	}
	//ws.List.PushFront(sysTime.UnixMilli())
	//ok, _ := redis.Client.SelectDbSetNx(config.GetInt("redis.db"), fmt.Sprintf("SubmitTrade:%v", sysTime.UnixMilli()/100), "suo", 10)
	//if !ok {
	//	pkg.Notice(fmt.Sprintf("检测到用户【%s】频繁下单,klineTime=%v", params.AppUserId, params.KlineTime))
	//	SocketResp(ws, websocket.TextMessage, getMessage(params.LanguageCode, params.Nation, "FailOperate"))
	//	return
	//}
	// 双重锁

	// demo 订单只能同时存在五笔
	if 3 == params.AmountType {
		var total int64
		err := mysql.DB.Debug().
			Raw(`SELECT COUNT(1) FROM t_trading_order_virtual tov WHERE tov.quiz_result = -1 AND tov.app_user_id = ?`, params.AppUserId).
			Scan(&total).Error
		if err != nil {
			logger.Error(err)
			SocketResp(ws, websocket.TextMessage, getMessage(params.LanguageCode, params.Nation, "FailOperate"))
			return
		}
		if total >= 5 {
			SocketResp(ws, websocket.TextMessage, getMessage(params.LanguageCode, params.Nation, "DemoLimit"))
			return
		}
	}
	// 校验是否存在该交易点数据
	var size int64
	err := mysql.DB.Debug().
		Raw(`SELECT COUNT(tkl.c_close) FROM t_kline_log tkl,t_trading_assets tta WHERE tkl.contract_code = tta.assets_code AND tkl.ts = ? AND tta.id = ? AND tkl.c_close = ?`, params.KlineTime, params.AssetsId, params.BeginAssetsPrice).
		Scan(&size).Error
	if err != nil {
		logger.Error(err)
		SocketResp(ws, websocket.TextMessage, getMessage(params.LanguageCode, params.Nation, "FailOperate"))
		return
	}
	if size == 0 {
		logger.Error(errors.New("校验是否存在该交易点数据"))
		SocketResp(ws, websocket.TextMessage, getMessage(params.LanguageCode, params.Nation, "FailOperate"))
		return
	}
	// 校验资产 资产已经休市，请更换其他资产
	var asset models.TradingAssets
	mysql.DB.Debug().Model(models.TradingAssets{}).Where("id", params.AssetsId).Find(&asset)
	if asset.IsValid != 1 {
		SocketResp(ws, websocket.TextMessage, getMessage(params.LanguageCode, params.Nation, "AssetClosed"))
		return
	}
	// 前端传递费率和数据库读取到的不一致
	if asset.ReturnRate != params.ReturnRate {
		logger.Info("参数错误：returnRate不一致：", asset.ReturnRate, "---", params.ReturnRate)
		// 发送通知
		pkg.Notice("下单风控检测出现 returnRate 不一致，详细信息", asset, params)
		//SocketResp(ws, websocket.TextMessage, getMessage(params.LanguageCode, params.Nation, "FailOperate"))
		//return
	}
	var appUser models.AppUser
	mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", params.AppUserId).Find(&appUser)
	// 进行交易规则校验，虚拟交易不进行校验
	if params.AmountType != 3 {
		ok, msg := tradeJudge(appUser, params)
		if ok {
			SocketResp(ws, websocket.TextMessage, msg)
			return
		}
	}
	// 交易金额类型 1 余额 2 bonus 3 虚拟货币
	orderId := fmt.Sprintf("%v%v", time.Now().Format("0060102"), uuid.New().ID())
	marshal, err := json.Marshal(gin.H{
		"id":  "07",
		"msg": "ok",
		"data": gin.H{
			"tradeOrderId": orderId,
		},
	})
	if err != nil {
		logger.Error(err)
		SocketResp(ws, websocket.TextMessage, getMessage(params.LanguageCode, params.Nation, "FailOperate"))
		return
	}
	switch params.AmountType {
	case 1:
		if appUser.Balance < params.Amounts {
			SocketResp(ws, websocket.TextMessage, getMessage(params.LanguageCode, params.Nation, "InsufficientBalance"))
			return
		}
		ok := createTradingOrder(orderId, params, appUser, asset)
		if ok {
			goadjust.SubmitEvent(goadjust.EventParams{
				EventName:        "trade",
				AppUserId:        appUser.AppUserId,
				IsRepeatedSubmit: "0",
				Mobile:           appUser.Mobile,
				EventBindId:      appUser.FlowBindId,
				Email:            appUser.Email,
				FlowJsonString:   appUser.FlowJson,
				DeviceBrand:      appUser.DeciveBrand,
			}, appUser.Nation, "adjust")
			// 休息两百毫秒再返回
			//time.Sleep(time.Millisecond * 200)
			SocketResp(ws, websocket.TextMessage, string(marshal))
		} else {
			SocketResp(ws, websocket.TextMessage, getMessage(params.LanguageCode, params.Nation, "FailOperate"))
		}
		break
	case 2:
		if appUser.BonusBalance < params.Amounts {
			SocketResp(ws, websocket.TextMessage, getMessage(params.LanguageCode, params.Nation, "InsufficientBalance"))
			return
		}
		ok := createTradingOrder(orderId, params, appUser, asset)
		if ok {
			goadjust.SubmitEvent(goadjust.EventParams{
				EventName:        "trade",
				AppUserId:        appUser.AppUserId,
				IsRepeatedSubmit: "0",
				Mobile:           appUser.Mobile,
				EventBindId:      appUser.FlowBindId,
				Email:            appUser.Email,
				FlowJsonString:   appUser.FlowJson,
				DeviceBrand:      appUser.DeciveBrand,
			}, appUser.Nation, "adjust")
			// 休息两百毫秒再返回
			//time.Sleep(time.Millisecond * 200)
			SocketResp(ws, websocket.TextMessage, string(marshal))
		} else {
			SocketResp(ws, websocket.TextMessage, getMessage(params.LanguageCode, params.Nation, "FailOperate"))
		}
		break
	case 3:
		if appUser.VirtualBalance < params.Amounts {
			SocketResp(ws, websocket.TextMessage, getMessage(params.LanguageCode, params.Nation, "InsufficientBalance"))
			return
		}
		createVirtualTradingOrder(orderId, params, appUser, asset)
		// 休息一百毫秒再返回
		//time.Sleep(time.Millisecond * 100)
		SocketResp(ws, websocket.TextMessage, string(marshal))
		break
	default:
		break
	}
}

// 交易规则判断

func tradeJudge(appUser models.AppUser, params requests.SubmitTradeParams) (bool, string) {
	// 获取用户所在国家的时区
	var appItem models.AppItem
	err := mysql.DB.Debug().Model(models.AppItem{}).Where("item_code", appUser.ItemCode).Where("item_national_name", appUser.Nation).Find(&appItem).Error
	if err != nil {
		logger.Error(err)
		logger.Error(errors.New("交易规则未配置"))
		return true, getMessage(params.LanguageCode, params.Nation, "FailOperate")
	}
	timeZoneArr := strings.SplitAfter(appItem.TimeZone, "T")
	timeZone := timeZoneArr[1]
	var userTradeMessage map[string]interface{}
	err = mysql.DB.Debug().Raw(`SELECT a.app_user_id AS 'appUserId', a.is_new AS 'isNew', a.black_list AS 'blackList',
ifnull(b.tradingCount, 0) AS 'tradingCount', ifnull(b.tradingAmounts, 0) AS 'tradingAmounts',
ifnull(c.rechargeAmounts, 0) AS 'rechargeAmounts', ifnull(c.rechargeCount, 0) AS 'rechargeCount',
ifnull(c.rechargeTodayAmounts, 0) AS 'rechargeTodayAmounts', ifnull(c.rechargeTodayCount, 0) AS 'rechargeTodayCount',
ifnull(d.withdrawAmounts, 0) AS 'withdrawAmounts', ifnull(d.withdrawCount, 0) AS 'withdrawCount',
ifnull(d.withdrawTodayAmounts, 0) AS 'withdrawTodayAmounts', ifnull(d.withdrawTodayCount, 0) AS 'withdrawTodayCount',
ifnull(e.hisTradingCount, 0) AS 'hisTradingCount', ifnull(e.hisTradingAmounts, 0) AS 'hisTradingAmounts'
FROM (SELECT a.app_user_id, '1' AS 'biao', a.is_new, a.black_list
FROM t_app_user AS a
WHERE 1=1 AND a.app_user_id = ? AND a.nation = ?) AS a
INNER JOIN (SELECT COUNT(1) AS 'tradingCount', SUM(a.amounts) AS 'tradingAmounts', '1' AS 'biao'
FROM t_trading_order AS a
WHERE 1=1 AND a.app_user_id = ?
AND DATE_FORMAT(CONVERT_TZ(a.create_date, '+07:00', ?),'%Y-%m-%d') = DATE_FORMAT(CONVERT_TZ(NOW(), '+07:00', ?),'%Y-%m-%d')
AND a.amount_type = 1 AND a.nation = ?) AS b ON a.biao = b.biao
INNER JOIN (SELECT SUM(b.amount) AS 'rechargeAmounts', COUNT(1) AS 'rechargeCount',
SUM(CASE WHEN DATE_FORMAT(CONVERT_TZ(b.create_date, '+07:00', ?),'%Y-%m-%d') = DATE_FORMAT(CONVERT_TZ(NOW(), '+07:00', ?),'%Y-%m-%d') THEN b.amount ELSE 0 END ) AS 'rechargeTodayAmounts',
SUM(CASE WHEN DATE_FORMAT(CONVERT_TZ(b.create_date, '+07:00', ?),'%Y-%m-%d') = DATE_FORMAT(CONVERT_TZ(NOW(), '+07:00', ?),'%Y-%m-%d') THEN 1 ELSE 0 END ) AS 'rechargeTodayCount',
'1' AS 'biao'
FROM t_recharge_order AS b
WHERE 1=1 AND b.app_user_id = ? AND b.recharge_order_status = 1 AND b.nation = ?) AS c ON a.biao = c.biao
INNER JOIN (SELECT SUM(b.amount) AS 'withdrawAmounts', COUNT(1) AS 'withdrawCount',
SUM(CASE WHEN DATE_FORMAT(CONVERT_TZ(b.create_date, '+07:00', ?),'%Y-%m-%d') = DATE_FORMAT(CONVERT_TZ(NOW(), '+07:00', ?),'%Y-%m-%d') THEN b.amount ELSE 0 END ) AS 'withdrawTodayAmounts',
SUM(CASE WHEN DATE_FORMAT(CONVERT_TZ(b.create_date, '+07:00', ?),'%Y-%m-%d') = DATE_FORMAT(CONVERT_TZ(NOW(), '+07:00', ?),'%Y-%m-%d') THEN 1 ELSE 0 END ) AS 'withdrawTodayCount',
'1' AS 'biao'
FROM t_withdraw_order AS b
WHERE 1=1 AND b.app_user_id = ? AND b.withdraw_order_status IN (1,2) AND b.nation = ?) AS d ON a.biao = d.biao
INNER JOIN (SELECT COUNT(1) AS 'hisTradingCount', SUM(a.amounts) AS 'hisTradingAmounts', '1' AS 'biao'
FROM t_trading_order AS a
WHERE 1=1 AND a.app_user_id = ? AND a.amount_type = 1 AND a.nation = ?) AS e ON a.biao = e.biao`, appUser.AppUserId, appUser.Nation, appUser.AppUserId, timeZone, timeZone, appUser.Nation, timeZone, timeZone, timeZone, timeZone, appUser.AppUserId, appUser.Nation, timeZone, timeZone, timeZone, timeZone, appUser.AppUserId, appUser.Nation, appUser.AppUserId, appUser.Nation).Scan(&userTradeMessage).Error
	if err != nil {
		logger.Error(err)
		return true, getMessage(params.LanguageCode, params.Nation, "FailOperate")
	}
	logger.Info(fmt.Sprintf("%T", userTradeMessage))
	logger.Info(userTradeMessage["tradingCount"])
	logger.Info(userTradeMessage["tradingAmounts"])
	// 查询用户交易规则
	var userOperationRules []models.UserOperationRules
	mysql.DB.Debug().Model(models.UserOperationRules{}).
		Where("is_valid", 1).
		Where("is_show", 1).
		Where("rule_type_code", "trading").
		Where("nation", appUser.Nation).
		Where("nation", appUser.Nation).
		Find(&userOperationRules)
	for _, item := range userOperationRules {
		// 最小交易金额限制，命中规则拦截
		if strings.EqualFold("minTradingAmounts", item.RuleCode) {
			if DecisionAlgorithmInterface(int(params.Amounts), int(item.Value), item.CalculateMethod) {
				return true, fmt.Sprintf(`{"id":"10409","msg":"%v"}`, strings.ReplaceAll(item.ValueTips, "[value]", fmt.Sprintf("[%v]", item.Value)))
			}
		}
		// 最大交易金额限制，命中规则拦截
		if strings.EqualFold("maxTradingAmounts", item.RuleCode) {
			if DecisionAlgorithmInterface(int(params.Amounts), int(item.Value), item.CalculateMethod) {
				return true, fmt.Sprintf(`{"id":"10409","msg":"%v"}`, strings.ReplaceAll(item.ValueTips, "[value]", fmt.Sprintf("[%v]", item.Value)))
			}
		}
		// 最多交易次数限制，命中规则拦截
		if strings.EqualFold("maxTradingTimes", item.RuleCode) {
			if DecisionAlgorithmInterface(int(userTradeMessage["tradingCount"].(int64)), int(item.Value), item.CalculateMethod) {
				return true, fmt.Sprintf(`{"id":"10409","msg":"%v"}`, strings.ReplaceAll(item.ValueTips, "[value]", fmt.Sprintf("[%v]", item.Value)))
			}
		}
		// 日交易总金额限制，命中规则拦截【不判断影响不大】
		if strings.EqualFold("dayMaxTradingAmounts", item.RuleCode) {
			if DecisionAlgorithmInterface(int(userTradeMessage["tradingAmounts"].(float64)), int(item.Value), item.CalculateMethod) {
				return true, fmt.Sprintf(`{"id":"10409","msg":"%v"}`, strings.ReplaceAll(item.ValueTips, "[value]", fmt.Sprintf("[%v]", item.Value)))
			}
		}
	}
	// 动态统计数据
	return false, ""
}

// 决策运算算法工具类

func DecisionAlgorithmInterface(a, b, calculateMethod int) bool {
	flag := false
	// 运算方式 1> 2< 3>= 4<= 5<> 6in 7equals 8=
	switch calculateMethod {
	case 1:
		flag = a > b
		break
	case 2:
		flag = a < b
		break
	case 3:
		flag = a >= b
		break
	case 4:
		flag = a <= b
		break
	case 5:
		flag = a != b
		break
	case 6:
		flag = a == b
		break
	case 7:
		flag = a == b
		break
	case 8:
		flag = a == b
		break
	default:
		break
	}
	return flag
}

// 创建真实交易订单
func createTradingOrder(orderId string, params requests.SubmitTradeParams, appUser models.AppUser, asset models.TradingAssets) bool {
	Pass := true
	createOrder := new(models.TradingOrder)
	createOrder.TradeOrderId = orderId
	createOrder.Amounts = params.Amounts
	createOrder.AmountType = params.AmountType
	createOrder.AppUserId = appUser.AppUserId
	createOrder.AssetsId = asset.Id
	createOrder.AssetsLogoUrl = asset.AssetsLogoUrl
	createOrder.AssetsName = asset.AssetsName
	createOrder.AssetsCode = asset.AssetsCode
	createOrder.PlanGainAmounts = (1 + asset.ReturnRate) * params.Amounts
	createOrder.ReturnRate = asset.ReturnRate
	beginAssetsPrice, err := helpers.FormatFloat(params.BeginAssetsPrice, asset.PointNumber)
	if err != nil {
		logger.Error(err)
		return false
	}
	createOrder.BeginAssetsPrice = beginAssetsPrice
	createOrder.CreateDate = time.Now()
	createOrder.UpdateDate = time.Now()
	createOrder.CreateTs = time.Now().UnixMilli()
	createOrder.UpdateTs = time.Now().UnixMilli()
	createOrder.IsValid = 1
	createOrder.ItemCode = config.GetString("app.item_code")
	createOrder.QuizResult = -1
	switch params.AmountType {
	case 1:
		createOrder.RealCostAmounts = params.Amounts
		createOrder.Balance = appUser.Balance
		break
	case 2:
		createOrder.RealCostAmounts = 0
		createOrder.Balance = appUser.BonusBalance
		break
	case 3:
		createOrder.RealCostAmounts = 0
		createOrder.Balance = appUser.VirtualBalance
		break
	default:
		break
	}
	createOrder.RealGainAmounts = 0
	createOrder.Nation = appUser.Nation
	createOrder.CurrencyType = appUser.CurrencyType
	createOrder.EqualReturnAmounts = 0
	createOrder.TradeBeginDate = time.Unix(0, helpers.StringToInt64(params.TradingBeginDate)*int64(time.Millisecond))
	createOrder.TradeEndDate = time.Unix(0, helpers.StringToInt64(params.TradeEndDate)*int64(time.Millisecond))
	createOrder.KlineTime = params.KlineTime
	createOrder.TradeBeginTimestamp = createOrder.TradeBeginDate.UnixMilli()
	createOrder.TradeEndTimestamp = createOrder.TradeEndDate.UnixMilli()
	createOrder.UpDownType = params.UpDownType
	createOrder.LockingKey = orderId
	err = mysql.DB.Debug().Model(models.TradingOrder{}).Create(&createOrder).Error
	if err != nil {
		logger.Error(err)
		return false
	}
	// 变更用户余额
	switch params.AmountType {
	case 1:
		Pass = UpdateUserBalance(orderId, appUser.AppUserId, "2", params.Amounts, 2)
		break
	case 2:
		Pass = UpdateUserBonusBalance(orderId, appUser.AppUserId, "2", params.Amounts, 7)
		break
	case 3:
		err = mysql.DB.Model(models.AppUser{}).Where("app_user_id", appUser.AppUserId).Updates(map[string]interface{}{
			"virtual_balance": appUser.VirtualBalance - params.Amounts,
		}).Error
		if err != nil {
			logger.Error(err)
			Pass = false
		}
		break
	default:
		break
	}
	return Pass
}

// 创建虚拟交易订单
func createVirtualTradingOrder(orderId string, params requests.SubmitTradeParams, appUser models.AppUser, asset models.TradingAssets) bool {
	createOrder := new(models.TradingOrderVirtual)
	createOrder.TradeOrderId = orderId
	createOrder.Amounts = params.Amounts
	createOrder.AmountType = params.AmountType
	createOrder.AppUserId = appUser.AppUserId
	createOrder.AssetsId = asset.Id
	createOrder.AssetsLogoUrl = asset.AssetsLogoUrl
	createOrder.AssetsName = asset.AssetsName
	createOrder.AssetsCode = asset.AssetsCode
	createOrder.PlanGainAmounts = (1 + asset.ReturnRate) * params.Amounts
	createOrder.ReturnRate = asset.ReturnRate
	beginAssetsPrice, err := helpers.FormatFloat(params.BeginAssetsPrice, asset.PointNumber)
	if err != nil {
		logger.Error(err)
		return false
	}
	createOrder.BeginAssetsPrice = beginAssetsPrice
	createOrder.CreateDate = time.Now()
	createOrder.UpdateDate = time.Now()
	createOrder.CreateTs = time.Now().UnixMilli()
	createOrder.UpdateTs = time.Now().UnixMilli()
	createOrder.IsValid = 1
	createOrder.ItemCode = config.GetString("app.item_code")
	createOrder.QuizResult = -1
	createOrder.RealCostAmounts = 0
	createOrder.EqualReturnAmounts = 0
	createOrder.TradeBeginDate = time.Unix(0, helpers.StringToInt64(params.TradingBeginDate)*int64(time.Millisecond))
	createOrder.TradeEndDate = time.Unix(0, helpers.StringToInt64(params.TradeEndDate)*int64(time.Millisecond))
	createOrder.TradeBeginTimestamp = createOrder.TradeBeginDate.UnixMilli()
	createOrder.TradeEndTimestamp = createOrder.TradeEndDate.UnixMilli()
	createOrder.UpDownType = params.UpDownType
	createOrder.CurrencyType = appUser.CurrencyType
	createOrder.Nation = appUser.Nation
	createOrder.VirtualBalance = appUser.VirtualBalance
	err = mysql.DB.Debug().Model(models.TradingOrderVirtual{}).Create(&createOrder).Error
	if err != nil {
		logger.Error(err)
		return false
	}
	// 变更用户余额
	err = mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", appUser.AppUserId).Updates(map[string]interface{}{
		"virtual_balance": appUser.VirtualBalance - params.Amounts,
	}).Error
	if err != nil {
		logger.Error(err)
		return false
	}
	return true
}

type UserBalance struct {
	Balance      float64 `json:"balance"`
	BonusBalance float64 `json:"bonus_balance"`
}

func UpdateUserBalance(sourceIndex, appUserId, types string, amount float64, changeType int) bool {
	DB := mysql.DB.Begin().Debug()
	var userBalance UserBalance
	err := DB.Raw(`SELECT balance,bonus_balance FROM t_app_user WHERE app_user_id = ? FOR UPDATE`, appUserId).Scan(&userBalance).Error
	if err != nil {
		logger.Error(err)
		DB.Rollback()
		return false
	}
	balance := userBalance.Balance
	var changeBalance float64 = 0
	if strings.EqualFold("1", types) {
		balance = balance + amount
		changeBalance = amount
		logger.Info("增加余额", appUserId, changeBalance)
	} else if strings.EqualFold("2", types) {
		balance = balance - amount
		changeBalance = -amount
		logger.Info("减少余额", appUserId, changeBalance)
	} else {
		logger.Info("UpdateUserBalance")
	}
	var appUser models.AppUser
	DB.Model(models.AppUser{}).Where("app_user_id", appUserId).Find(&appUser)
	var appUserBalanceChangeLog models.AppUserBalanceChangeLog
	appUserBalanceChangeLog.AppUserId = appUser.AppUserId
	appUserBalanceChangeLog.ChangeBalance = changeBalance
	appUserBalanceChangeLog.ChangeType = changeType
	appUserBalanceChangeLog.CurrencyType = appUser.CurrencyType
	appUserBalanceChangeLog.ItemCode = appUser.ItemCode
	appUserBalanceChangeLog.Nation = appUser.Nation
	appUserBalanceChangeLog.SourceIndex = sourceIndex
	appUserBalanceChangeLog.TimerBalance = userBalance.Balance
	appUserBalanceChangeLog.CreateDate = time.Now()
	appUserBalanceChangeLog.UpdateDate = time.Now()
	err = DB.Model(models.AppUserBalanceChangeLog{}).Create(&appUserBalanceChangeLog).Error
	if err != nil {
		logger.Error(err)
		DB.Rollback()
		return false
	}
	err = DB.Model(models.AppUser{}).Where("app_user_id", appUser.AppUserId).Updates(map[string]interface{}{
		"update_date": time.Now(),
		"update_ts":   time.Now().UnixMilli(),
		"balance":     balance,
	}).Error
	if err != nil {
		logger.Error(err)
		DB.Rollback()
		return false
	}
	DB.Commit()
	return true
}

func UpdateUserBonusBalance(sourceIndex, appUserId, types string, amount float64, changeType int) bool {
	DB := mysql.DB.Debug().Begin()
	var userBalance UserBalance
	err := DB.Raw(`SELECT balance,bonus_balance FROM t_app_user WHERE app_user_id = ? FOR UPDATE`, appUserId).Scan(&userBalance).Error
	if err != nil {
		logger.Error(err)
		DB.Rollback()
		return false
	}
	bonusBalance := userBalance.BonusBalance
	var changeBalance float64 = 0
	if strings.EqualFold("1", types) {
		bonusBalance = bonusBalance + amount
		changeBalance = amount
		logger.Info("增加bonus", appUserId, changeBalance)
	} else if strings.EqualFold("2", types) {
		bonusBalance = bonusBalance - amount
		changeBalance = -amount
		logger.Info("减少bonus", appUserId, changeBalance)
	} else {
		logger.Info("UpdateUserBonusBalance")
	}
	var appUser models.AppUser
	err = DB.Model(models.AppUser{}).Where("app_user_id", appUserId).Find(&appUser).Error
	if err != nil {
		logger.Error(err)
		DB.Rollback()
		return false
	}
	var appUserBalanceChangeLog models.AppUserBalanceChangeLog
	appUserBalanceChangeLog.AppUserId = appUser.AppUserId
	appUserBalanceChangeLog.ChangeBalance = changeBalance
	appUserBalanceChangeLog.ChangeType = changeType
	appUserBalanceChangeLog.CurrencyType = appUser.CurrencyType
	appUserBalanceChangeLog.ItemCode = appUser.ItemCode
	appUserBalanceChangeLog.Nation = appUser.Nation
	appUserBalanceChangeLog.SourceIndex = sourceIndex
	appUserBalanceChangeLog.TimerBalance = userBalance.BonusBalance
	appUserBalanceChangeLog.CreateDate = time.Now()
	appUserBalanceChangeLog.UpdateDate = time.Now()
	err = DB.Model(models.AppUserBalanceChangeLog{}).Create(&appUserBalanceChangeLog).Error
	if err != nil {
		logger.Error(err)
		DB.Rollback()
		return false
	}
	err = DB.Model(models.AppUser{}).Where("app_user_id", appUser.AppUserId).Updates(map[string]interface{}{
		"update_date":   time.Now(),
		"update_ts":     time.Now().UnixMilli(),
		"bonus_balance": bonusBalance,
	}).Error
	logger.Info(bonusBalance, userBalance)
	if err != nil {
		logger.Error(err)
		DB.Rollback()
		return false
	}
	DB.Commit()
	return true
}
