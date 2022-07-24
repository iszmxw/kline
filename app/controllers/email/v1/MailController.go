package v1

import (
	"github.com/gin-gonic/gin"
	"goapi/app/requests"
	"goapi/pkg/echo"
	"goapi/pkg/validate"
)

// k线图服务

type MailController struct {
	BaseController
}

/**接口编号:【U01】
 * .发送邮件【POST】
 */
// MailController
// 历史行情 mysql 查询

func (h *MailController) SendEmail(c *gin.Context) {
	var params requests.SendEmailParams
	_ = c.Bind(&params)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	echo.Success(c, nil, "")
}
