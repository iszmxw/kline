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
		}

		{ // V3 版本

		}

	}
}
