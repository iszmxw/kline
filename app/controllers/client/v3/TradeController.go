package v3

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"goapi/app/models"
	"goapi/app/requests"
	"goapi/app/response"
	goadjust "goapi/pkg/adjust"
	"goapi/pkg/config"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	wss "goapi/pkg/websocket"
	"strings"
	"time"
)

// 交易模块

type TradeController struct {
	BaseController
}

// 提交下单

func (h *TradeController) SubmitTrade(socketId string, messageType int, ws *wss.WsConn, subInfo string) {
	logger.RequestId = socketId
	// 下单参数
	var SubmitTradeParams requests.SubmitTradeParams
	err1 := json.Unmarshal([]byte(subInfo), &SubmitTradeParams)
	if err1 != nil {
		logger.Error(err1)
		socketResp(ws, messageType, `{"id":"10409","data":"ParamsCanNotBeNull"}`)
		return
	}
	if SubmitTradeParams.Amounts <= 0 {
		socketResp(ws, messageType, `{"id":"10409","data":"ParamsCanNotBeNull"}`)
		return
	}
	if SubmitTradeParams.AmountType <= 0 {
		socketResp(ws, messageType, `{"id":"10409","data":"ParamsCanNotBeNull"}`)
		return
	}
	if len(SubmitTradeParams.AssetsId) == 0 {
		socketResp(ws, messageType, `{"id":"10409","data":"ParamsCanNotBeNull"}`)
		return
	}
	if SubmitTradeParams.TradingBeginDate.UnixMilli() <= 0 {
		socketResp(ws, messageType, `{"id":"10409","data":"ParamsCanNotBeNull"}`)
		return
	}
	if SubmitTradeParams.TradeEndDate.UnixMilli() <= 0 {
		socketResp(ws, messageType, `{"id":"10409","data":"ParamsCanNotBeNull"}`)
		return
	}
	if len(SubmitTradeParams.UpDownType) == 0 {
		socketResp(ws, messageType, `{"id":"10409","data":"ParamsCanNotBeNull"}`)
		return
	}
	if SubmitTradeParams.BeginAssetsPrice <= 0 {
		socketResp(ws, messageType, `{"id":"10409","data":"ParamsCanNotBeNull"}`)
		return
	}
	if SubmitTradeParams.KlineTime <= 0 {
		socketResp(ws, messageType, `{"id":"10409","data":"ParamsCanNotBeNull"}`)
		return
	}
	sysTime := time.Now()
	// 请求超时拦截
	if SubmitTradeParams.TradingBeginDate.UnixMilli() < sysTime.UnixMilli()-10000 && 3 != SubmitTradeParams.AmountType {
		socketResp(ws, messageType, `{"id":"10409","data":"RequestTimeout"}`)
		return
	}
	// 时间超前拦截
	if SubmitTradeParams.TradingBeginDate.UnixMilli() > sysTime.UnixMilli() && 3 != SubmitTradeParams.AmountType {
		socketResp(ws, messageType, `{"id":"10409","data":"PleaseCheckTheNetworkTime"}`)
		return
	}
	// demo 订单只能同时存在五笔
	var total int64
	mysql.DB.Debug().Model(models.TradingOrderVirtual{}).Where("quiz_result", -1).Where("app_user_id", SubmitTradeParams.TradeOrderId).Count(&total)
	if total >= 5 && 3 == SubmitTradeParams.AmountType {
		socketResp(ws, messageType, `{"id":"10409","data":"DemoLimit"}`)
		return
	}
	// 校验是否存在该交易点数据
	var size int64
	err := mysql.DB.Debug().
		Raw(`SELECT COUNT(tkl.c_close) FROM t_kline_log tkl,t_trading_assets tta WHERE tkl.contract_code = tta.assets_code AND tkl.ts = ? AND tta.id = ? AND tkl.c_close = ?`,
			SubmitTradeParams.KlineTime, SubmitTradeParams.AssetsId, SubmitTradeParams.BeginAssetsPrice).
		Row().
		Scan(&size).Error
	if err != nil {
		socketResp(ws, messageType, `{"id":"10409","data":"校验是否存在该交易点数据"}`)
		return
	}
	if size == 0 {
		socketResp(ws, messageType, `{"id":"10409","data":"校验是否存在该交易点数据"}`)
		return
	}
	// 校验资产 资产已经休市，请更换其他资产
	var asset models.TradingAssets
	mysql.DB.Debug().Model(models.TradingAssets{}).Where("id", SubmitTradeParams.AssetsId).Find(&asset)
	if asset.IsValid != 1 {
		socketResp(ws, messageType, `{"id":"10409","data":"AssetClosed"}`)
		return
	}
	var appUser models.AppUser
	mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", SubmitTradeParams.AssetsId).Find(&appUser)
	// 进行交易规则校验，虚拟交易不进行校验
	if SubmitTradeParams.AmountType != 3 {
		ok, msg := tradeJuedge(appUser, SubmitTradeParams)
		if ok {
			socketResp(ws, messageType, fmt.Sprintf(`{"id":"10409","data":"%v"}`, msg))
			return
		}
	}
	// 交易金额类型 1 余额 2 bonus 3 虚拟货币
	switch SubmitTradeParams.AmountType {
	case 1:
		if appUser.Balance < SubmitTradeParams.Amounts {
			socketResp(ws, messageType, fmt.Sprintf(`{"id":"10409","data":"%v"}`, "余额不足"))
			return
		}
		ok := createTradingOrder(SubmitTradeParams, appUser, asset)
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
			socketResp(ws, messageType, fmt.Sprintf(`{"id":"1000","data":"%v"}`, "ok"))
		}
		break
	case 2:
		if appUser.BonusBalance < SubmitTradeParams.Amounts {
			socketResp(ws, messageType, fmt.Sprintf(`{"id":"10409","data":"%v"}`, "bonus余额不足"))
			return
		}
		ok := createTradingOrder(SubmitTradeParams, appUser, asset)
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
			socketResp(ws, messageType, fmt.Sprintf(`{"id":"1000","data":"%v"}`, "ok"))
		}
		break
	case 3:
		if appUser.VirtualBalance < SubmitTradeParams.Amounts {
			socketResp(ws, messageType, fmt.Sprintf(`{"id":"10409","data":"%v"}`, "虚拟余额不足"))
			return
		}
		createVirtualTradingOrder(SubmitTradeParams, appUser, asset)
		socketResp(ws, messageType, fmt.Sprintf(`{"id":"1000","data":"%v"}`, "ok"))
		break
	default:
		break
	}
}

