package task_route

import (
	"github.com/robfig/cron/v3"
	"goapi/service/binary-timer/task_plan"
)

// RegisterTaskRoutes 注册定时任务路由

func RegisterTaskRoutes(crontab *cron.Cron) {
	// 4、【定时任务】重置邮箱剩余条数 执行时间/周期：- 每天的凌晨12:02分
	_, _ = crontab.AddFunc("*/3 * * * * ?", task_plan.ActivityInvitationGiftBonus)
	// 【活动】1、每日8点整开始执行，赠送Bonus入账
	_, _ = crontab.AddFunc("*/3 * * * * ?", task_plan.ActivityInvitationGiftBonus)
	// 【活动】2、每天0点1分开始执行，汇总交易亏损数据
	_, _ = crontab.AddFunc("*/5 * * * * ?", task_plan.ActivityLossReturnActivity)
	// 3、每10分钟执行一次验证码失效处理
	_, _ = crontab.AddFunc("*/5 * * * * ?", task_plan.ClientRefreshRechargeTimeoutStatus)
	// 4、每天零时1分执行，后台看板活跃数据
	_, _ = crontab.AddFunc("*/5 * * * * ?", task_plan.ClientRealTradeResult)
	// 5、同步分时间段统计注册、充值、交易 定时任务 每隔1分钟
	_, _ = crontab.AddFunc("*/5 * * * * ?", task_plan.ClientDemoTradeResult)
	// 6、每2秒钟检查站内信发送状态并更新
	_, _ = crontab.AddFunc("*/5 * * * * ?", task_plan.OthersVerificationStatus)
	// 7、每分钟执行一次，自动刷新充值超时状态定时任务
	_, _ = crontab.AddFunc("*/5 * * * * ?", task_plan.WebStatisticsYesterday)
	// 8、每隔2秒自动刷新真实交易结果定时任务
	_, _ = crontab.AddFunc("*/5 * * * * ?", task_plan.WebCheckNewUserData)
	// 9、每2秒执行一次自动刷新模拟交易结果定时任务
	_, _ = crontab.AddFunc("*/5 * * * * ?", task_plan.WebSiteInformation)

}
