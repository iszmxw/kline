package requests

type CashFlowParams struct {
	Base
	PageInfo
	Today     string `json:"today"`
	OrderType string `json:"orderType"`
}

type CreatePaidCodeParams struct {
	Base
	PageInfo
	Amount       float64 `json:"amounts" validate:"required"`   // 充值金额
	ChannelId    int     `json:"channelId" validate:"required"` // 充值渠道ID
	Mobile       int     `json:"mobile"`                        // 客户手机号
	LastStayPage string  `json:"lastStayPage"`                  // 最后停留的页面
}

type SubmitWithdrawParams struct {
	Base
	PageInfo
	UserBankId int     `json:"userBankId" validate:"required"` // 提现银行ID
	Amount     float64 `json:"amounts" validate:"required"`    // 提现金额
	Type       string  `json:"type"`                           // type 1-银行卡提现 2-数字货币提现
	Channel    string  `json:"channel"`                        // channel
	Address    string  `json:"address"`                        // address
}

type OrderDetailParams struct {
	Base
	PageInfo
	OrderId string `json:"orderId" validate:"required"` // 订ID
}
