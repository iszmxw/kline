package mapper

import (
	"fmt"
	"goapi/app/models"
	"goapi/app/requests"
	"goapi/app/response"
	"goapi/pkg/mysql"
	"strings"
	"time"
)

func GetUserTradeMessage(appItem response.ItemNationMessageParams, tAppUser models.AppUser) (userTradeMessage response.UserTradeMessage, err error) {
	sqlStr := `
SELECT a.app_user_id AS appUserId ,
        a.is_new AS isNew ,
        a.black_list AS blackList ,
        ifnull(b.tradingCount,0) as tradingCount,
        ifnull(b.tradingAmounts,0) as tradingAmounts,
        ifnull(c.rechargeAmounts,0) as rechargeAmounts,
        ifnull(c.rechargeCount,0) as rechargeCount,
        ifnull(c.rechargeTodayAmounts,0) as rechargeTodayAmounts,
        ifnull(c.rechargeTodayCount,0) as rechargeTodayCount,
        ifnull(d.withdrawAmounts,0) as withdrawAmounts,
        ifnull(d.withdrawCount,0) as withdrawCount,
        ifnull(d.withdrawTodayAmounts,0) as withdrawTodayAmounts,
        ifnull(d.withdrawTodayCount,0) as withdrawTodayCount,
        ifnull(e.hisTradingCount,0) as hisTradingCount,
        ifnull(e.hisTradingAmounts,0) as hisTradingAmounts
FROM (
         SELECT a.app_user_id,
                '1' AS biao,
                 a.is_new,
                a.black_list
         FROM t_app_user AS a
         WHERE 1=1
           AND a.app_user_id = #{appUserId}
           AND a.nation = #{nation}
     ) AS a INNER JOIN
     (
         SELECT COUNT(1) AS tradingCount,
                 SUM(a.amounts) AS tradingAmounts,
                 '1' AS biao
         FROM t_trading_order AS a
         WHERE 1=1
           AND a.app_user_id = #{appUserId}
			AND DATE_FORMAT(CONVERT_TZ(a.create_date, '+07:00', #{timeZone}),'%Y-%m-%d') = DATE_FORMAT(CONVERT_TZ(NOW(), '+07:00', #{timeZone}),'%Y-%m-%d')
			AND a.amount_type = 1
			AND a.nation = #{nation}
     ) AS b ON a.biao = b.biao INNER JOIN
     (
         SELECT SUM(b.amount) AS rechargeAmounts,
                 COUNT(1) AS rechargeCount,
                 SUM(CASE WHEN DATE_FORMAT(CONVERT_TZ(b.create_date, '+07:00', #{timeZone}),'%Y-%m-%d') = DATE_FORMAT(CONVERT_TZ(NOW(), '+07:00', #{timeZone}),'%Y-%m-%d') THEN b.amount ELSE 0 END ) AS rechargeTodayAmounts,
                 SUM(CASE WHEN DATE_FORMAT(CONVERT_TZ(b.create_date, '+07:00', #{timeZone}),'%Y-%m-%d') = DATE_FORMAT(CONVERT_TZ(NOW(), '+07:00', #{timeZone}),'%Y-%m-%d') THEN 1 ELSE 0 END ) AS rechargeTodayCount,
                 '1' AS biao
         FROM t_recharge_order AS b
         WHERE 1=1
           AND b.app_user_id = #{appUserId}
			AND b.recharge_order_status = 1
			AND b.nation = #{nation}
     ) AS c ON a.biao = c.biao  INNER JOIN
     (
         SELECT SUM(b.amount) AS withdrawAmounts,
                 COUNT(1) AS withdrawCount,
                 SUM(CASE WHEN DATE_FORMAT(CONVERT_TZ(b.create_date, '+07:00', #{timeZone}),'%Y-%m-%d') = DATE_FORMAT(CONVERT_TZ(NOW(), '+07:00', #{timeZone}),'%Y-%m-%d') THEN b.amount ELSE 0 END ) AS withdrawTodayAmounts,
                 SUM(CASE WHEN DATE_FORMAT(CONVERT_TZ(b.create_date, '+07:00', #{timeZone}),'%Y-%m-%d') = DATE_FORMAT(CONVERT_TZ(NOW(), '+07:00', #{timeZone}),'%Y-%m-%d') THEN 1 ELSE 0 END ) AS withdrawTodayCount,
                 '1' AS biao
         FROM t_withdraw_order AS b
         WHERE 1=1
           AND b.app_user_id = #{appUserId}
			AND b.withdraw_order_status IN (1,2)
            AND b.nation = #{nation}
     ) AS d  ON a.biao = d.biao INNER JOIN
     (
         SELECT COUNT(1) AS hisTradingCount,
                 SUM(a.amounts) AS hisTradingAmounts,
                 '1' AS biao
         FROM t_trading_order AS a
         WHERE 1=1
           AND a.app_user_id = #{appUserId}
			AND a.amount_type = 1
			AND a.nation = #{nation}
     ) AS e ON a.biao = e.biao
`
	timeZone := strings.SplitAfter(appItem.TimeZone, "T")
	sqlStr = strings.ReplaceAll(sqlStr, "#{appUserId}", tAppUser.AppUserId)
	sqlStr = strings.ReplaceAll(sqlStr, "#{nation}", tAppUser.Nation)
	sqlStr = strings.ReplaceAll(sqlStr, "#{timeZone}", timeZone[1])
	err = mysql.DB.Debug().Raw(sqlStr).Scan(&userTradeMessage).Error
	return userTradeMessage, err
}

