package v3

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/app/models/keyenum"
	"goapi/app/requests"
	"goapi/app/response"
	"goapi/pkg/commonUtils"
	"goapi/pkg/echo"
	"goapi/pkg/email"
	"goapi/pkg/firebase"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
	"goapi/pkg/validate"
	"strings"
	"time"
)

// 通知公告服务

type BindController struct {
	BaseController
}

type PushSetRule struct {
	PushPlatForm string `json:"pushPlatForm"`
	TopicName    string `json:"topicName"`
	IsRequest    int    `json:"isRequest"`
	ImageUrl     string `json:"imageUrl"`
}

// BindTokenHandler 绑定推送令牌
func (h *BindController) BindTokenHandler(c *gin.Context) {
	var params requests.BindTokenParams
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	if len(params.PushToken) == 0 {
		// redis 设置了可以不用绑定直接返回成功
		s, _ := redis.Client.Get(keyenum.KeyInfo(keyenum.PUSH_TOKEN_CHECK).Key)
		if strings.EqualFold("1", s) {
			echo.Success(c, nil, "SUCCESS")
			return
		}
	}
	if params.SubmitType == 0 || len(params.PushPlatForm) == 0 {
		echo.Success(c, nil, "SUCCESS")
		return
	}
	var SystemApiRule models.SystemApiRule
	where := cmap.New().Items()
	where["is_start"] = 1
	where["item_code"] = params.ItemCode
	where["rule_code"] = "PUSH_SET_RULE"
	where["language_code"] = params.LanguageCode
	where["nation"] = params.Nation
	mysql.DB.Debug().Model(models.SystemApiRule{}).Where(where).Find(&SystemApiRule)
	if SystemApiRule.Id == 0 {
		echo.Error(c, "Failed", "")
		return
	}
	var ruleJson PushSetRule
	err := json.Unmarshal([]byte(SystemApiRule.RuleJson), &ruleJson)
	if err != nil {
		logger.Error(err)
		echo.Error(c, "Failed", err.Error())
		return
	}
	var TAppPushUserConnect models.TAppPushUserConnect
	where2 := cmap.New().Items()
	if len(params.PushPlatForm) > 0 {
		where2["push_plat_form"] = params.PushPlatForm
	}
	if len(params.AppUserId) > 0 {
		where2["app_user_id"] = params.AppUserId
	}
	if params.SubmitType != 0 {
		where2["submit_type"] = params.SubmitType
	}
	if len(params.PushToken) > 0 {
		where2["token_index"] = strings.ToLower(params.PushToken[:16])
	}
	mysql.DB.Debug().Model(models.TAppPushUserConnect{}).Where(where2).Order("create_data desc").Find(&TAppPushUserConnect)
	if TAppPushUserConnect.PushId == 0 {
		// 创建
		mysql.DB.Debug().Model(models.TAppPushUserConnect{}).Create(&models.TAppPushUserConnect{
			PushBindToken: params.PushToken,
			PushPlatForm:  params.PushPlatForm,
			AppUserId:     params.AppUserId,
			DeciveId:      params.DeviceId,
			DeciveBrand:   params.DeviceBrand,
			CreateDate:    time.Now(),
			UpdateDate:    time.Now(),
			DataJsonUrl:   "",
			ItemCode:      params.ItemCode,
			SubmitType:    int64(params.SubmitType),
			TokenIndex:    strings.ToLower(params.PushToken[:16]),
			Nation:        params.Nation,
		})
	} else {
		// 更新
		mysql.DB.Debug().Model(models.TAppPushUserConnect{}).Where("push_id", TAppPushUserConnect.PushId).Updates(map[string]interface{}{
			"push_bind_token": params.PushToken,
			"update_date":     time.Now(),
			"submit_type":     int64(params.SubmitType),
			"token_index":     strings.ToLower(params.PushToken[:16]),
		})
	}
	mysql.DB.Debug().Model(models.TAppPushUserConnectLog{}).Create(&models.TAppPushUserConnectLog{
		PushBindToken: params.PushToken,
		PushPlatForm:  params.PushPlatForm,
		AppUserId:     params.AppUserId,
		DeciveId:      params.DeviceId,
		DeciveBrand:   params.DeviceBrand,
		CreateDate:    time.Now(),
		SubmitType:    int64(params.SubmitType),
		ConnectType:   1,
		TokenIndex:    strings.ToLower(params.PushToken[:16]),
		ItemCode:      params.ItemCode,
		Nation:        params.Nation,
	})
	if ruleJson.IsRequest == 1 {
		tCollaboratorParams := commonUtils.GeTCollaboratorParams(params.ItemCode, "firebase", "", "")
		if tCollaboratorParams.ID == 0 {
			echo.Error(c, "Failed", "")
			return
		}
		// 直接提供服务账号密钥的 JSON 字符串
		serviceAccountKeyJSON := tCollaboratorParams.ParamJson
		// 要订阅的主题
		topic := ruleJson.TopicName + params.Nation
		// 要订阅的设备令牌
		tokens := []string{
			params.PushToken,
		}
		toTopic := firebase.SubscribeToTopic(serviceAccountKeyJSON, topic, tokens)
		fmt.Printf("Successfully subscribed to topic: %d tokens were subscribed\n", toTopic.SuccessCount)
	}
	echo.Success(c, nil, "SUCCESS")
}

