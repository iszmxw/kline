package models

import "time"

type TradingAssetsMonitoring struct {
	ID                  int64     `json:"id" gorm:"column:id"`                                     // 资产表主键ID
	ItemCode            string    `json:"item_code" gorm:"column:item_code"`                       // 应用编码
	AssetsCode          string    `json:"assets_code" gorm:"column:assets_code"`                   // 资产名称
	IsValid             int       `json:"is_valid" gorm:"column:is_valid"`                         // 是否有效1-是 0-否
	AreaTimer           int       `json:"area_timer" gorm:"column:area_timer"`                     // 设置区域时长
	IntervalTimer       int       `json:"interval_timer" gorm:"column:interval_timer"`             // 设置间隔触发时间
	InadvanceTimer      int64     `json:"inadvance_timer" gorm:"column:inadvance_timer"`           // 设置提前触发时间（拿这个时间间隔作为判断条件）
	LossThreshold       float64   `json:"loss_threshold" gorm:"column:loss_threshold"`             // 触发属性：区域平台亏损金额>=xx美金
	ShockAmplitude      float64   `json:"shock_amplitude" gorm:"column:shock_amplitude"`           // 震荡幅度
	VolatilityAmplitude float64   `json:"volatility_amplitude" gorm:"column:volatility_amplitude"` // 波动幅度
	VolatilityFrequency string    `json:"volatility_frequency" gorm:"column:volatility_frequency"` // 波动频率
	Creator             string    `json:"creator" gorm:"column:creator"`                           // 创建人
	Updater             string    `json:"updater" gorm:"column:updater"`                           // 更新人
	CreateDate          time.Time `json:"create_date" gorm:"column:create_date"`                   // 创建日期
	UpdateDate          time.Time `json:"update_date" gorm:"column:update_date"`                   // 更新日期
}

func (m *TradingAssetsMonitoring) TableName() string {
	return "t_trading_assets_monitoring"
}
