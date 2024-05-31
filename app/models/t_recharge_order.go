package models

import "time"

// TRechargeOrder 充值订单记录表
type TRechargeOrder struct {
	RechargeId          int64     `json:"recharge_id" gorm:"recharge_id"`                     // 充值订单记录表主键ID
	OnlyCode            string    `json:"only_code" gorm:"only_code"`                         // 唯一关联流水号
	CreateDate          time.Time `json:"create_date" gorm:"create_date"`                     // 创建日期
	UpdateDate          time.Time `json:"update_date" gorm:"update_date"`                     // 更新日期
	PayCode             string    `json:"pay_code" gorm:"pay_code"`                           // 付款连接或付款码
	BindPayChannel      string    `json:"bind_pay_channel" gorm:"bind_pay_channel"`           // 绑定的在线支付渠道 wyp王杨pay insta instaMoney
	AppUserId           string    `json:"app_user_id" gorm:"app_user_id"`                     // 关联客户ID
	ItemCode            string    `json:"item_code" gorm:"item_code"`                         // 关联项目编码
	Status              string    `json:"status" gorm:"status"`                               // 订单状态 01：未支付 02：已支付 05：转入退款 09：支付失败
	RechargeOrderStatus int64     `json:"recharge_order_status" gorm:"recharge_order_status"` // 充值状态 1充值成功 0充值失败 2处理中 3已超时
	OutTradeNo          string    `json:"out_trade_no" gorm:"out_trade_no"`                   // 商户订单号
	TradeType           string    `json:"trade_type" gorm:"trade_type"`                       // 下单固定传 pay.submit
	MchId               string    `json:"mch_id" gorm:"mch_id"`                               // 平台分配给商户的商户号
	Channel             string    `json:"channel" gorm:"channel"`                             // 支付类型
	MobilePhone         string    `json:"mobile_phone" gorm:"mobile_phone"`                   // 手机号
	Amount              float64   `json:"amount" gorm:"amount"`                               // 交易金额  单位为印尼盾，两位小数
	AmountPay           float64   `json:"amount_pay" gorm:"amount_pay"`                       // 根据汇率换算出用户要支付的USDT或者其他货币的数量
	Email               string    `json:"email" gorm:"email"`                                 // 邮箱
	NotifyUrl           string    `json:"notify_url" gorm:"notify_url"`                       // 异步通知地址
	Payer               string    `json:"payer" gorm:"payer"`                                 // 还款人姓名
	SignKey             string    `json:"sign_key" gorm:"sign_key"`                           // 签名秘钥
	ReturnCode          string    `json:"return_code" gorm:"return_code"`                     // 返回状态码  0：成功、1：失败
	ReturnMsg           string    `json:"return_msg" gorm:"return_msg"`                       // 应答描述
	OutChannelNo        string    `json:"out_channel_no" gorm:"out_channel_no"`               // 平台订单号
	PayCodeWay          string    `json:"pay_code_way" gorm:"pay_code_way"`                   // 付款连接类型
	ErrCode             string    `json:"err_code" gorm:"err_code"`                           // 错误代码
	ErrCodeDes          string    `json:"err_code_des" gorm:"err_code_des"`                   // 错误返回的信息描述
	EndValidDate        time.Time `json:"end_valid_date" gorm:"end_valid_date"`               // 截止有效日期
	TransTime           string    `json:"trans_time" gorm:"trans_time"`                       // 请求格式化日期
	CostTime            int64     `json:"cost_time" gorm:"cost_time"`                         // 请求耗时
	LoanId              int64     `json:"loan_id" gorm:"loan_id"`                             // 关联订单ID
	PaycodeCreator      string    `json:"paycode_creator" gorm:"paycode_creator"`             // 生成还款码的操作者
	PaycodeCreatorType  int64     `json:"paycode_creator_type" gorm:"paycode_creator_type"`   // 生成还款码操作者类型 1客户 2后台用户
	IsLockingFirst      int64     `json:"is_locking_first" gorm:"is_locking_first"`           // 是否为第一次锁定订单还款记录 1是 0否
	DataJson            string    `json:"data_json" gorm:"data_json"`                         // 扩展信息JSON
	BarcodeUrl          string    `json:"barcode_url" gorm:"barcode_url"`                     // 便利店的付款条形码
	PayAccountCode      string    `json:"pay_account_code" gorm:"pay_account_code"`           // 还款账号编码
	Fee                 float64   `json:"fee" gorm:"fee"`                                     // 服务费 【支付方收取我们的】
	IsFirstRechage      int64     `json:"is_first_rechage" gorm:"is_first_rechage"`           // 是否首充 1是 0否
	RegistDateRecharge  int64     `json:"regist_date_recharge" gorm:"regist_date_recharge"`   // 是否注册日充值 1是 0否
	CurrencyType        string    `json:"currency_type" gorm:"currency_type"`                 // 货币类型 IDR USD CNY
	Nation              string    `json:"nation" gorm:"nation"`                               // 绑定国家
	CreateTs            int64     `json:"create_ts" gorm:"create_ts"`                         // 创建充值订单时间戳
	UpdateTs            int64     `json:"update_ts" gorm:"update_ts"`                         // 更新充值订单时间戳
	EndValidTs          int64     `json:"end_valid_ts" gorm:"end_valid_ts"`                   // 订单截至有效时间戳
	Photo               string    `json:"photo" gorm:"photo"`                                 // 照片凭证图片
}

// TableName 表名称
func (*TRechargeOrder) TableName() string {
	return "t_recharge_order"
}