func GetInvitTop1Info(params requests.MineInvitUserParams, beginTimestamp, endTimestamp time.Time) (invitTop1Info *response.UserInvitTopOneResult) {
	sqlStr2 := `
		SELECT
            COUNT(1) AS invitBonus,
            t.app_user_id AS appUserId,
            SUM(t.amounts) AS invitAmounts
        FROM
            t_app_bonus_log t
        LEFT JOIN t_app_user b ON b.app_user_id = t.app_user_id
        WHERE
            t.item_code = #{itemCode}
          AND t.activity_code = 'invite_politely'
          AND t.nation = #{nation}
          AND t.create_ts  BETWEEN #{beginTimestamp} AND #{endTimestamp}
          AND b.is_logout = 0
        GROUP BY t.app_user_id
        HAVING COUNT(1) >= 3
        ORDER BY invitBonus DESC
            LIMIT 1
`
	sqlStr2 = strings.ReplaceAll(sqlStr2, "#{itemCode}", params.ItemCode)
	sqlStr2 = strings.ReplaceAll(sqlStr2, "#{nation}", params.Nation)
	sqlStr2 = strings.ReplaceAll(sqlStr2, "#{beginTimestamp}", fmt.Sprintf("%v", beginTimestamp.UnixMilli()))
	sqlStr2 = strings.ReplaceAll(sqlStr2, "#{endTimestamp}", fmt.Sprintf("%v", endTimestamp.UnixMilli()))
	mysql.DB.Raw(sqlStr2).Scan(&invitTop1Info)
	return
}

func GetMineInvitUserList(itemCode, referralCode, index, ends, nation string) (invitUserResults []response.MineInvitUserResults) {
	sqlStr := `
		SELECT
	        ot.appUserId,
	        ot.isLogout,
	        ot.email,
	        ot.avatar,
            IF( ot.invitStatus = 1, '-1', IFNULL(ot.orderAmount,0) ) AS invitStatus
        FROM
            (
            SELECT
                t.app_user_id AS appUserId,
                t.email,
                IFNULL(t.new_avatar,t.avatar) as avatar,
                t.is_logout as isLogout,
                t.is_new AS invitStatus,(
                    SELECT
                        COUNT( 1 )
                    FROM
                        t_trading_order tto
                    WHERE
                        tto.app_user_id = t.app_user_id
                      and tto.nation = #{nation}
                      and tto.quiz_result != -1
                    GROUP BY
                        t.app_user_id
                ) AS orderAmount
            FROM
                t_app_user t
            WHERE
                t.parent_code = #{referralCode}
                AND t.item_code = #{itemCode}
                AND t.nation = #{nation}
            ORDER BY
                t.create_date DESC
                LIMIT #{index},#{ends}
            ) ot
`
	sqlStr = strings.ReplaceAll(sqlStr, "#{nation}", nation)
	sqlStr = strings.ReplaceAll(sqlStr, "#{referralCode}", referralCode)
	sqlStr = strings.ReplaceAll(sqlStr, "#{itemCode}", itemCode)
	sqlStr = strings.ReplaceAll(sqlStr, "#{index}", fmt.Sprintf("%v", index))
	sqlStr = strings.ReplaceAll(sqlStr, "#{ends}", fmt.Sprintf("%v", ends))
	mysql.DB.Raw(sqlStr).Scan(&invitUserResults)
	return
}

