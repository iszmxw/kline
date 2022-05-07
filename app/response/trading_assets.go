package response

import (
	"goapi/pkg/helpers"
)

// 资产表

type TradingAssets struct {
	Id            int64              `json:"id"`              //资产表主键ID
	ItemCode      string             `json:"item_code"`       //应用编码
	AssetsName    string             `json:"assets_name"`     //资产名称
	AssetsCode    string             `json:"assets_code"`     //资产编码
	IsValid       int                `json:"is_valid"`        //是否有效 1是 0否
	IsShow        int                `json:"is_show"`         //是否展示 1是 0否 删除标识符
	AssetsLogoUrl string             `json:"assets_logo_url"` //资产logo链接
	Sort          int                `json:"sort"`            //排序
	CreateDate    helpers.TimeNormal `json:"create_date"`     //创建日期
	UpdateDate    helpers.TimeNormal `json:"update_date"`     //更新日期
	Creator       string             `json:"creator"`         //创建人
	Updater       string             `json:"updater"`         //更新人
	ReturnRate    float64            `json:"return_rate"`     //收益率
}
