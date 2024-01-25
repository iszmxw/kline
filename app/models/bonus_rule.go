package models

type BonusRule struct {
	RealAmounts     float64 `json:"realAmounts"`     //真实交易总额
	BonusAmounts    float64 `json:"bonusAmounts"`    //bonus交易总额
	RechargeAmounts float64 `json:"rechargeAmounts"` //充值金额
}
