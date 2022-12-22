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
		Kanban.Any("/getEsHisKlineList.json", WebV1Group.KlineController.HistoryEsHandler)
	}

	// 设置模块
	Config := WebV1.Group("/config")
	{
		// 【E34】 所有资产对应的数据源
		Config.Any("/AllowKlineSource.json", WebV1Group.KlineController.AllowKlineSource)
		// 【E34】 获取当前币种可以设置的数据源
		Config.Any("/KlineSourceList.json", WebV1Group.KlineController.KlineSourceList)
		// 【E34】 设置 KlineSource 信息
		Config.Any("/setKlineSource.json", WebV1Group.KlineController.SetKlineSourceHandler)
		// 【E35】 查看所有 KlineSource 配置
		Config.Any("/getAllKlineSource.json", WebV1Group.KlineController.GetAllKlineSourceHandler)
		// 【E36】 所有 redis 的Key
		Config.Any("/getRedisKey.json", WebV1Group.KlineController.GetRedisKeyHandler)
		// 【E37】 设置token
		Config.Any("/setKlineToken.json", WebV1Group.KlineController.SetKlineToken)
	}

}
