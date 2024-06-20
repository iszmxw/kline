package v3

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"goapi/app/requests"
	"goapi/app/response"
	"goapi/pkg/echo"
	"goapi/pkg/validate"
)

// 银行模块

type BankController struct {
	BaseController
}

// GetUserBankCardsHandler
// @Summary No.06.1-获取用户提现银行卡列表
// @Description 获取用户提现银行卡列表的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 6.银行模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param PageInfo body requests.PageInfo true "获取用户提现银行卡列表的参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.BankCards "data数据结构体"
// @Router /No.06.1/app/v3/bank/getUserBankCards.json [post]
func (h *BankController) GetUserBankCardsHandler(c *gin.Context) {
	var params requests.PageInfo
	var result response.BankCards
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}

	echo.Success(c, result, "ok")
}

// DeleteUserBankCardHandler
// @Summary No.06.2-删除用户银行卡
// @Description 删除用户银行卡的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 6.银行模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param SetUserBankCardParams body requests.SetUserBankCardParams true "删除用户银行卡的参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Router /No.06.2/app/v3/bank/deleteUserBankCard.json [post]
func (h *BankController) DeleteUserBankCardHandler(c *gin.Context) {
	var params requests.SetUserBankCardParams
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	echo.Success(c, nil, "ok")
}

// GetPaidFastAmountsHandler
// @Summary No.06.3-获取充值快捷金额列表
// @Description 获取充值快捷金额列表的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 6.银行模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param GetPaidFastAmountsParams body requests.GetPaidFastAmountsParams true "获取充值快捷金额列表的参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.RechargeFastAmountResult "data数据结构体"
// @Router /No.06.3/app/v3/bank/getPaidFastAmounts.json [post]
func (h *BankController) GetPaidFastAmountsHandler(c *gin.Context) {
	var params requests.GetPaidFastAmountsParams
	var result response.RechargeFastAmountResult
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	echo.Success(c, result, "ok")
}

// GetPaidTypeHandler
// @Summary No.06.4-获取支付方式
// @Description 获取支付方式的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 6.银行模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param Params body requests.Base true "获取支付方式的参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.RechargeTypeBeans "data数据结构体"
// @Router /No.06.4/app/v3/bank/getPaidType.json [post]
func (h *BankController) GetPaidTypeHandler(c *gin.Context) {
	var params requests.Base
	var result response.RechargeTypeBeans
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	echo.Success(c, result, "ok")
}