func GetUserMailMessageListCounts(itemCode, appUserId, deviceType, nation string, date int64) (total int64) {
	sqlStr := `
		SELECT
      count(tomsl.app_user_id)
      FROM t_notice_log tomsl
      LEFT JOIN t_notice_message toms ON toms.message_id = tomsl.message_id
      LEFT JOIN t_notice tom ON tom.mail_id = tomsl.mail_id
      WHERE
      tomsl.item_code = #{itemCode}
      AND tomsl.nation = #{nation}
      AND tomsl.decive_brand = #{deviceType}
      AND (tomsl.app_user_id = #{appUserId} OR tomsl.app_user_id = 'ALL')
      AND tom.send_state = 1
      AND tom.send_time <![CDATA[>=]]> #{date}
      AND (tom.send_time / 1000) >= UNIX_TIMESTAMP(DATE_SUB(CURDATE(),INTERVAL 7 DAY))
      ORDER BY tom.send_time DESC
`
	sqlStr = strings.ReplaceAll(sqlStr, "#{itemCode}", itemCode)
	sqlStr = strings.ReplaceAll(sqlStr, "#{appUserId}", appUserId)
	sqlStr = strings.ReplaceAll(sqlStr, "#{deviceType}", deviceType)
	sqlStr = strings.ReplaceAll(sqlStr, "#{nation}", nation)
	sqlStr = strings.ReplaceAll(sqlStr, "#{date}", fmt.Sprintf("%v", date))
	mysql.DB.Raw(sqlStr).Scan(&total)
	return
}

func GetUserMailMessageList(itemCode, appUserId, deviceType, nation string, date int64) (userMailMessageList []response.UserMailMessageResult) {
	sqlStr := `
		SELECT toms.mail_title as title,
      toms.notice_content AS content,
      toms.url AS jumpUrl,
      tom.send_time AS sendTime ,
      tomsl.is_show AS isFirstShow
      FROM t_notice_log tomsl
      LEFT JOIN t_notice_message toms ON toms.message_id = tomsl.message_id
      LEFT JOIN t_notice tom ON tom.mail_id = tomsl.mail_id
      WHERE
      tomsl.item_code = #{itemCode}
      AND tomsl.nation = #{nation}
      AND tomsl.decive_brand = #{deviceType}
      AND (tomsl.app_user_id = #{appUserId} OR tomsl.app_user_id = 'ALL')
      AND tom.send_state = 1
      AND tom.send_time <![CDATA[>=]]> #{date}
      AND (tom.send_time / 1000) >= UNIX_TIMESTAMP(DATE_SUB(CURDATE(),INTERVAL 7 DAY))
      ORDER BY tom.send_time DESC
`
	sqlStr = strings.ReplaceAll(sqlStr, "#{itemCode}", itemCode)
	sqlStr = strings.ReplaceAll(sqlStr, "#{appUserId}", appUserId)
	sqlStr = strings.ReplaceAll(sqlStr, "#{deviceType}", deviceType)
	sqlStr = strings.ReplaceAll(sqlStr, "#{nation}", nation)
	sqlStr = strings.ReplaceAll(sqlStr, "#{date}", fmt.Sprintf("%v", date))
	mysql.DB.Raw(sqlStr).Scan(&userMailMessageList)
	return
}

func UpdateMailStatus(itemCode, nation, deviceType, appUserId string) {
	sqlStr := `
		update t_notice_log tomsl
      set tomsl.read_state = 1
      where tomsl.is_show = 1
      AND tomsl.item_code = #{itemCode}
      AND tomsl.nation = #{nation}
      AND tomsl.decive_brand = #{deviceType}
      AND (tomsl.app_user_id = #{appUserId} OR tomsl.app_user_id = 'ALL')
`
	sqlStr = strings.ReplaceAll(sqlStr, "#{itemCode}", itemCode)
	sqlStr = strings.ReplaceAll(sqlStr, "#{nation}", nation)
	sqlStr = strings.ReplaceAll(sqlStr, "#{deviceType}", deviceType)
	sqlStr = strings.ReplaceAll(sqlStr, "#{appUserId}", appUserId)
	mysql.DB.Exec(sqlStr)
	return
}

