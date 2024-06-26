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
