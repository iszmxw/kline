package models

type BonusRule struct {
	RealAmounts     float64 `json:"real_amounts"`     //真实交易总额
	BonusAmounts    float64 `json:"bonus_amounts"`    //bonus交易总额
	RechargeAmounts float64 `json:"recharge_amounts"` //充值金额
}
