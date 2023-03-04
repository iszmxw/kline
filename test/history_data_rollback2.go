package main

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/bootstrap"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"strings"
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
	// 初始化 SQL 回滚
	bootstrap.SetupDB()
}

type Result struct {
	Id                       int64
	Everyday                 string
	oldUserRecharged         int64
	oldUserRechargeTimes     int64
	oldUserLoss              float64
	oldUserRechargedRate     float64
	oldUserRechargeTimesRate float64
	oldUserLossRate          float64
}

func main() {
	bootstrap.SetupRedis(conf.GetInt("redis.db"))
	defer bootstrap.RedisClose()

	sql := `
SELECT
	IFNULL( '?', 0 ) AS Everyday,
	IFNULL( t3.oldUserRecharged, 0 ) AS oldUserRecharged,
	IFNULL( t3.newRegisteredRechargTimes, 0 ) AS newRegisteredRechargTimes,
	IFNULL( t3.oldUserRechargeTimes, 0 ) AS oldUserRechargeTimes,
	IFNULL( t8.oldUserLoss, 0 ) AS oldUserLoss,
	IFNULL( t3.oldUserRecharged / t3.rechargeUserNumber, 0 ) AS oldUserRechargedRate,
	IFNULL( t3.oldUserRechargeTimes / t3.tradeNumber, 0 ) AS oldUserRechargeTimesRate,
    IFNULL( t9.newUserLoss ,0 ) AS newUserLoss,
	IFNULL( t8.oldUserLoss / t6.userLoss, 0 ) AS oldUserLossRate 
FROM
	(
		(
		SELECT
			IFNULL( COUNT( DISTINCT tau.app_user_id ), 0 ) AS createNumber,
			IFNULL( COUNT( DISTINCT tau.decive_id ), 0 ) AS registUserNumber,
			DATE_FORMAT( tau.create_date, '%Y-%m-%d' ) AS userCreateDate 
		FROM
			t_app_user tau 
		WHERE
			DATE_FORMAT( tau.create_date, '%Y-%m-%d' ) > date_sub( DATE_FORMAT( '?', '%Y-%m-%d' ), INTERVAL 1 DAY ) 
			AND DATE_FORMAT( tau.create_date, '%Y-%m-%d' ) <= DATE_FORMAT( '?', '%Y-%m-%d' ) 
		) t2
		LEFT JOIN (
		SELECT
			IFNULL( SUM( tro.amount ), 0 ) AS rechargeAmount,
			IFNULL( COUNT( tro.recharge_id ), 0 ) AS tradeNumber,
			IFNULL( COUNT( DISTINCT tro.app_user_id ), 0 ) AS rechargeUserNumber,
			IFNULL( SUM( tro.fee ), 0 ) AS rechargeFee,
			IFNULL( SUM( tro.regist_date_recharge ), 0 ) AS registRecharge,
			IFNULL( COUNT( DISTINCT tro.app_user_id, IF ( tro.is_first_rechage = 1, 1, NULL )), 0 ) AS newRechargeNumber,
			IFNULL( COUNT( DISTINCT tro.app_user_id, IF ( tro.is_first_rechage = 0, 1, NULL )), 0 ) AS oldUserRecharged,
			IFNULL( COUNT( IF ( tro.is_first_rechage = 1, 1, NULL )), 0 ) AS newRegisteredRechargTimes,
			IFNULL( COUNT( IF ( tro.is_first_rechage = 0, 1, NULL )), 0 ) AS oldUserRechargeTimes,
			IFNULL( SUM( IF ( tro.is_first_rechage = 1, tro.amount, 0 )), 0 ) AS newRechargeAmount,
			DATE_FORMAT( tro.create_date, '%Y-%m-%d' ) AS rechargeCreateDate 
		FROM
			t_recharge_order tro 
		WHERE
			DATE_FORMAT( tro.create_date, '%Y-%m-%d' ) > date_sub( DATE_FORMAT( '?', '%Y-%m-%d' ), INTERVAL 1 DAY ) 
			AND DATE_FORMAT( tro.create_date, '%Y-%m-%d' ) <= DATE_FORMAT( '?', '%Y-%m-%d' ) 
			AND tro.recharge_order_status = 1 
		) t3 ON t3.rechargeCreateDate = DATE_FORMAT( '?', '%Y-%m-%d' )
		LEFT JOIN (
		SELECT
			IFNULL( SUM( two.amount ), 0 ) AS withdrawAmounts,
			IFNULL( COUNT( DISTINCT two.app_user_id ), 0 ) AS withdrawUserNumber,
			IFNULL( COUNT( two.withdraw_id ), 0 ) AS withdrawNumber,
			IFNULL( SUM( two.fee ), 0 ) AS withdrawFee,
			IFNULL( SUM( two.service_fee ), 0 ) AS withdrawServiceFee,
			DATE_FORMAT( two.create_date, '%Y-%m-%d' ) AS withdrawCreateDate 
		FROM
			t_withdraw_order two 
		WHERE
			DATE_FORMAT( two.create_date, '%Y-%m-%d' ) > date_sub( DATE_FORMAT( '?', '%Y-%m-%d' ), INTERVAL 1 DAY ) 
			AND DATE_FORMAT( two.create_date, '%Y-%m-%d' ) <= DATE_FORMAT( '?', '%Y-%m-%d' ) 
			AND two.withdraw_order_status = 1 
		) t4 ON t4.withdrawCreateDate = DATE_FORMAT( '?', '%Y-%m-%d' )
		LEFT JOIN (
		SELECT
			IFNULL( COUNT( DISTINCT IF ( tto.amount_type = 1, tto.app_user_id, NULL )), 0 ) AS tradingUserAmounts,
			IFNULL( COUNT( DISTINCT IF ( tto.amount_type = 1, tto.trade_order_id, NULL )), 0 ) AS tradeNumber,
			IFNULL( SUM( IF ( tto.amount_type = 1, tto.amounts, NULL )), 0 ) AS tradeAmounts,
			IFNULL( SUM( IF ( tto.amount_type = 1, tto.real_cost_amounts, NULL )), 0 ) AS realCostAmounts,
			IFNULL( SUM( IF ( tto.amount_type = 1, tto.real_gain_amounts, NULL )), 0 ) AS realGainAmounts,
			IFNULL( COUNT( DISTINCT IF ( tto.amount_type = 2, tto.app_user_id, NULL )), 0 ) AS bonusTradeNumber,
			IFNULL( SUM( IF ( tto.amount_type = 2, tto.real_cost_amounts, NULL )), 0 ) AS bonusRealCostAmounts,
			IFNULL( SUM( IF ( tto.amount_type = 2, tto.real_gain_amounts, NULL )), 0 ) AS bonusRealGainAmounts,
			IFNULL( SUM( CASE WHEN tto.quiz_result = 1 THEN 1 ELSE 0 END ), 0 ) AS userWinOrder,
			IFNULL( SUM( CASE WHEN tto.quiz_result = 0 THEN 1 ELSE 0 END ), 0 ) AS userLossOrder,
			IFNULL( SUM( IF ( tto.amount_type = 2, tto.amounts, 0 )), 0 ) AS bonusTradeAmounts,
			IFNULL( SUM( IF ( tto.amount_type = 1, tto.user_gain_amounts, 0 )), 0 ) AS userLoss,
			DATE_FORMAT( tto.create_date, '%Y-%m-%d' ) AS tradingCreateDate 
		FROM
			t_trading_order tto 
		WHERE
			DATE_FORMAT( tto.create_date, '%Y-%m-%d' ) > date_sub( DATE_FORMAT( '?', '%Y-%m-%d' ), INTERVAL 1 DAY ) 
			AND DATE_FORMAT( tto.create_date, '%Y-%m-%d' ) <= DATE_FORMAT( '?', '%Y-%m-%d' ) 
		) t6 ON t6.tradingCreateDate = DATE_FORMAT( '?', '%Y-%m-%d' )
		LEFT JOIN (
		SELECT
			IFNULL( SUM( tabl.amounts ), 0 ) AS bonusAmounts,
			IFNULL( COUNT( DISTINCT tabl.app_user_id ), 0 ) AS bonusTradingUserAmounts,
			DATE_FORMAT( tabl.create_date, '%Y-%m-%d' ) AS bonusCreateDate 
		FROM
			t_app_bonus_log tabl 
		WHERE
			DATE_FORMAT( tabl.create_date, '%Y-%m-%d' ) > date_sub( DATE_FORMAT( '?', '%Y-%m-%d' ), INTERVAL 1 DAY ) 
			AND DATE_FORMAT( tabl.create_date, '%Y-%m-%d' ) <= DATE_FORMAT( '?', '%Y-%m-%d' ) 
		) t5 ON t5.bonusCreateDate = DATE_FORMAT( '?', '%Y-%m-%d' )
		INNER JOIN (
		SELECT
			COUNT( tauacd.id ) AS numberInstallers 
		FROM
			t_app_user_adjust_callback_data tauacd 
		WHERE
			tauacd.event_name = 'install' 
			AND DATE_FORMAT( tauacd.create_date, '%Y-%m-%d' ) = DATE_FORMAT( '?', '%Y-%m-%d' ) 
		) t7
		LEFT JOIN (
		SELECT
			IFNULL( SUM( IF ( tto.amount_type = 1, tto.user_gain_amounts, 0 )), 0 ) AS oldUserLoss,
			DATE_FORMAT( tto.create_date, '%Y-%m-%d' ) AS tradingCreateDate 
		FROM
			t_trading_order tto
			LEFT JOIN t_app_user tau ON tto.app_user_id = tau.app_user_id 
		WHERE
			tau.is_new = 0 
			AND DATE_FORMAT( tto.create_date, '%Y-%m-%d' ) > date_sub( DATE_FORMAT( '?', '%Y-%m-%d' ), INTERVAL 1 DAY ) 
			AND DATE_FORMAT( tto.create_date, '%Y-%m-%d' ) <= DATE_FORMAT( '?', '%Y-%m-%d' ) 
		) t8 ON t8.tradingCreateDate = DATE_FORMAT( '?', '%Y-%m-%d' )
		LEFT JOIN (
        SELECT
			IFNULL( SUM( IF ( tto.amount_type = 1, tto.user_gain_amounts, 0 )), 0 ) AS newUserLoss,
			DATE_FORMAT( tto.create_date, '%Y-%m-%d' ) AS tradingCreateDate 
		FROM
			t_trading_order tto
			LEFT JOIN t_app_user tau ON tto.app_user_id = tau.app_user_id 
		WHERE
			tau.is_new = 1 
			AND DATE_FORMAT( tto.create_date, '%Y-%m-%d' ) > date_sub( DATE_FORMAT( '?', '%Y-%m-%d' ), INTERVAL 1 DAY ) 
			AND DATE_FORMAT( tto.create_date, '%Y-%m-%d' ) <= DATE_FORMAT( '?', '%Y-%m-%d' ) 
		) t9 ON t9.tradingCreateDate = DATE_FORMAT( '?', '%Y-%m-%d' ) 

	)
`

	// 查询历史所有数据
	historySql := `SELECT
	tdad.id AS Id,-- id
	DATE_FORMAT( tdad.create_date, '%Y-%m-%d' ) AS Everyday,-- 获取提币手续费
	IFNULL( tdad.withdrawal_commission, 0 ) AS WithdrawalCommission,-- 获取用户亏损
	IFNULL( tdad.user_loss, 0 ) AS UserLoss,-- 获取用户亏损
	IFNULL( tdad.platform_profit_loss, 0 ) AS PlatformProfitLoss 
FROM
	t_daily_active_data tdad`
	var ListHistory []Result
	mysql.DB.Raw(historySql).Scan(&ListHistory)
	for _, val := range ListHistory { // 遍历所有重新计算查询并且更新
		var result Result
		data := cmap.New().Items()
		timeStr := Parse2(val.Everyday).Format("2006-01-02")
		logger.Info(fmt.Sprint(val.Everyday, "   ", timeStr))
		mysql.DB.Debug().Raw(strings.ReplaceAll(sql, "?", timeStr)).Scan(&data)
		logger.Info(fmt.Sprint(val.Id, "  ", result, "  ", "ok"))
		logger.Info(fmt.Sprint(val.Id, "  ", data))
		Update(val.Id, data)
	}
}

func Update(id int64, result map[string]interface{}) {
	err := mysql.DB.Debug().Model(models.DailyActiveData{}).Where("id", id).
		Updates(map[string]interface{}{
			"new_registered_recharg_times": result["newRegisteredRechargTimes"],
			"old_user_recharged":           result["oldUserRecharged"],
			"old_user_recharged_rate":      result["oldUserRechargedRate"],
			"old_user_recharge_times":      result["oldUserRechargeTimes"],
			"old_user_recharge_times_rate": result["oldUserRechargeTimesRate"],
			"new_user_loss":                result["newUserLoss"],
			"old_user_loss":                result["oldUserLoss"],
			"old_user_loss_rate":           result["oldUserLossRate"],
		}).
		Error
	if err != nil {
		logger.Error(err)
		panic(err)
	}
}

func Parse2(strTime string) time.Time {
	var result time.Time
	// 东七
	result, err := time.ParseInLocation("2006-01-02", strTime, time.FixedZone("CST", 7*3600)) // 2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
	if err != nil {
		logger.Error(err)
	}
	return result
}
