package v3

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/app/requests"
	"goapi/pkg/echo"
	"goapi/pkg/validate"
)

// 通知公告服务

type UserController struct {
	BaseController
}

// RegisterHandler 注册
func (h *UserController) RegisterHandler(c *gin.Context) {
	deviceidentity := c.GetHeader("deviceidentity")
	issecert := c.GetHeader("issecert")
	sign := c.GetHeader("sign")
	token := c.GetHeader("token")
	fmt.Println(issecert, sign, token)
	if len(deviceidentity) > 0 {

	}
	//res := address.GetMessageByIp(c.ClientIP())
	var params requests.UserRegisterParams
	_ = c.Bind(&params)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	c.Request.Header.Set("Language", params.LanguageCode)
	echo.Success(c, nil, "SUCCESS")
}

// LoginHandler 登录
func (h *UserController) LoginHandler(c *gin.Context) {

}

// LogoutHandler 退出
func (h *UserController) LogoutHandler(c *gin.Context) {

}

// CancelingAccountHandler 注销
func (h *UserController) CancelingAccountHandler(c *gin.Context) {

}

// UnbindTokenHandler 解绑推送令牌
func (h *UserController) UnbindTokenHandler(c *gin.Context) {

}

// ChangePasswordHandler 修改密码
func (h *UserController) ChangePasswordHandler(c *gin.Context) {

}
