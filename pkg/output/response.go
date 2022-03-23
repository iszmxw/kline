package output

import (
	"github.com/gin-gonic/gin"
)

// Rjson 成功返回封装 参数 data interface{} 类型为可接受任意类型
func Rjson(c *gin.Context, result interface{}, msg string, code int64) {
	reqId, _ := c.Get("Tracking-Id")
	c.JSON(200, gin.H{
		//返回数据
		"code":    code,
		"data":    result,
		"reqId":   reqId,
		"message": msg,
	})
	c.Abort()
	return
}
