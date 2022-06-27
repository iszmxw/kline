package main

import (
	"fmt"
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("test")
	// 初始化 SQL
	bootstrap.SetupDB()
}

func main() {
	bootstrap.SetupRedis(0)
	defer bootstrap.RedisClose()

//	sql := `
//SELECT
//	IFNULL( '?', 0 ) as Everyday,
//	-- 获取提币手续费
//	IFNULL( t4.withdrawFee, 0 ) AS WithdrawalCommission,
//	-- 获取用户亏损
//	IFNULL( t6.userLoss, 0 ) AS UserLoss,
//	-- 设置平台盈亏
//  (IFNULL( t4.withdrawFee, 0 ) - IFNULL( t6.userLoss, 0 )) AS PlatformProfitLoss
//FROM
//	(
//		SELECT
//			IFNULL( COUNT( DISTINCT tau.app_user_id ), 0 ) AS createNumber,
//			IFNULL( COUNT( DISTINCT tau.decive_id ), 0 ) AS registUserNumber,
//			DATE_FORMAT( tau.create_date, '%Y-%m-%d' ) AS userCreateDate
//		FROM
//			t_app_user tau
//		WHERE
//			tau.create_date >= date_sub( '?', INTERVAL 1 DAY )
//			AND tau.create_date < '?'
//		) t2
//		LEFT JOIN (
//		SELECT
//			IFNULL( SUM( tro.amount ), 0 ) AS rechargeAmount,
//			IFNULL( COUNT( tro.recharge_id ), 0 ) AS tradeNumber,
//			IFNULL( COUNT( DISTINCT tro.app_user_id ), 0 ) AS rechargeUserNumber,
//			IFNULL( SUM( tro.fee ), 0 ) AS rechargeFee,
//			IFNULL( SUM( tro.regist_date_recharge ), 0 ) AS registRecharge,
//			IFNULL( COUNT( DISTINCT tro.app_user_id, IF ( tro.is_first_rechage = 1, 1, NULL )), 0 ) AS newRechargeNumber,
//			IFNULL( SUM( IF ( tro.is_first_rechage = 1, tro.amount, 0 )), 0 ) AS newRechargeAmount,
//			DATE_FORMAT( tro.create_date, '%Y-%m-%d' ) AS rechargeCreateDate
//		FROM
//			t_recharge_order tro
//		WHERE
//			tro.create_date >= date_sub( '?', INTERVAL 1 DAY )
//			AND tro.create_date < '?'
//			AND tro.recharge_order_status = 1
//		) t3 ON t3.rechargeCreateDate = date_sub( '?', INTERVAL 1 DAY )
//		LEFT JOIN (
//		SELECT
//			IFNULL( SUM( two.amount ), 0 ) AS withdrawAmounts,
//			IFNULL( COUNT( DISTINCT two.app_user_id ), 0 ) AS withdrawUserNumber,
//			IFNULL( COUNT( two.withdraw_id ), 0 ) AS withdrawNumber,
//			IFNULL( SUM( two.fee ), 0 ) AS withdrawFee,
//			IFNULL( SUM( two.service_fee ), 0 ) AS withdrawServiceFee,
//			DATE_FORMAT( two.create_date, '%Y-%m-%d' ) AS withdrawCreateDate
//		FROM
//			t_withdraw_order two
//		WHERE
//			two.create_date >= date_sub( '?', INTERVAL 1 DAY )
//			AND two.create_date < '?'
//			AND two.withdraw_order_status = 1
//		) t4 ON t4.withdrawCreateDate = date_sub( '?', INTERVAL 1 DAY )
//		LEFT JOIN (
//		SELECT
//			IFNULL( COUNT( DISTINCT IF ( tto.amount_type = 1, tto.app_user_id, NULL )), 0 ) AS tradingUserAmounts,
//			IFNULL( COUNT( DISTINCT IF ( tto.amount_type = 1, tto.trade_order_id, NULL )), 0 ) AS tradeNumber,
//			IFNULL( SUM( IF ( tto.amount_type = 1, tto.amounts, NULL )), 0 ) AS tradeAmounts,
//			IFNULL( SUM( IF ( tto.amount_type = 1, tto.real_cost_amounts, NULL )), 0 ) AS realCostAmounts,
//			IFNULL( SUM( IF ( tto.amount_type = 1, tto.real_gain_amounts, NULL )), 0 ) AS realGainAmounts,
//			IFNULL( COUNT( IF ( tto.amount_type = 2, tto.app_user_id, NULL )), 0 ) AS bonusTradeNumber,
//			IFNULL( SUM( IF ( tto.amount_type = 2, tto.real_cost_amounts, NULL )), 0 ) AS bonusRealCostAmounts,
//			IFNULL( SUM( IF ( tto.amount_type = 2, tto.real_gain_amounts, NULL )), 0 ) AS bonusRealGainAmounts,
//			IFNULL( SUM( CASE WHEN tto.quiz_result = 1 THEN 1 ELSE 0 END ), 0 ) AS userWinOrder,
//			IFNULL( SUM( CASE WHEN tto.quiz_result = 0 THEN 1 ELSE 0 END ), 0 ) AS userLossOrder,
//			IFNULL( SUM( IF ( tto.amount_type = 2, tto.amounts, 0 )), 0 ) AS bonusTradeAmounts,
//			IFNULL( SUM( IF ( tto.amount_type = 1, tto.user_gain_amounts, 0 )), 0 ) AS userLoss,
//			DATE_FORMAT( tto.create_date, '%Y-%m-%d' ) AS tradingCreateDate
//		FROM
//			t_trading_order tto
//		WHERE
//			tto.create_date >= date_sub( '?', INTERVAL 1 DAY )
//			AND tto.create_date < '?'
//		) t6 ON t6.tradingCreateDate = date_sub( '?', INTERVAL 1 DAY )
//		LEFT JOIN (
//		SELECT
//			IFNULL( SUM( tabl.amounts ), 0 ) AS bonusAmounts,
//			IFNULL( COUNT( DISTINCT tabl.app_user_id ), 0 ) AS bonusTradingUserAmounts,
//			DATE_FORMAT( tabl.create_date, '%Y-%m-%d' ) AS bonusCreateDate
//		FROM
//			t_app_bonus_log tabl
//		WHERE
//			tabl.create_date >= date_sub( '?', INTERVAL 1 DAY )
//			AND tabl.create_date < '?'
//		) t5 ON t5.bonusCreateDate = date_sub( '?', INTERVAL 1 DAY )
//		INNER JOIN (
//		SELECT
//			COUNT( tauacd.id ) AS numberInstallers
//		FROM
//			t_app_user_adjust_callback_data tauacd
//		WHERE
//			tauacd.event_name = 'install'
//			AND tauacd.create_date >= date_sub( '?', INTERVAL 1 DAY )
//			AND tauacd.create_date < '?'
//		) t7
//`
	type Result struct {
		Everyday             string
		WithdrawalCommission float64
		UserLoss             float64
		PlatformProfitLoss   float64
	}
	// 查询历史所有数据
	historySql := `SELECT
	DATE_FORMAT( tdad.create_date, '%Y-%m-%d' ) AS Everyday,-- 获取提币手续费
	IFNULL( tdad.withdrawal_commission, 0 ) AS WithdrawalCommission,-- 获取用户亏损
	IFNULL( tdad.user_loss, 0 ) AS UserLoss,-- 获取用户亏损
	IFNULL( tdad.platform_profit_loss, 0 ) AS PlatformProfitLoss 
FROM
	t_daily_active_data tdad`
	var ListHistory []Result
	mysql.DB.Raw(historySql).Scan(&ListHistory)
	for _, val := range ListHistory { // 遍历所有重新计算查询并且更新
		//var result Result
		timeStr := Parse(val.Everyday).Add(-1).Format("2006-01-02")
		logger.Info(timeStr)
		logger.Info(fmt.Sprint(val.Everyday, timeStr))
		//mysql.DB.Raw(strings.ReplaceAll(sql, "?", timeStr)).Scan(&result)
		//logger.Info(fmt.Sprint(result, val))
	}
}

func Parse(strTime string) time.Time {
	var result time.Time
	// 东七
	result, err := time.ParseInLocation("2006-01-02 15:04:05", strTime, time.FixedZone("CST", 7*3600)) // 2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
	if err != nil {
		logger.Error(err)
	}
	return result
}
