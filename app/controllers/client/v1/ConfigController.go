package v1

import (
	"goapi/app/models"
	"goapi/app/requests"
	"goapi/app/response"
	"goapi/pkg/echo"
	"goapi/pkg/mysql"
	"goapi/pkg/validate"

	"github.com/gin-gonic/gin"
)

type ConfigController struct {
	BaseController
}

func (h *ConfigController) GetActiveHostHandler(c *gin.Context) {
	var params requests.Base
	var result []response.AppActiveHostResult
	_ = c.Bind(&params)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	mysql.DB.Debug().Model(models.AppActiveHost{}).Where("nation", params.Nation).Select("active_host").Find(&result)
	echo.Success(c, result, "")
}

func (h *ConfigController) GetLanguagePageHandler(c *gin.Context) {

}

func (h *ConfigController) GetLanguagePageV2Handler(c *gin.Context) {

}

func (h *ConfigController) CheckAppVersionHandler(c *gin.Context) {

}

func (h *ConfigController) GetTutorialByTypeHandler(c *gin.Context) {

}

func (h *ConfigController) GetLoginCodeByTestHandler(c *gin.Context) {

}

func (h *ConfigController) ClearTestAccountHandler(c *gin.Context) {

}

func (h *ConfigController) GetExistingCountryHandler(c *gin.Context) {

}

func (h *ConfigController) GetImagesHandler(c *gin.Context) {

}
