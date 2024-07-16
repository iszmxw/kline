package response

type GetCashFlow struct {
	List        []CashFlow `json:"list"`
	Total       int64      `json:"total"`
	HasNextPage bool       `json:"hasNextPage"`
	TotalPage   int        `json:"totalPage"`
}

// CashFlow 表示现金流信息
type CashFlow struct {
	OrderId             string  `json:"orderId"`             // 订单ID
	Amount              float64 `json:"amount"`              // 金额
	ShowAmount          string  `json:"showAmount"`          // 格式化后的展示金额
	Status              string  `json:"status"`              // 状态
	PaymentChannel      string  `json:"paymentChannel"`      // 渠道
	SubmitTime          int64   `json:"submitTime"`          // 提交时间
	ArrivalTime         int64   `json:"arrivalTime"`         // 到账时间
	OrderType           string  `json:"orderType"`           // 订单类型
	FailureReason       string  `json:"failureReason"`       // 拒绝理由
	CurrencyType        string  `json:"currencyType"`        // 货币符号
	PhotoPaymentVoucher string  `json:"photoPaymentVoucher"` // 照片支付凭证
	IsUpload            string  `json:"isUpload"`            // 是否需要上传 1-是 0-否
}

// OrderDetail 订单详情
type OrderDetail struct {
	Status       int64   `json:"status"`       // 状态
	PayCurrency  string  `json:"payCurrency"`  // 支付类型
	Account      string  `json:"account"`      // 账户邮箱或者手机号
	Address      string  `json:"address"`      // 地址
	Amount       float64 `json:"amount"`       // 交易金额
	CurrencyType string  `json:"currencyType"` // 货币符号
	AmountPay    float64 `json:"amountPay"`    // 交易金额
}