func QueryUserVipInfo(appUserId string) (VipUserInfo response.VipUserInfo) {
	sqlStr := `
		SELECT
        b.user_level AS userLevel,
        IFNULL(b.new_avatar,b.avatar) AS avatar,
        a.vip_logo_url AS vipLogoUrl,
        a.recharge_icon AS rechargeIcon,
        IFNULL( SUM( c.amount ), 0 ) AS amount,
        IFNULL( count(c.recharge_id), 0 ) AS tradeTimes,
        b.app_user_id AS appUserId,
        a.vip_name AS vipName,
        a.recharge_standard AS rechargeStandard,
        a.trading_standard AS tradingStandard
        FROM
        t_setting_vip_rule a
        LEFT JOIN t_app_user b ON a.vip_level = b.user_level
        AND a.item_code = b.item_code
        AND a.nation = b.nation
        LEFT JOIN t_recharge_order c ON c.app_user_id = b.app_user_id
        AND c.recharge_order_status = 1
        AND c.STATUS = '02'
        WHERE
        b.app_user_id = #{appUserId}
`
	sqlStr = strings.ReplaceAll(sqlStr, "#{appUserId}", appUserId)
	mysql.DB.Raw(sqlStr).Scan(&VipUserInfo)
	return
}

func GetTradingAssets(assetsName string, index, ends *int) (TradingAssetsBean []response.TradingAssetsBean) {
	sqlStr := `
		select a.id,
		       a.assets_name as 'assetsName',
		       a.assets_code as 'assetsCode',
		       a.assets_logo_url as 'assetsLogoUrl',
		       a.return_rate as 'returnRate',
		       a.assets_status as 'assetsStatus',
			   a.is_valid as 'isValid',
			   a.is_show as 'isShow',
		       a.point_number as 'pointNumber',
		       a.assets_type_code as 'assetsTypeCode' ,
		       a.assets_type_name as 'assetsTypeName' 
		from t_trading_assets as a 
		where 1=1
		${assetsName}
		ORDER BY a.sort ASC
		LIMIT #{index},#{ends}
`
	sqlStr = strings.ReplaceAll(sqlStr, "${assetsName}", fmt.Sprintf(`and a.assets_name like '%%%v%%'`, assetsName))
	sqlStr = strings.ReplaceAll(sqlStr, "#{index}", fmt.Sprintf("%v", index))
	sqlStr = strings.ReplaceAll(sqlStr, "#{ends}", fmt.Sprintf("%v", ends))
	mysql.DB.Raw(sqlStr).Scan(&TradingAssetsBean)
	return
}

func GetTradingOrderVirtual(appUserId, itemCode, tradeOrderId, nation string, amountType []string, index, ends *int) (GetTradingOrderParams []response.GetTradingOrderParams) {
	sqlStr := `
		SELECT a.*,
		       b.point_number AS 'pointNumber'
		FROM(
	        select
	        a.trade_order_id AS tradeOrderId,
	        a.app_user_id AS appUserId,
	        a.amounts AS amounts,
	        a.plan_gain_amounts AS planGainAmounts,
			(a.real_gain_amounts + IFNULL(a.equal_return_amounts, 0))  AS realGainAmounts,
	        a.item_code AS itemCode,
	        a.assets_id AS assetsId,
	        a.assets_name AS assetsName,
	        a.assets_logo_url AS assetsLogoUrl,
	        a.begin_assets_price AS beginAssetsPrice,
	        a.end_assets_price AS endAssetsPrice,
	        a.trade_begin_date AS tradeBeginDate,
	        a.trade_end_date AS tradeEndDate,
	        a.quiz_result AS quizResult,
	        a.is_valid AS isValid,
	        a.create_date AS createDate,
	        a.update_date AS updateDate,
	        a.amount_type AS amountType,
	        a.up_down_type AS upDownType,
	        a.real_cost_amounts  AS realCostAmounts,
	        a.return_rate AS returnRate,
			a.currency_type AS currencyType
	        from t_trading_order_virtual a
	        where a.is_valid = 1
	        AND a.app_user_id = #{appUserId}
	        AND a.item_code = #{itemCode}
	        `
	if len(amountType) > 0 {
		sqlStr += ` AND a.amount_type in(#{amountType})`
	}
	if len(tradeOrderId) > 0 {
		sqlStr += ` and a.trade_order_id = #{tradeOrderId}`
	}
	if len(nation) > 0 {
		sqlStr += ` and a.nation = #{nation}`
	}
	sqlStr += `
        ) AS a 
		INNER JOIN t_trading_assets AS b
		ON a.assetsId = b.id
		where 1=1
	    ORDER BY a.createDate DESC 
		LIMIT #{index},#{ends}
`
	sqlStr = strings.ReplaceAll(sqlStr, "#{appUserId}", appUserId)
	sqlStr = strings.ReplaceAll(sqlStr, "#{itemCode}", itemCode)
	sqlStr = strings.ReplaceAll(sqlStr, "#{amountType}", fmt.Sprintf("%v", strings.Join(amountType, ",")))
	sqlStr = strings.ReplaceAll(sqlStr, "#{tradeOrderId}", tradeOrderId)
	sqlStr = strings.ReplaceAll(sqlStr, "#{nation}", nation)
	sqlStr = strings.ReplaceAll(sqlStr, "#{index}", fmt.Sprintf("%v", index))
	sqlStr = strings.ReplaceAll(sqlStr, "#{ends}", fmt.Sprintf("%v", ends))
	mysql.DB.Raw(sqlStr).Scan(&GetTradingOrderParams)
	return
}

