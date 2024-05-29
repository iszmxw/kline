package echo

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"strings"
)

// Response 统一响应结构
type Response struct {
	Code  int         `json:"code"`  // 响应码，1 表示成功，其他表示失败
	Msg   string      `json:"msg"`   // 响应消息，SUCCESS 表示成功
	Data  interface{} `json:"data"`  // 响应数据，有数据的时候返回数据，详情请看data数据结构体
	ReqID string      `json:"reqId"` // 请求 ID，用于追踪请求
}

// Rjson 成功返回封装 参数 data interface{} 类型为可接受任意类型
func Rjson(c *gin.Context, result interface{}, msg string, code int) {
	reqId, _ := c.Get("Tracking-Id")
	rdata := new(Response)
	rdata.ReqID = reqId.(string)
	rdata.Code = code
	if result != nil {
		rdata.Data = result
	}
	if len(msg) > 0 {
		rdata.Msg = msg
	} else {
		rdata.Msg = "success."
	}
	c.JSON(200, rdata)
	return
}

// Error  错误返回封装
func Error(c *gin.Context, codeName string, lastMsg string) {
	var logInfo []zapcore.Field
	if len(logger.RequestId) > 0 {
		logInfo = append(logInfo, zap.Any("RequestId", logger.RequestId))
	}
	// 本地枚举获取错误代码
	code, firstMsg := GetCode(codeName, c.Request.Header.Get("Language")) // 最终code
	if strings.EqualFold("未知错误", firstMsg) {
		// 数据库获取错误代码
		code, firstMsg = showWord(codeName, c.Request.Header.Get("Language"), c.Request.Header.Get("itemCode"))
	}
	fullMsg := firstMsg + lastMsg // 最终msg
	logInfo = append(logInfo, zap.Any(code, fullMsg))
	logger.Logger.WithOptions(zap.AddCallerSkip(1)).Info("返回错误", logInfo...)
	Rjson(c, nil, fullMsg, helpers.StringToInt(code))
}

// Success  错误返回封装
func Success(c *gin.Context, result interface{}, msg string) {
	var logInfo []zapcore.Field
	if len(logger.RequestId) > 0 {
		logInfo = append(logInfo, zap.Any("RequestId", logger.RequestId))
	}
	logInfo = append(logInfo, zap.Any("返回数据", result))
	logger.Logger.WithOptions(zap.AddCallerSkip(1)).Info("成功返回", logInfo...)
	Rjson(c, result, msg, 1)
}
