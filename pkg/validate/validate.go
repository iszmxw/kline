package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"goapi/pkg/echo"
)

// 校验参数

func ParamsError(c *gin.Context, params interface{}) bool {
	ValidateLang := new(echo.ValidateLang)
	ValidateLang.Validate = validator.New()
	vErr := ValidateLang.SetLang("zh").Validate.Struct(params)
	if vErr != nil {
		msg := ValidateLang.Translate(vErr, params)
		echo.Error(c, "PARAMS_ERROR", "："+msg)
		return true
	}
	return false
}