func GetTradingOrder(appUserId, itemCode, tradeOrderId, nation string, Now int64, amountType []string, index, ends *int) (GetTradingOrderParams []response.GetTradingOrderParams) {
	sqlStr := `
		SELECT a.*,
		       b.point_number AS 'pointNumber'
		FROM(
		        select
	            a.trade_order_id AS tradeOrderId,
	            a.app_user_id AS appUserId,
	            a.amounts AS amounts,
	            a.plan_gain_amounts AS planGainAmounts,
				(a.real_gain_amounts + IFNULL(a.equal_return_amounts, 0)) AS realGainAmounts,
	            a.item_code AS itemCode,
	            a.assets_id AS assetsId,
	            a.assets_name AS assetsName,
	            a.assets_logo_url AS assetsLogoUrl,
	            a.begin_assets_price AS beginAssetsPrice,
	            a.end_assets_price AS endAssetsPrice,
	            a.trade_begin_date AS tradeBeginDate,
	            a.trade_end_date AS tradeEndDate,
	            a.quiz_result AS quizResult,
	            a.is_valid AS isValid,
	            a.create_date AS createDate,
	            a.update_date AS updateDate,
	            a.amount_type AS amountType,
	            a.up_down_type AS upDownType,
	            a.real_cost_amounts  AS realCostAmounts,
	            a.user_gain_amounts AS userGainAmounts,
	            a.return_rate AS returnRate,
	            a.currency_type AS currencyType
	        from t_trading_order a
	        where a.is_valid = 1
	            AND a.app_user_id = #{appUserId}
	            AND a.item_code = #{itemCode}
`
	if len(amountType) > 0 {
		sqlStr += ` AND a.amount_type in(#{amountType})`
	}
	if len(tradeOrderId) > 0 {
		sqlStr += ` and a.trade_order_id = #{tradeOrderId}`
	}
	if Now > 0 {
		sqlStr += ` and a.create_ts >= #{Now}}`
	}
	if len(nation) > 0 {
		sqlStr += ` and a.nation = #{nation}`
	}
	sqlStr += `
        ) AS a 
		INNER JOIN t_trading_assets AS b
		ON a.assetsId = b.id
		where 1=1
	    ORDER BY a.createDate DESC
		LIMIT #{index},#{ends}
`
	sqlStr = strings.ReplaceAll(sqlStr, "#{appUserId}", appUserId)
	sqlStr = strings.ReplaceAll(sqlStr, "#{itemCode}", itemCode)
	sqlStr = strings.ReplaceAll(sqlStr, "#{amountType}", fmt.Sprintf("%v", strings.Join(amountType, ",")))
	sqlStr = strings.ReplaceAll(sqlStr, "#{tradeOrderId}", tradeOrderId)
	sqlStr = strings.ReplaceAll(sqlStr, "#{Now}", fmt.Sprintf("%v", Now))
	sqlStr = strings.ReplaceAll(sqlStr, "#{nation}", nation)
	sqlStr = strings.ReplaceAll(sqlStr, "#{index}", fmt.Sprintf("%v", index))
	sqlStr = strings.ReplaceAll(sqlStr, "#{ends}", fmt.Sprintf("%v", ends))
	mysql.DB.Raw(sqlStr).Scan(&GetTradingOrderParams)
	return
}
