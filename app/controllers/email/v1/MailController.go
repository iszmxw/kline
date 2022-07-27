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

func (h *MailController) QueryMargin(c *gin.Context) {
	var params requests.QueryMargin
	_ = c.Bind(&params)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	echo.Success(c, nil, "")
}

/**接口编号:【U01】
 * .发送邮件【POST】
 */
// MailController

func (h *MailController) SendOneEmail(c *gin.Context) {
	var params requests.SendOneEmail
	_ = c.Bind(&params)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	echo.Success(c, nil, "")
}

/**接口编号:【U01】
 * .发送邮件【POST】
 */
// MailController

func (h *MailController) SendBatchEmail(c *gin.Context) {
	var params requests.SendBatchEmail
	_ = c.Bind(&params)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	echo.Success(c, nil, "")
}

/**接口编号:【U01】
 * .发送邮件【POST】
 */
// MailController

func (h *MailController) ResetMargin(c *gin.Context) {
	echo.Success(c, nil, "")
}
