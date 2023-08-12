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
				// 生产队列数据
				other.Any("/pushList.json", ClientV3Group.OtherController.PushList)
				// 测试mongo
				other.Any("/mongo.json", ClientV3Group.OtherController.Mongo)
			}
		}

	}
}
