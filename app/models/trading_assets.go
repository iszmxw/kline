package models

import "time"

// 资产表

type TradingAssets struct {
	Id            int64     `json:"id" gorm:"column:id; PRIMARY_KEY"` //资产表主键ID
	ItemCode      string    `json:"item_code"`                        //应用编码
	AssetsName    string    `json:"assets_name"`                      //资产名称
	AssetsCode    string    `json:"assets_code"`                      //资产编码
	IsValid       int       `json:"is_valid"`                         //是否有效 1是 0否
	IsShow        int       `json:"is_show"`                          //是否展示 1是 0否 删除标识符
	AssetsLogoUrl string    `json:"assets_logo_url"`                  //资产logo链接
	Sort          int       `json:"sort"`                             //排序
	CreateDate    time.Time `json:"create_date"`                      //创建日期
	UpdateDate    time.Time `json:"update_date"`                      //更新日期
	Creator       string    `json:"creator"`                          //创建人
	Updater       string    `json:"updater"`                          //更新人
	ReturnRate    float64   `json:"return_rate"`                      //收益率
	PointNumber   int       `json:"point_number"`                     //保留小数点位数
	AssetsStatus  int       `json:"assets_status"`                    //资产状态 1正常 2休市 3下架
	SleepTime     string    `json:"sleep_time"`                       //休市时间 {"start":"6-04:00","ts_long":"172800"} start开始时间，从周几开始几点开始，休息多少秒
}

type SleepTime struct {
	Start  string `json:"start"`
	TsLong string `json:"ts_long"`
}
