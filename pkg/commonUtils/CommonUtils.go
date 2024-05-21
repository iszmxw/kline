package commonUtils

import (
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/pkg/mysql"
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
