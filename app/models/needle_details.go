package models

import "time"

type NeedleDetails struct {
	ID                   int64     `json:"id" gorm:"column:id"`                                           // 资产表主键ID
	AssetsCode           string    `json:"assets_code" gorm:"column:assets_code"`                         // 资产名称
	StartTimer           int64     `json:"start_timer" gorm:"column:start_timer"`                         // 执行插针开始时间
	EndTimer             int64     `json:"end_timer" gorm:"column:end_timer"`                             // 执行插针结束时间
	OrderNum             int       `json:"order_num" gorm:"column:order_num"`                             // 干预订单数
	LossThreshold        float64   `json:"loss_threshold" gorm:"column:loss_threshold"`                   // 干预用户数
	PeopleNum            int       `json:"people_num" gorm:"column:people_num"`                           // 干预用户数
	MarketQuoteOdds      int64     `json:"market_quote_odds" gorm:"column:market_quote_odds"`             // 市场报价胜率
	MarketQuoteWinAmount float64   `json:"market_quote_win_amount" gorm:"column:market_quote_win_amount"` // 市场报价盈亏金额
	NeedleAfterOdds      float64   `json:"needle_after_odds" gorm:"column:needle_after_odds"`             // 插针后胜率
	NeedleAfterWinAmount float64   `json:"needle_after_win_amount" gorm:"column:needle_after_win_amount"` // 插针后胜率
	StartDate            time.Time `json:"start_date" gorm:"column:start_date"`                           // 执行插针开始时间
	EndDate              time.Time `json:"end_date" gorm:"column:end_date"`                               // 执行插针结束时间
	CreateDate           time.Time `json:"create_date" gorm:"column:create_date"`                         // 创建日期
	UpdateDate           time.Time `json:"update_date" gorm:"column:update_date"`                         // 更新日期
}

func (m *NeedleDetails) TableName() string {
	return "t_needle_details"
}
