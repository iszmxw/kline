package routes

import (
	"github.com/gin-gonic/gin"
	WebCtl "goapi/app/controllers/web/v1"
	WebMw "goapi/app/middlewares/v1"
)

var WebV1Group = new(WebCtl.Group)

// RegisterWebRoutes 注册路由
func RegisterWebRoutes(ApiRoute *gin.RouterGroup) {
	// 路由分组 Web 接口
	ApiRoute.Use(WebMw.Web())
	WebV1 := ApiRoute.Group("/web/v1")
	// 看板模块
	Kanban := WebV1.Group("/kanban")
	{
		// 【E31】 wsIndex.json
		Kanban.Any("/wsIndex.json", WebV1Group.KlineController.WsIndexHandler)
		// 【E32】 历史行情-mysql读取
		Kanban.Any("/getHisKlineList.json", WebV1Group.KlineController.HistoryHandler)
		// 【E33】 历史行情-es读取
		Kanban.Any("/getEsHisKlineList", WebV1Group.KlineController.HistoryEsHandler)
		// 【E34】 设置 config 信息
		WebV1.Any("/setConfig.json", WebV1Group.KlineController.SetConfigHandler)
		// 【E35】 查看配置
		WebV1.Any("/getConfig.json", WebV1Group.KlineController.GetConfigHandler)
		// 【E36】
		WebV1.Any("/getRedisKey.json", WebV1Group.KlineController.GetRedisKeyHandler)
	}
}
