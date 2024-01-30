package models

import "time"

type TSettingVipRule struct {
	ID                  int64     `json:"id"`                  // VIP配置表自增ID
	Nation              string    `json:"nation"`              // 国家
	LanguageCode        string    `json:"languageCode"`        // 语言编码
	ItemCode            string    `json:"itemCode"`            // 关联项目编码
	IsValid             int       `json:"isValid"`             // 是否启用 1是 0否
	IsShow              int       `json:"isShow"`              // 是否展示 1展示 0不展示[删除标识符]
	VipLevel            int       `json:"vipLevel"`            // 会员等级
	VipName             string    `json:"vipName"`             // 会员等级前缀
	VipLogoUrl          string    `json:"vipLogoUrl"`          // 会员等级图标URL
	RechargeIcon        string    `json:"rechargeIcon"`        // 充值页面快捷金额对应小图标
	VipBanner           string    `json:"vipBanner"`           // 会员卡片图标
	Color               string    `json:"color"`               // 进度条颜色
	Action              int       `json:"action"`              // 提升特权按钮：1去充值、2去交易
	RechargeStandard    float64   `json:"rechargeStandard"`    // 充值达标金额
	TradingStandard     int       `json:"tradingStandard"`     // 交易达标次数
	UsdRechargeStandard int       `json:"usdRechargeStandard"` // 是否使用充值金额作为标准 1是 0否
	UsdTrdingStandard   int       `json:"usdTrdingStandard"`   // 是否使用交易次数作为标准 1是 0否
	LostReturnRate      float64   `json:"lostReturnRate"`      // 亏损返还比例
	WithdrawTimes       int       `json:"withdrawTimes"`       // 单日提现次数
	ParallelTrans       int64     `json:"parallelTrans"`       // 同时交易订单数
	MaxTransAmount      float64   `json:"maxTransAmount"`      // 最大交易金额
	WithdrawAmounts     int       `json:"withdrawAmounts"`     // 单日提现金额
	WidrawFeeJson       string    `json:"widrawFeeJson"`       // 提现手续费规则
	StandardShowWord    string    `json:"standardShowWord"`    // 会员达标条件
	WelfareShowWord     string    `json:"welfareShowWord"`     // 会员享受福利文案
	WithdrawShowWord    string    `json:"withdrawShowWord"`    // 会员提现页面规则文案
	RechargeShowWord    string    `json:"rechargeShowWord"`    // 充值页面快捷金额文案
	Creator             string    `json:"creator"`             // 创建人
	Updater             string    `json:"updater"`             // 修改人
	CreatTime           time.Time `json:"creatTime"`           // 创建时间
	UpdateTime          time.Time `json:"updateTime"`          // 修改时间
	Remarks             string    `json:"remarks"`             // 备注
}

func (m *TSettingVipRule) TableName() string {
	return "t_setting_vip_rule"
}
