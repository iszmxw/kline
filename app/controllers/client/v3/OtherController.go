package v3

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/pkg/echo"
	"goapi/pkg/redis"
)

// k线图服务

type OtherController struct {
	BaseController
}

// 获取 15 库的所有键值对

func (h *OtherController) PushList(c *gin.Context) {
	index := c.Query("index")
	item := c.Query("item")
	if len(item) > 0 && len(index) > 0 {
		ok, err := redis.Client.LPushDb(15, fmt.Sprintf("list-mql5:%s", index), item)
		if err != nil {
			echo.Error(c, "Failed", err.Error())
			return
		}
		echo.Success(c, gin.H{"result": ok}, "")
		return
	}
	echo.Error(c, "Failed", "")
}
