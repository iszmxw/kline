package routes

import (
	clientV1 "goapi/app/controllers/client/v1"
	clientV2 "goapi/app/controllers/client/v2"
	middlewares "goapi/app/middlewares/v1"

	"github.com/gin-gonic/gin"
)

// RegisterClientRoutes 注册Client路由
func RegisterClientRoutes(router *gin.RouterGroup) {
	router.Use(middlewares.AddressLimit()) // 地区限制
	//router.Use(middlewares.Client())
	// 路由分组 客户端 模块
	AppRoute := router.Group("/app")
	{
		{ // V1 版本
			ClientV1Group := new(clientV1.Group)
			V1Route := AppRoute.Group("/v1")
			// k线图服务
			trade := V1Route.Group("/trade")
			{
				// 【C07】 历史行情-mysql读取
				trade.Any("/getHisKlineList.json", ClientV1Group.KlineController.HistoryHandler)
				// 【C08】 历史行情-es读取
				trade.Any("/getEsHisKlineList.json", ClientV1Group.KlineController.HistoryEsHandler)
				// 【C09】 wsIndex.json
				trade.Any("/wsIndex.json", ClientV1Group.KlineController.WsIndexHandler)
				// 【C10】 wsReturnRate.json
				trade.Any("/wsReturnRate.json", ClientV1Group.KlineController.WsReturnRateHandler)
				// 删除kline_log表数据，维护数据表过大问题
				trade.Any("/maintenanceTableData.json", ClientV1Group.KlineController.MaintenanceTableData)
			}

			// 【A】默认模块
			{
				// 【A01】获取域名池【POST】
				AppRoute.Group("/v1/config").POST("/getActiveHost.json", ClientV1Group.ConfigController.GetActiveHostHandler)
				// 【A02】获取系统多语言包【POST】
				AppRoute.Group("/v1/config").POST("/getLanguagePage.json", ClientV1Group.ConfigController.GetLanguagePageHandler)
				// 【A03】校验app版本【POST】
				AppRoute.Group("/config/v1").POST("/checkAppVersion.json", ClientV1Group.ConfigController.CheckAppVersionHandler)
				// 【A04】获取引导教程【POST】
				AppRoute.Group("/v1/config").POST("/getTutorialByType.json", ClientV1Group.ConfigController.GetTutorialByTypeHandler)
				// 【A05】获取测试账号发送验证码【POST】
				AppRoute.Group("/v1/config").POST("/getLoginCodeByTest.json", ClientV1Group.ConfigController.GetLoginCodeByTestHandler)
				// 【A06】清除测试账号【POST】
				AppRoute.Group("/v1/config").POST("/clearTestAccount.json", ClientV1Group.ConfigController.ClearTestAccountHandler)
				// 【A09】获取国家信息【POST】
				AppRoute.Group("/v1/config").POST("/getExistingCountry.json", ClientV1Group.ConfigController.GetExistingCountryHandler)
				// 【A10】获取系统多语言包V2【POST】
				AppRoute.Group("/v1/config").POST("/getLanguagePageV2.json", ClientV1Group.ConfigController.GetLanguagePageV2Handler)
				// 【A11】获取系统展示图片【POST】
				AppRoute.Group("/v1/config").POST("/getImages.json", ClientV1Group.ConfigController.GetImagesHandler)
			}

		}

		{ // V2 版本
			ClientV2Group := new(clientV2.Group)
			V2Route := AppRoute.Group("/v2")
			// k线图服务
			trade := V2Route.Group("/trade")
			{
				// 【C07】 历史行情-mysql读取
				trade.Any("/getHisKlineList.json", ClientV2Group.KlineController.HistoryHandler)
				// 【C08】 历史行情-es读取
				trade.Any("/getEsHisKlineList.json", ClientV2Group.KlineController.HistoryEsHandler)
				// 【C09】 wsIndex.json
				trade.Any("/wsIndex.json", ClientV2Group.KlineController.WsIndexHandler)
				// 【C10】 wsReturnRate.json
				trade.Any("/wsReturnRate.json", ClientV2Group.KlineController.WsReturnRateHandler)
			}

			// 通知
			notice := V2Route.Group("/notice")
			{
				// 站内信通知
				notice.Any("/siteLetterNotification.json", ClientV2Group.NoticeController.StationLetterHandler)
			}

			// 代理落地页接口
			agent := V2Route.Group("/agent")
			{
				// 落地页app注册模块
				agent.Any("/register.json", ClientV2Group.AgentController.RegisterHandler)
				// 博眼球，数据造假，时时推送假数据
				agent.Any("/statistics.json", ClientV2Group.AgentController.StatisticsHandler)
			}
		}

		{ // V3 版本

		}

	}
}
