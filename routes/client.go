package routes

import (
	"github.com/gin-gonic/gin"
	clientV1 "goapi/app/controllers/client/v1"
)

// RegisterClientRoutes 注册Client路由
func RegisterClientRoutes(router *gin.RouterGroup) {
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
				// 【C11】 设置 config 信息
				trade.Any("/setConfig.json", ClientV1Group.KlineController.SetConfigHandler)
				// 【C12】 查看配置
				trade.Any("/getConfig.json", ClientV1Group.KlineController.GetConfigHandler)
				trade.Any("/getRedisKey.json", ClientV1Group.KlineController.GetRedisKeyHandler)
			}
		}

		{ // V2 版本

		}

		{ // V3 版本

		}

	}
}
