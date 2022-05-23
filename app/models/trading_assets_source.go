package models

import (
	"goapi/pkg/config"
	"time"
)

// 数据源表，一个数据源对应多个资产，不同的数据源有不同的休市时间

type TradingAssetsSource struct {
	Id              int       `json:"id"`               //数据源表主键ID
	ItemCode        string    `json:"item_code"`        //应用编码
	AssetsCode      string    `json:"assets_code"`      //资产编码
	KlineSource     string    `json:"kline_source"`     //k线来源,数据源
	SubscribeSymbol string    `json:"subscribe_symbol"` //数据源对应的订阅符号
	SleepTime       string    `json:"sleep_time"`       //休市时间 {"start":"6-04:00","ts_long":"172800"} start开始时间，从周几开始几点开始，休息多少秒
	IsValid         int       `json:"is_valid"`         //是否有效 1是 0否
	Remark          string    `json:"remark"`           //数据源备注
	Creator         string    `json:"creator"`          //创建人
	Updater         string    `json:"updater"`          //更新人
	CreateDate      time.Time `json:"create_date"`      //创建日期
	UpdateDate      time.Time `json:"update_date"`      //更新日期
}

func (m *TradingAssetsSource) TableName() string {
	prefix := config.GetString("database.mysql.prefix")
	return prefix + "trading_assets_source"
}
