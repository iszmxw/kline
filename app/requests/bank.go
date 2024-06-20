package requests

// SetUserBankCardParams 删除银行卡
type SetUserBankCardParams struct {
	Bank     string `json:"bank" validate:"required"`     // 银行名称
	CardNo   string `json:"cardNo" validate:"required"`   // 银行卡号
	BankCode string `json:"bankCode" validate:"required"` // 银行编码
}

type GetPaidFastAmountsParams struct {
	PaidTypeId int `json:"paidTypeId"` // 充值类型ID
}
