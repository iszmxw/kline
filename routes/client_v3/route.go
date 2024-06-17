package client_v3

import (
	clientV3 "goapi/app/controllers/client/v3"
	middlewaresV3 "goapi/app/middlewares/v3"

	"github.com/gin-gonic/gin"
)

// RegisterClientRoutes 注册Client路由
func RegisterClientRoutes(router *gin.RouterGroup) {
	router.Use(middlewaresV3.AddressLimit()) // 地区限制
	router.Use(middlewaresV3.Client())
	// 路由分组 客户端 模块
	AppRoute := router.Group("/app")
	{
		{ // V3 版本
			ClientV3Group := new(clientV3.Group)
			V3Route := AppRoute.Group("/v3")
			// k线图服务
			trade := V3Route.Group("/trade")
			{
				// 【C09】 wsIndex.json
				trade.Any("/wsIndex.json", ClientV3Group.KlineController.WsIndexHandler)
			}

			// 代理落地页接口
			agent := V3Route.Group("/agent")
			{
				// 落地页app注册模块
				agent.Any("/register.json", ClientV3Group.AgentController.RegisterHandler)
			}

			// 邀请推广落地页接口
			promote := V3Route.Group("/promote")
			{
				// 落地页app注册模块
				promote.Any("/register.json", ClientV3Group.PromoteController.RegisterHandler)
				// 博眼球，数据造假，时时推送假数据
				promote.Any("/statistics.json", ClientV3Group.AgentController.StatisticsHandler)
			}

			// 设置token接口
			config := V3Route.Group("/config")
			{
				// 设置token
				config.Any("/setKlineToken.json", ClientV3Group.ConfigController.SetKlineToken)
				// 检查多语言是否包含中文，然后发送警告
				config.Any("/checkLanguage.json", ClientV3Group.ConfigController.CheckLanguage)
				// 清除 adId
				config.Any("/delAdId.json", ClientV3Group.ConfigController.DelAdId)
				// iqoption 的登录 token 设置
				config.Any("/iqoption.json", ClientV3Group.ConfigController.Iqoption)
			}

			// 其他接口
			other := V3Route.Group("/other")
			{
				// 接口信息
				other.Any("/info.json", ClientV3Group.OtherController.Info)
				// 生产队列数据
				other.Any("/pushList.json", ClientV3Group.OtherController.PushList)
				// 测试mongo
				other.Any("/mongo.json", ClientV3Group.OtherController.Mongo)
			}
			// TODO:: 新接口开发中
			// 1.用户模块
			user := V3Route.Group("/user")
			{
				// 注册 [ok]
				user.Any("/register.json", ClientV3Group.UserController.RegisterHandler)
				// 登录 [ok]
				user.Any("/login.json", ClientV3Group.UserController.LoginHandler)
				// 退出 [ok]
				user.Any("/logout.json", ClientV3Group.UserController.LogoutHandler)
				// 修改密码 [ok]
				user.Any("/changePassword.json", ClientV3Group.UserController.ChangePasswordHandler)
				// 注销 [ok]
				user.Any("/cancelingAccount.json", ClientV3Group.UserController.CancelingAccountHandler)
				// 多类型获取验证码 [ok]
				user.Any("/getCodeByType.json", ClientV3Group.UserController.GetCodeByTypeHandler)
			}

			// 2.绑定模块
			bind := V3Route.Group("/bind")
			{
				// 绑定推送令牌 [ok]
				bind.Any("/bindToken.json", ClientV3Group.BindController.BindTokenHandler)
				// 解绑推送令牌 [ok]
				bind.Any("/unbindToken.json", ClientV3Group.BindController.UnbindTokenHandler)
				// 绑定用户邮箱 [ok]
				bind.Any("/bindAccountEmail.json", ClientV3Group.BindController.BindAccountEmailHandler)
			}

			// 3.配置模块
			{
				// 获取系统多语言包 v2 [ok]
				config.Any("/getLanguagePageV2.json", ClientV3Group.ConfigController.GetLanguagePageV2Handler)
				// 获取国家信息 [ok]
				config.Any("/getExistingCountryV2.json", ClientV3Group.ConfigController.GetExistingCountryV2Handler)
				// 获取用户限制规则列表 [ok]
				config.Any("/getUserOperationRules.json", ClientV3Group.ConfigController.GetUserOperationRulesHandler)
				// 获取配置信息 [ok]
				config.Any("/getUserTradingRules.json", ClientV3Group.ConfigController.GetUserTradingRulesHandler)
				// 获取数字货币提现相关的信息 [ok]
				config.Any("/getBlockWithdrawInfo.json", ClientV3Group.ConfigController.GetBlockWithdrawInfoHandler)
				// 获取支持提现银行列表 [ok]
				config.Any("/getUserWithdrawList.json", ClientV3Group.ConfigController.GetUserWithdrawListHandler)
			}

			// 4.活动模块
			activity := V3Route.Group("/activity")
			{
				// 获取用户亏损返还数据 v2 [ok]
				activity.Any("/getUserLossReturn.json", ClientV3Group.ActivityController.GetUserLossReturnHandler)
				// 获取 Bonus 类型 [ok]
				activity.Any("/getBonusTypes.json", ClientV3Group.ActivityController.GetBonusTypesHandler)
				// 获取 Bonus 赠送记录 [ok]
				activity.Any("/getBonusRecords.json", ClientV3Group.ActivityController.GetBonusRecordsHandler)
				// 获取 banner 数据 [ok]
				activity.Any("/getBanners.json", ClientV3Group.ActivityController.GetBannersHandler)
				// 获取 FAQ 管理列表 [ok]
				activity.Any("/getFAQManageList.json", ClientV3Group.ActivityController.GetFAQManageListHandler)
				// 获取客户排名列表 [ok]
				activity.Any("/getUserRank.json", ClientV3Group.ActivityController.GetUserRankHandler)
				// 获取本周签到列表v2 [ok]
				activity.Any("/getWeekDataV2.json", ClientV3Group.ActivityController.GetWeekDataV2Handler)
				// 签到启停 firebase 通知 [ok]
				activity.Any("/signNotify.json", ClientV3Group.ActivityController.SignNotifyHandler)
				// 每日签到 [todo]
				activity.Any("/submit.json", ClientV3Group.ActivityController.SubmitHandler)
			}

			// 5.我的模块
			mine := V3Route.Group("/mine")
			{
				// 获取客户账号信息 [todo]
				mine.Any("/getUserAccountMessage.json", ClientV3Group.MineController.GetUserAccountMessageHandler)
				// 更新用户信息 [todo]
				mine.Any("/submitUpdateUserInfo.json", ClientV3Group.MineController.SubmitUpdateUserInfoHandler)
				// 上报新手教程完成状态 [todo]
				mine.Any("/sumbmitPartnerEventByType.json", ClientV3Group.MineController.SumbmitPartnerEventByTypeHandler)
				// 获取邀请用户进度 [todo]
				mine.Any("/getInvitUsers.json", ClientV3Group.MineController.GetInvitUsersHandler)
				// 获取用户邀请码及父邀请码 [todo]
				mine.Any("/getInvitationCode.json", ClientV3Group.MineController.GetInvitationCodeHandler)
				// 获取站内信列表 [todo]
				mine.Any("/getUserMailMessage.json", ClientV3Group.MineController.GetUserMailMessageHandler)
				// 上传照片 [todo]
				mine.Any("/uploadPhoto.json", ClientV3Group.MineController.UploadPhotoHandler)
			}

			// 6.银行模块
			bank := V3Route.Group("/bank")
			{
				// 获取用户提现银行卡列表 [todo]
				bank.Any("/getUserBankCards.json", ClientV3Group.BankController.GetUserBankCardsHandler)
				// 删除用户银行卡 [todo]
				bank.Any("/deleteUserBankCard.json", ClientV3Group.BankController.DeleteUserBankCardHandler)
				// 获取充值快捷金额列表 [todo]
				bank.Any("/getPaidFastAmounts.json", ClientV3Group.BankController.GetPaidFastAmountsHandler)
				// 获取支付方式 [todo]
				bank.Any("/getPaidType.json", ClientV3Group.BankController.GetPaidTypeHandler)
			}

			// 7.订单模块
			order := V3Route.Group("/order")
			{
				// 生成充值码 [todo]
				order.Any("/createPaidCode.json", ClientV3Group.OrderController.CreatePaidCodeHandler)
				// 提交提现 [todo]
				order.Any("/submitWithdraw.json", ClientV3Group.OrderController.SubmitWithdrawHandler)
				// 获取订单详细信息 [todo]
				order.Any("/orderDetail.json", ClientV3Group.OrderController.OrderDetailHandler)
				// 获取资金流水列表 [todo]
				order.Any("/getCashFlow.json", ClientV3Group.OrderController.GetCashFlowHandler)
			}

			// 8.VIP模块
			vip := V3Route.Group("/vip")
			{
				// 查询用户 VIP 信息 [todo]
				vip.Any("/getUserVip.json", ClientV3Group.VipController.GetUserVipHandler)
				// 个人中心页会员标识列表 [todo]
				vip.Any("/getSettingVip.json", ClientV3Group.VipController.GetSettingVipHandler)
			}

			// 9.认证模块
			credit := V3Route.Group("/credit")
			{
				// 提交邮箱、身份认证 [todo]
				credit.Any("/submitCredit.json", ClientV3Group.CreditController.SubmitCreditHandler)
				// 获取用户认证状态和认证信息 [todo]
				credit.Any("/getUserCreditMessage.json", ClientV3Group.CreditController.GetUserCreditMessageHandler)
			}

			// 10.交易模块
			{
				// 获取资产列表 [todo]
				trade.Any("/getAssetsList.json", ClientV3Group.TradeController.GetAssetsListHandler)
				// 查询交易历史 [todo]
				trade.Any("/getTradeHistorys.json", ClientV3Group.TradeController.GetTradeHistoryHandler)
			}

		}

	}
}
