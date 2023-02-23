package client_v2

import (
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

			// 代理落地页接口
			config := V2Route.Group("/config")
			{
				// 设置token
				config.Any("/setKlineToken.json", ClientV2Group.ConfigController.SetKlineToken)
			}
		}

	}
}
