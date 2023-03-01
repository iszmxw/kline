package models

import "time"

// 用户余额变化表记录ID

type AppUserBalanceChangeLog struct {
	Id            int64     `json:"id"`             //时间段统计表记录ID
	AppUserId     string    `json:"app_user_id"`    //客户ID
	ChangeType    int       `json:"change_type"`    //变更类型 1充值到账 2交易扣除 3交易获胜 4提现扣除 5提现失败返还  6bonus交易平返还  7bonus交易下单 8balance平返还 9手动修改用户余额
	ChangeBalance float64   `json:"change_balance"` //变动金额
	CreateDate    time.Time `json:"create_date"`    //创建日期
	UpdateDate    time.Time `json:"update_date"`    //更新日期
	Nation        string    `json:"nation"`         //国家
	ItemCode      string    `json:"item_code"`      //应用编码
	CurrencyType  string    `json:"currency_type"`  //货币类型 IDR USD CNY
	SourceIndex   string    `json:"source_index"`   //来源唯一键
	TimerBalance  float64   `json:"timer_balance"`  //操作时余额
}
