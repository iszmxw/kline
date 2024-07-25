package v3

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/app/models/SubmitTypeCode"
	"goapi/app/requests"
	"goapi/app/response"
	"goapi/app/service/credit"
	"goapi/app/service/user"
	"goapi/pkg/commonUtils"
	"goapi/pkg/echo"
	"goapi/pkg/email"
	"goapi/pkg/logger"
	"goapi/pkg/mobile"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
	"goapi/pkg/validate"
	"strings"
)

// 银行模块

type CreditController struct {
	BaseController
}

// SubmitCreditHandler
// @Summary No.09.1-提交邮箱、身份认证
// @Description 提交邮箱、身份认证的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 9.认证模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param SubmitCreditParams body requests.SubmitCreditParams true "提交邮箱、身份认证的参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.SuccessLanguagePageV2 "data数据结构体"
// @Router /No.09.1/app/v3/credit/submitCredit.json [post]
func (h *CreditController) SubmitCreditHandler(c *gin.Context) {
	var params requests.SubmitCreditParams
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	appUser, _ := c.Get("user")
	appUserId := appUser.(response.ClientUserBeans).AppUserId
	tAppUser := user.GetUser(appUserId)
	if tAppUser == nil {
		echo.Error(c, "UserMessageWrong", "")
		return
	}
	var creditData models.AppUserCredit
	mysql.DB.Debug().Model(models.AppUserCredit{}).Where("app_user_id", appUserId).Find(&creditData)
	switch params.SubmitType {
	case SubmitTypeCode.Info(SubmitTypeCode.MobileCredit).Code:
		{
			if len(params.Account) == 0 {
				echo.Error(c, "ParamsCanNotBeNull", "")
				return
			}
			// 拦截已认证完客户，不能修改认证信息
			if creditData.IsFinishCredit == 1 {
				echo.Error(c, "InformationAlreadyExists", "")
				return
			}
			params.Account = mobile.HandleMobileParams(params.Account, params.Nation)
			credit.UpdateMobile(c, params, tAppUser)
		}
		break
	case SubmitTypeCode.Info(SubmitTypeCode.EmailCredit).Code:
		{
			if len(params.Account) == 0 {
				echo.Error(c, "ParamsCanNotBeNull", "")
				return
			}
			// 拦截已认证完客户，不能修改认证信息
			if creditData.IsFinishCredit == 1 {
				echo.Error(c, "InformationAlreadyExists", "")
				return
			}
			// 邮箱格式校验
			if !email.IsEmailPattern(params.Account) {
				echo.Error(c, "ErrorEmail", "")
				return
			}
			credit.UpdateEmail(c, params, tAppUser)
		}
		break
	case SubmitTypeCode.Info(SubmitTypeCode.IdentityCredit).Code:
		{
			// 拦截已认证完客户，不能修改认证信息
			if creditData.IsFinishCredit == 1 {
				echo.Error(c, "InformationAlreadyExists", "")
				return
			}
			if len(params.IdBackUrl) == 0 {
				echo.Error(c, "ParamsCanNotBeNull", "")
				return
			}
			if len(params.IdFrontUrl) == 0 {
				echo.Error(c, "ParamsCanNotBeNull", "")
				return
			}
			if len(params.IdentityNo) == 0 {
				echo.Error(c, "ParamsCanNotBeNull", "")
				return
			}
			if len(params.UserName) == 0 {
				echo.Error(c, "ParamsCanNotBeNull", "")
				return
			}
			credit.UpdateIdentity(c, params, tAppUser)
		}
		break
	case SubmitTypeCode.Info(SubmitTypeCode.BankCredit).Code:
		{
			if len(params.BankCardNo) == 0 {
				echo.Error(c, "ParamsCanNotBeNull", "")
				return
			}
			if len(params.BankHolder) == 0 {
				echo.Error(c, "ParamsCanNotBeNull", "")
				return
			}
			if params.BankId == 0 && strings.EqualFold("Indonesia", params.Nation) {
				echo.Error(c, "ParamsCanNotBeNull", "")
				return
			}
			if len(params.BankName) == 0 && strings.EqualFold("India", params.Nation) {
				echo.Error(c, "ParamsCanNotBeNull", "")
				return
			}
			if len(params.BankCode) == 0 && strings.EqualFold("India", params.Nation) {
				echo.Error(c, "ParamsCanNotBeNull", "")
				return
			}

			var tipartList []models.TTipartiteConfiguration
			where := cmap.New().Items()
			where["item_code"] = params.ItemCode
			where["nation"] = params.Nation
			where["is_valid"] = 1
			where["use_type"] = "withdraw"
			mysql.DB.Debug().Model(models.TTipartiteConfiguration{}).Where(where).Find(&tipartList)
			tipart := tipartList[0]
			configBank := new(models.TConfigBankInfo)
			if params.BankId == 0 {
				configBank.BankCode = params.BankCode
				configBank.BankName = params.BankName
				configBank.BankId = params.BankCode
				configBank.BindPayChannel = tipart.PartnerCode
			} else {
				// 新版本绑卡规则获取银行编码规则
				var tSystemBank *models.TSystemBank
				mysql.DB.Debug().Model(models.TSystemBank{}).Where("id", params.BankId).Find(&tSystemBank)
				if tSystemBank == nil {
					logger.Error(errors.New("新版本绑卡规则获取银行编码规则-失败未找到银行列表数据"))
					echo.Error(c, "ParamsCanNotBeNull", "")
					return
				}
				var configBankTem *models.TConfigBankInfo
				mysql.DB.Debug().Model(models.TConfigBankInfo{}).Where(map[string]interface{}{
					"system_bank_id":   tSystemBank.ID,
					"bind_pay_channel": tipart.PartnerCode,
				}).Find(&configBankTem)
				if configBankTem == nil {
					logger.Error(errors.New("新版本绑卡规则获取银行编码规则-失败未找到银行列表数据"))
					echo.Error(c, "ParamsCanNotBeNull", "")
					return
				}
			}
			err, codeName := credit.UpdateBank(c, params, tAppUser, configBank)
			if err != nil {
				echo.Error(c, codeName, "")
				return
			}
		}
		break
	case SubmitTypeCode.Info(SubmitTypeCode.InvitCredit).Code:
		{
			if len(params.InvitCode) == 0 {
				echo.Error(c, "ParamsCanNotBeNull", "")
				return
			}
			credit.SubmitInvit(params, tAppUser)
		}
		break
	default:
		break
	}
	echo.Success(c, nil, "ok")
}

