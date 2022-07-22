package task_plan

import (
	"fmt"
	"time"
)

/**
 * 活动
 * 1、每日8点整开始执行，赠送Bonus入账
 */

func ActivityInvitationGiftBonus() {
	fmt.Println("1、每日8点整开始执行，赠送Bonus入账", time.Now())
}

/**
 * 活动
 * 2、每天0点1分开始执行，汇总交易亏损数据
 */

func ActivityLossReturnActivity() {
	fmt.Println("2、每天0点1分开始执行，汇总交易亏损数据", time.Now())
}
