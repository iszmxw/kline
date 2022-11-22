package models

import "time"


type AppCurrencyRate struct {
	Id         int64     `json:"id"`
	CurrencyA  string    `json:"currency_a"`  //货币类型 IDR USD CNY
	CurrencyB  string    `json:"currency_b"`  //货币类型 IDR USD CNY
	ANum       float64   `json:"a_num"`       //货币a数量
	BNum       float64   `json:"b_num"`       //货币b数量
	Rate       float64   `json:"rate"`        //汇率转换手续费
	IsValid    int       `json:"is_valid"`    //是否有效1-生效 0-未生效
	CreateDate time.Time `json:"create_date"` //创建时间
	UpdateDate time.Time `json:"update_date"` //更新时间
}
