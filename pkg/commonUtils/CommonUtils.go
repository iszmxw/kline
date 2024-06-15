package commonUtils

import (
	"encoding/json"
	"errors"
	"fmt"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/app/service/login"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
	"strings"
)

func GeTCollaboratorParams(itemCode, partnerCode, accountCode, nation string) (TCollaboratorParams models.TCollaboratorParams) {
	where := cmap.New().Items()
	where["item_code"] = itemCode
	where["partner_code"] = partnerCode
	where["is_start"] = 1
	if len(accountCode) > 0 {
		where["account_code"] = accountCode
	}
	if len(nation) > 0 {
		where["nation"] = nation
	}
	mysql.DB.Debug().Model(models.TCollaboratorParams{}).Where(where).Find(&TCollaboratorParams)
	return
}

// 提交对象转化成JSON对象，保存到OSS文件服务器

func SendStrToFileServiceGetUrl(str string) (submitParamUrl string) {
	//getFileUrl := "https://fileupload.ssmartoption.com/file/uploadBase64Str"

	return
}

func GetEmailTemplate(nation, languageCode, name string) (arr map[string]string) {
	arr = make(map[string]string)
	arr["status"] = "1"
	// 获取验证码模板
	var TSystemEmailTemplate models.TSystemEmailTemplate
	where := cmap.New().Items()
	where["language_code"] = strings.ToLower(languageCode)
	where["name"] = name
	mysql.DB.Debug().Model(models.TSystemEmailTemplate{}).Where(where).Find(&TSystemEmailTemplate)
	if TSystemEmailTemplate.ID == 0 {
		logger.Error(errors.New("没有路模板可用"))
		arr["status"] = "0"
		return
	}
	appItem := login.GetAppItem(nation)
	if len(appItem.ItemCode) == 0 {
		arr["status"] = "0"
		return
	}
	arr["url"] = appItem.GoMail
	arr["type"] = TSystemEmailTemplate.Type
	arr["title"] = TSystemEmailTemplate.Title
	arr["body"] = TSystemEmailTemplate.Body
	return
}

func GetVipInfo(itemCode, nation string, vipLevel int) *models.TSettingVipRule {
	var resp *models.TSettingVipRule
	if vipLevel == -1 {
		return nil
	}
	key := fmt.Sprintf("config:setting_vip_rule:%v%v", vipLevel, nation)
	get, _ := redis.Client.Get(key)
	if len(get) > 0 {
		err := json.Unmarshal([]byte(get), &resp)
		if err == nil {
			return resp
		}
	} else {
		err := mysql.DB.Model(models.TSettingVipRule{}).Where(map[string]interface{}{
			"item_code": itemCode,
			"vip_level": vipLevel,
			"is_valid":  1,
			"nation":    nation,
		}).Find(&resp).Error
		if err != nil {
			return nil
		}
		marshal, err := json.Marshal(resp)
		if err != nil {
			return nil
		}
		_, _ = redis.Client.Add(key, string(marshal), 0)
	}
	return resp

}
