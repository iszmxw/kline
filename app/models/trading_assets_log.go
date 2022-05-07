package models

import "time"

// 资产表赔率修改记录表

type TradingAssetsLog struct {
	Id                int64     `json:"id" gorm:"column:id; PRIMARY_KEY"` //资产表主键ID
	ItemCode          string    `json:"item_code"`                        //应用编码
	AssetsCode        string    `json:"assets_code"`                      //资产编码
	CreateDate        time.Time `json:"create_date"`                      //创建日期
	ReturnRateHistory float64   `json:"return_rate_history"`              //历史收益率
	ReturnRateNow     float64   `json:"return_rate_now"`                  //新的收益率
	KlineSource       string    `json:"kline_source"`                     //赔率来源，可能是采集的数据源，也可能是自我们自己估算的
}