// 交易规则判断

func tradeJuedge(appUser models.AppUser, params requests.SubmitTradeParams) (bool, string) {
	// 获取用户所在国家的时区
	var appItem models.AppItem
	err := mysql.DB.Debug().Model(models.AppItem{}).Where("item_code", appUser.ItemCode).Where("item_national_name", appUser.Nation).Find(&appItem).Error
	if err != nil {
		logger.Error(err)
		return true, "交易规则未配置"
	}
	timeZone := strings.SplitAfter(appItem.TimeZone, "T")
	var userTradeMessage response.UserTradeMessage
	err = mysql.DB.Debug().Raw(`SELECT a.app_user_id AS 'appUserId' ,
			       a.is_new AS 'isNew' ,
			       a.black_list AS 'blackList' ,
			       ifnull(b.tradingCount,0) as 'tradingCount',
			       ifnull(b.tradingAmounts,0) as 'tradingAmounts',
			       ifnull(c.rechargeAmounts,0) as 'rechargeAmounts',
			       ifnull(c.rechargeCount,0) as 'rechargeCount',
			       ifnull(c.rechargeTodayAmounts,0) as 'rechargeTodayAmounts',
			       ifnull(c.rechargeTodayCount,0) as 'rechargeTodayCount',
			       ifnull(d.withdrawAmounts,0) as 'withdrawAmounts',
			       ifnull(d.withdrawCount,0) as 'withdrawCount',
			       ifnull(d.withdrawTodayAmounts,0) as 'withdrawTodayAmounts',
			       ifnull(d.withdrawTodayCount,0) as 'withdrawTodayCount',
			       ifnull(e.hisTradingCount,0) as 'hisTradingCount',
			       ifnull(e.hisTradingAmounts,0) as 'hisTradingAmounts'
			FROM (
				SELECT a.app_user_id,
				       '1' AS 'biao',
				       a.is_new,
				       a.black_list
				FROM t_app_user AS a
				WHERE 1=1
				AND a.app_user_id = ?
				AND a.nation = ?
			) AS a INNER JOIN
			(
				SELECT COUNT(1) AS 'tradingCount',
				       SUM(a.amounts) AS 'tradingAmounts',
				       '1' AS 'biao'
				FROM t_trading_order AS a
				WHERE 1=1
				AND a.app_user_id = ?
				AND DATE_FORMAT(CONVERT_TZ(a.create_date, '+07:00', ?),'%Y-%m-%d') = DATE_FORMAT(CONVERT_TZ(NOW(), '+07:00', ?),'%Y-%m-%d')
				AND a.amount_type = 1
				AND a.nation = ?
			) AS b ON a.biao = b.biao INNER JOIN
			(
				SELECT SUM(b.amount) AS 'rechargeAmounts',
				       COUNT(1) AS 'rechargeCount',
				       SUM(CASE WHEN DATE_FORMAT(CONVERT_TZ(b.create_date, '+07:00', ?),'%Y-%m-%d') = DATE_FORMAT(CONVERT_TZ(NOW(), '+07:00', ?),'%Y-%m-%d') THEN b.amount ELSE 0 END ) AS 'rechargeTodayAmounts',
				       SUM(CASE WHEN DATE_FORMAT(CONVERT_TZ(b.create_date, '+07:00', ?),'%Y-%m-%d') = DATE_FORMAT(CONVERT_TZ(NOW(), '+07:00', ?),'%Y-%m-%d') THEN 1 ELSE 0 END ) AS 'rechargeTodayCount',
				       '1' AS 'biao'
				FROM t_recharge_order AS b
				WHERE 1=1
				AND b.app_user_id = ?
				AND b.recharge_order_status = 1
				AND b.nation = ?
			) AS c ON a.biao = c.biao  INNER JOIN
			(
				SELECT SUM(b.amount) AS 'withdrawAmounts',
				       COUNT(1) AS 'withdrawCount',
				       SUM(CASE WHEN DATE_FORMAT(CONVERT_TZ(b.create_date, '+07:00', ?),'%Y-%m-%d') = DATE_FORMAT(CONVERT_TZ(NOW(), '+07:00', ?),'%Y-%m-%d') THEN b.amount ELSE 0 END ) AS 'withdrawTodayAmounts',
				       SUM(CASE WHEN DATE_FORMAT(CONVERT_TZ(b.create_date, '+07:00', ?),'%Y-%m-%d') = DATE_FORMAT(CONVERT_TZ(NOW(), '+07:00', ?),'%Y-%m-%d') THEN 1 ELSE 0 END ) AS 'withdrawTodayCount',
				       '1' AS 'biao'
				FROM t_withdraw_order AS b
				WHERE 1=1
				AND b.app_user_id = ?
				AND b.withdraw_order_status IN (1,2)
					AND b.nation = ?
			) AS d  ON a.biao = d.biao INNER JOIN
			(
				SELECT COUNT(1) AS 'hisTradingCount',
				       SUM(a.amounts) AS 'hisTradingAmounts',
				       '1' AS 'biao'
				FROM t_trading_order AS a
				WHERE 1=1
				AND a.app_user_id = ?
				AND a.amount_type = 1
				AND a.nation = ?
			) AS e ON a.biao = e.biao
	`, appUser.AppUserId, appUser.Nation, appUser.AppUserId, timeZone, timeZone, appUser.Nation, timeZone, timeZone, timeZone, timeZone, appUser.AppUserId, appUser.Nation, timeZone, timeZone, timeZone, timeZone, appUser.AppUserId, appUser.Nation, appUser.AppUserId, appUser.Nation).
		Row().
		Scan(&userTradeMessage)
	if err != nil {
		return true, ""
	}
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
				return true, strings.ReplaceAll(item.ValueTips, "[value]", fmt.Sprintf("[%v]", item.Value))
			}
		}
		// 最大交易金额限制，命中规则拦截
		if strings.EqualFold("maxTradingAmounts", item.RuleCode) {
			if DecisionAlgorithmInterface(int(params.Amounts), int(item.Value), item.CalculateMethod) {
				return true, strings.ReplaceAll(item.ValueTips, "[value]", fmt.Sprintf("[%v]", item.Value))
			}
		}
		// 最多交易次数限制，命中规则拦截
		if strings.EqualFold("maxTradingTimes", item.RuleCode) {
			if DecisionAlgorithmInterface(userTradeMessage.TradingCount, int(item.Value), item.CalculateMethod) {
				return true, strings.ReplaceAll(item.ValueTips, "[value]", fmt.Sprintf("[%v]", item.Value))
			}
		}
		// 日交易总金额限制，命中规则拦截【不判断影响不大】
		if strings.EqualFold("dayMaxTradingAmounts", item.RuleCode) {
			if DecisionAlgorithmInterface(userTradeMessage.TradingAmounts, int(item.Value), item.CalculateMethod) {
				return true, strings.ReplaceAll(item.ValueTips, "[value]", fmt.Sprintf("[%v]", item.Value))
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
func createTradingOrder(params requests.SubmitTradeParams, appUser models.AppUser, asset models.TradingAssets) bool {
	orderId := "G_" + strings.ReplaceAll(uuid.New().String(), "-", "")
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
	createOrder.TradeBeginDate = params.TradingBeginDate
	createOrder.TradeEndDate = params.TradeEndDate
	createOrder.KlineTime = params.KlineTime
	createOrder.TradeBeginTimestamp = params.TradingBeginDate.UnixMilli()
	createOrder.TradeEndTimestamp = params.TradeEndDate.UnixMilli()
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
		UpdateUserBalance(orderId, appUser.AppUserId, "2", params.Amounts, 2)
		break
	case 2:
		UpdateUserBonusBalance(orderId, appUser.AppUserId, "2", params.Amounts, 7)
		break
	case 3:
		err = mysql.DB.Model(models.AppUser{}).Where("app_user_id", appUser.AppUserId).Updates(map[string]interface{}{
			"virtual_balance": appUser.VirtualBalance - params.Amounts,
		}).Error
		if err != nil {
			logger.Error(err)
			return false
		}
		break
	default:
		break
	}
	return true
}

// 创建虚拟交易订单
func createVirtualTradingOrder(params requests.SubmitTradeParams, appUser models.AppUser, asset models.TradingAssets) bool {
	orderId := "G_" + strings.ReplaceAll(uuid.New().String(), "-", "")
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
	createOrder.TradeBeginDate = params.TradingBeginDate
	createOrder.TradeEndDate = params.TradeEndDate
	createOrder.TradeBeginTimestamp = params.TradingBeginDate.UnixMilli()
	createOrder.TradeEndTimestamp = params.TradeEndDate.UnixMilli()
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
	BonusBalance float64 `json:"bonusBalance"`
}

func UpdateUserBalance(sourceIndex, appUserId, types string, amount float64, changeType int) {
	DB := mysql.DB.Debug()
	DB.Begin()
	var userBalance UserBalance
	err := DB.Raw(`SELECT balance,bonus_balance as bonusBalance FROM t_app_user WHERE app_user_id = ? FOR UPDATE`, appUserId).Row().Scan(&userBalance)
	if err != nil {
		logger.Error(err)
		DB.Rollback()
		return
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
	DB.Model(models.AppUserBalanceChangeLog{}).Create(&appUserBalanceChangeLog)
	DB.Model(models.AppUser{}).Where("app_user_id", appUser.AppUserId).Updates(map[string]interface{}{
		"update_date": time.Now(),
		"update_ts":   time.Now().UnixMilli(),
		"balance":     balance,
	})
	DB.Commit()
}

func UpdateUserBonusBalance(sourceIndex, appUserId, types string, amount float64, changeType int) {
	DB := mysql.DB.Debug()
	DB.Begin()
	var userBalance UserBalance
	err := DB.Raw(`SELECT balance,bonus_balance as bonusBalance FROM t_app_user WHERE app_user_id = ? FOR UPDATE`, appUserId).Row().Scan(&userBalance)
	if err != nil {
		logger.Error(err)
		DB.Rollback()
		return
	}
	bonusBalance := userBalance.Balance
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
	DB.Model(models.AppUser{}).Where("app_user_id", appUserId).Find(&appUser)
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
	DB.Model(models.AppUserBalanceChangeLog{}).Create(&appUserBalanceChangeLog)
	DB.Model(models.AppUser{}).Where("app_user_id", appUser.AppUserId).Updates(map[string]interface{}{
		"update_date":   time.Now(),
		"update_ts":     time.Now().UnixMilli(),
		"bonus_balance": bonusBalance,
	})
	DB.Commit()
}
