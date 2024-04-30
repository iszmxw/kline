package v3

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"goapi/app/models"
	"goapi/app/requests"
	"goapi/pkg/echo"
	"goapi/pkg/helpers"
	"goapi/pkg/mysql"
	"goapi/pkg/validate"
)

// 通知公告服务

type UserController struct {
	BaseController
}

// RegisterHandler 注册
func (h *UserController) RegisterHandler(c *gin.Context) {
	HeaderParams := getHeaderParams(c)
	fmt.Println(HeaderParams.IsSecret, HeaderParams.Sign, HeaderParams.Token)
	var params requests.UserRegisterParams
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	// 校验邮箱的格式
	if !helpers.IsEmail(params.Account) {
		echo.Error(c, "ParamsError", "：邮箱格式不正确")
		return
	}
	// 检查用户是否已经注册
	var appUser models.AppUser
	mysql.DB.Debug().Model(appUser).Where("email", params.Account).Find(&appUser)
	if len(appUser.AppUserId) > 0 {
		// PleaseLoginOrRetrievePassword ，注册检测，已注册用户，提示用户去登录或者去找回密码
		echo.Error(c, "PleaseLoginOrRetrievePassword", "")
		return
	}
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
