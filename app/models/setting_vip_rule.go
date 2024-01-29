package models

import "time"

type TSettingVipRule struct {
	ID                  int64     `json:"id"`                    // VIP配置表自增ID
	Nation              string    `json:"nation"`                // 国家
	LanguageCode        string    `json:"language_code"`         // 语言编码
	ItemCode            string    `json:"item_code"`             // 关联项目编码
	IsValid             int       `json:"is_valid"`              // 是否启用 1是 0否
	IsShow              int       `json:"is_show"`               // 是否展示 1展示 0不展示[删除标识符]
	VipLevel            int       `json:"vip_level"`             // 会员等级
	VipName             string    `json:"vip_name"`              // 会员等级前缀
	VipLogoUrl          string    `json:"vip_logo_url"`          // 会员等级图标URL
	RechargeIcon        string    `json:"recharge_icon"`         // 充值页面快捷金额对应小图标
	VipBanner           string    `json:"vip_banner"`            // 会员卡片图标
	Color               string    `json:"color"`                 // 进度条颜色
	Action              int       `json:"action"`                // 提升特权按钮：1去充值、2去交易
	RechargeStandard    float64   `json:"recharge_standard"`     // 充值达标金额
	TradingStandard     int       `json:"trading_standard"`      // 交易达标次数
	UsdRechargeStandard int       `json:"usd_recharge_standard"` // 是否使用充值金额作为标准 1是 0否
	UsdTrdingStandard   int       `json:"usd_trding_standard"`   // 是否使用交易次数作为标准 1是 0否
	LostReturnRate      float64   `json:"lost_return_rate"`      // 亏损返还比例
	WithdrawTimes       int       `json:"withdraw_times"`        // 单日提现次数
	ParallelTrans       int64     `json:"parallel_trans"`        // 同时交易订单数
	MaxTransAmount      float64   `json:"max_trans_amount"`      // 最大交易金额
	WithdrawAmounts     int       `json:"withdraw_amounts"`      // 单日提现金额
	WidrawFeeJson       string    `json:"widraw_fee_json"`       // 提现手续费规则
	StandardShowWord    string    `json:"standard_show_word"`    // 会员达标条件
	WelfareShowWord     string    `json:"welfare_show_word"`     // 会员享受福利文案
	WithdrawShowWord    string    `json:"withdraw_show_word"`    // 会员提现页面规则文案
	RechargeShowWord    string    `json:"recharge_show_word"`    // 充值页面快捷金额文案
	Creator             string    `json:"creator"`               // 创建人
	Updater             string    `json:"updater"`               // 修改人
	CreatTime           time.Time `json:"creat_time"`            // 创建时间
	UpdateTime          time.Time `json:"update_time"`           // 修改时间
	Remarks             string    `json:"remarks"`               // 备注
}

func (m *TSettingVipRule) TableName() string {
	return "t_setting_vip_rule"
}
