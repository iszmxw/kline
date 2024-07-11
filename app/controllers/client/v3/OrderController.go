package v3

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"goapi/app/models/mapper"
	"goapi/app/requests"
	"goapi/app/response"
	"goapi/app/service/login"
	"goapi/pkg/decimal"
	"goapi/pkg/echo"
	"goapi/pkg/helpers"
	"goapi/pkg/page"
	"goapi/pkg/validate"
	"goapi/pkg/xstring"
	"goapi/pkg/xtime"
	"strings"
	"time"
)

// 银行模块

type OrderController struct {
	BaseController
}

// CreatePaidCodeHandler
// @Summary No.06.1-生成充值码
// @Description 生成充值码的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 7.订单模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param GetLanguagePageParams body requests.GetLanguagePageParams true "生成充值码的参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.SuccessLanguagePageV2 "data数据结构体"
// @Router /No.06.1/app/v3/order/createPaidCode.json [post]
func (h *OrderController) CreatePaidCodeHandler(c *gin.Context) {
	var params requests.GetLanguagePageParams
	var result response.SuccessLanguagePageV2
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	echo.Success(c, result, "ok")
}

// SubmitWithdrawHandler
// @Summary No.06.2-提交提现
// @Description 提交提现的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 7.订单模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param GetLanguagePageParams body requests.GetLanguagePageParams true "提交提现的参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.SuccessLanguagePageV2 "data数据结构体"
// @Router /No.06.2/app/v3/order/submitWithdraw.json [post]
func (h *OrderController) SubmitWithdrawHandler(c *gin.Context) {
	var params requests.GetLanguagePageParams
	var result response.SuccessLanguagePageV2
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	echo.Success(c, result, "ok")
}

// OrderDetailHandler
// @Summary No.06.3-获取订单详细信息
// @Description 获取订单详细信息的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 7.订单模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param GetLanguagePageParams body requests.GetLanguagePageParams true "获取订单详细信息的参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.SuccessLanguagePageV2 "data数据结构体"
// @Router /No.06.3/app/v3/order/orderDetail.json [post]
func (h *OrderController) OrderDetailHandler(c *gin.Context) {
	var params requests.GetLanguagePageParams
	var result response.SuccessLanguagePageV2
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	echo.Success(c, result, "ok")
}

// GetCashFlowHandler
// @Summary No.06.4-获取资金流水列表
// @Description 获取资金流水列表的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 7.订单模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param CashFlowParams body requests.CashFlowParams true "获取资金流水列表的参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.GetCashFlow "data数据结构体"
// @Router /No.06.4/app/v3/order/getCashFlow.json [post]
func (h *OrderController) GetCashFlowHandler(c *gin.Context) {
	var params requests.CashFlowParams
	var result response.GetCashFlow
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	if params.PageNum == 0 || params.PageSize == 0 {
		params.PageNum = 1
		params.PageSize = 5
	}
	// 国家
	nation := params.Nation
	// 保留小数位数
	appItems := login.GetAppItem(params.Nation)
	numberPoint := appItems.NumberPoint
	Now := xtime.GetTodayStartTime(appItems.TimeZone)
	user, _ := c.Get("user")
	pagea := page.NewPageApple(params.PageSize, params.PageNum)
	total := mapper.CashFlowListCount(params.ItemCode, nation, params.Today, params.OrderType, user.(response.ClientUserBeans).AppUserId, Now.UnixMilli(), pagea.Index, pagea.Ends)
	list := mapper.CashFlowList(params.ItemCode, nation, params.Today, params.OrderType, user.(response.ClientUserBeans).AppUserId, Now.UnixMilli(), pagea.Index, pagea.Ends)
	arr := []string{"成功", "已过期"}
	for i, flow := range list {
		if helpers.InArray(flow.FailureReason, arr) {
			// 去除中文
			list[i].FailureReason = ""
		}
		if xstring.IsAllChinese(flow.FailureReason) {
			// 去除中文
			list[i].FailureReason = ""
		}
		list[i].IsUpload = "0"
		// deposit订单用户支付5分钟后，还没有到账则显示按钮
		if strings.EqualFold("01", flow.Status) && (time.Now().UnixMilli()-flow.SubmitTime > 5*60*1000) && len(flow.PhotoPaymentVoucher) == 0 {
			list[i].IsUpload = "1"
		}
		list[i].ShowAmount = decimal.SplitDisplay(flow.Amount, nation, numberPoint)
	}
	result.List = list
	result.Total = total
	echo.Success(c, result, "ok")
}