// GetUserCreditMessageHandler
// @Summary No.09.2-获取用户认证状态和认证信息
// @Description 获取用户认证状态和认证信息的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 9.认证模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param Params body requests.Base true "获取用户认证状态和认证信息的参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.UserCreditMessage "data数据结构体"
// @Router /No.09.2/app/v3/credit/getUserCreditMessage.json [post]
func (h *CreditController) GetUserCreditMessageHandler(c *gin.Context) {
	var params requests.Base
	var result response.UserCreditMessage
	var credit *models.AppUserCredit
	var creditMessage response.CreditMessage
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	appUser, _ := c.Get("user")
	appUserId := appUser.(response.ClientUserBeans).AppUserId
	tAppUser := user.GetUser(appUserId)
	if tAppUser != nil {
		result.RejectReason = tAppUser.RejectReason
		creditMessage = response.CreditMessage{
			CreditMobile: tAppUser.CreditMobile,
			CreditEmail:  tAppUser.Email,
			IdentityNo:   tAppUser.IdentityNo,
			IdBackUrl:    tAppUser.IdBackUrl,
			IdFrontUrl:   tAppUser.IdFrontUrl,
			Gender:       int(tAppUser.Gender),
			UserName:     tAppUser.UserName,
			ParentCode:   tAppUser.ParentCode,
			ReferralCode: tAppUser.ReferralCode,
		}
	}
	mysql.DB.Debug().Model(models.AppUserCredit{}).Where("app_user_id", tAppUser.AppUserId).Find(&credit)
	if credit != nil {
		result.IdentityCredit = credit.IdentityCredit
		result.IsFinishCredit = credit.IsFinishCredit
		result.MobileCredit = credit.MobileCredit
		result.EmailCredit = credit.EmailCredit
		result.BankCredit = credit.BankCredit
	}
	rule, _ := redis.Client.WebGet(fmt.Sprintf("%v%v%v%v", "IQOPTION", params.Nation, params.ItemCode, "USEING_CREDIT_RULE"))
	if len(rule) > 0 {
		ruleJson := map[string]interface{}{}
		err := json.Unmarshal([]byte(rule), &ruleJson)
		if err != nil {
			echo.Error(c, "ParamsCanNotBeNull", "")
			return
		}
		result.HadOcr = ruleJson["hadOcr"].(int)
		result.IdentityCreditMSP = ruleJson["identityCreditMSP"].([]string)
		result.BankCreditMSP = ruleJson["bankCreditMSP"].([]string)
	} else {
		result.HadOcr = 0
		result.IdentityCreditMSP = []string{}
		result.BankCreditMSP = []string{}
	}
	result.CreditMessage = creditMessage
	withdrawType := commonUtils.GetWithdrawType(params.Nation, params.ItemCode)
	result.WithdrawType = withdrawType // 当前国家支持的提现类型 0 暂无支持的提现渠道 1 银行卡和数字货币都支持  2 数字货币提现  3 银行卡提现
	echo.Success(c, result, "ok")
}