// UnbindTokenHandler 解绑推送令牌
func (h *BindController) UnbindTokenHandler(c *gin.Context) {
	HeaderParams := getHeaderParams(c)
	var params requests.UnbindTokenParams
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	if len(HeaderParams.Token) > 0 {
		redis.Client.Delete(keyenum.KeyInfo(keyenum.CLIENT_TOKEN).Key + HeaderParams.Token)
	}
	if len(params.PushToken) == 0 {
		redis.Client.Delete(keyenum.KeyInfo(keyenum.CLIENT_TOKEN).Key + HeaderParams.Token)
		// redis 设置了可以不用绑定直接返回成功
		s, _ := redis.Client.Get(keyenum.KeyInfo(keyenum.PUSH_TOKEN_CHECK).Key)
		if strings.EqualFold("1", s) {
			echo.Success(c, nil, "SUCCESS")
			return
		}
	}
	if len(params.PushPlatForm) == 0 {
		echo.Success(c, nil, "SUCCESS")
		return
	}

	var SystemApiRule models.SystemApiRule
	where := cmap.New().Items()
	where["is_start"] = 1
	where["item_code"] = params.ItemCode
	where["rule_code"] = "PUSH_SET_RULE"
	where["language_code"] = params.LanguageCode
	where["nation"] = params.Nation
	mysql.DB.Debug().Model(models.SystemApiRule{}).Where(where).Find(&SystemApiRule)
	if SystemApiRule.Id == 0 {
		echo.Error(c, "Failed", "")
		return
	}
	var ruleJson PushSetRule
	err := json.Unmarshal([]byte(SystemApiRule.RuleJson), &ruleJson)
	if err != nil {
		logger.Error(err)
		echo.Error(c, "Failed", err.Error())
		return
	}
	var ListTAppPushUserConnect []models.TAppPushUserConnect
	where2 := cmap.New().Items()
	if len(params.PushPlatForm) > 0 {
		where2["push_plat_form"] = params.PushPlatForm
	}
	if len(params.AppUserId) > 0 {
		where2["app_user_id"] = params.AppUserId
	}
	if len(params.PushToken) > 0 {
		where2["token_index"] = strings.ToLower(params.PushToken[:16])
	}
	mysql.DB.Debug().Model(models.TAppPushUserConnect{}).Where(where2).Order("create_data desc").Find(&ListTAppPushUserConnect)
	for _, item := range ListTAppPushUserConnect {
		// 更新
		mysql.DB.Debug().Model(models.TAppPushUserConnect{}).Where("push_id", item.PushId).Updates(map[string]interface{}{
			"update_date": time.Now(),
			"submit_type": -3,
		})
	}
	mysql.DB.Debug().Model(models.TAppPushUserConnectLog{}).Create(&models.TAppPushUserConnectLog{
		PushBindToken: params.PushToken,
		PushPlatForm:  params.PushPlatForm,
		AppUserId:     params.AppUserId,
		DeciveId:      params.DeviceId,
		DeciveBrand:   params.DeviceBrand,
		CreateDate:    time.Now(),
		SubmitType:    -3,
		ConnectType:   1,
		TokenIndex:    strings.ToLower(params.PushToken[:16]),
		ItemCode:      params.ItemCode,
		Nation:        params.Nation,
	})
	if ruleJson.IsRequest == 1 {
		tCollaboratorParams := commonUtils.GeTCollaboratorParams(params.ItemCode, "firebase", "", "")
		if tCollaboratorParams.ID == 0 {
			echo.Error(c, "Failed", "")
			return
		}
		// 直接提供服务账号密钥的 JSON 字符串
		serviceAccountKeyJSON := tCollaboratorParams.ParamJson
		// 要解绑的主题
		topic := ruleJson.TopicName + params.Nation
		// 要解绑的设备令牌
		tokens := []string{
			params.PushToken,
		}
		fromTopic := firebase.UnsubscribeFromTopic(serviceAccountKeyJSON, topic, tokens)
		fmt.Printf("Successfully unsubscribed to topic: %d tokens were unsubscribed\n", fromTopic.SuccessCount)
	}
	echo.Success(c, nil, "SUCCESS")
}

// BindAccountEmailHandler 绑定用户邮箱
func (h *BindController) BindAccountEmailHandler(c *gin.Context) {
	var params requests.BindEmailParams
	var AppUser models.AppUser
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	if !email.IsEmailPattern(params.Email) {
		echo.Error(c, "ErrorEmail", "")
		return
	}
	user, _ := c.Get("user")
	ClientUserBeans := user.(response.ClientUserBeans)
	appUserId := ClientUserBeans.AppUserId
	if params.AppUserId != appUserId {
		echo.Error(c, "AccountCanNotNull", "")
		return
	}
	// 验证邮箱和KYC邮箱是否已被注册
	mysql.DB.Debug().Model(models.AppUser{}).Where("email", params.Email).Or("credit_email", params.Email).Find(&AppUser)
	if len(AppUser.AppUserId) > 0 {
		echo.Error(c, "EmailIsUsed", "")
		return
	}
	err := mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", ClientUserBeans.AppUserId).Updates(map[string]interface{}{
		"email":       params.Email,
		"update_date": time.Now(),
		"update_ts":   time.Now().UnixMilli(),
	}).Error
	if err != nil {
		echo.Error(c, "ErrorEmail", "")
		return
	}
	echo.Success(c, nil, "SUCCESS")
}
