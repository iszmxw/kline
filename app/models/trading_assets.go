package models

import (
	"time"
)

// 资产表

type TradingAssets struct {
	Id                      int64     `json:"id" gorm:"column:id; PRIMARY_KEY"`                                //资产表主键ID
	ItemCode                string    `json:"item_code" gorm:"column:item_code"`                               // 应用编码
	AssetsName              string    `json:"assets_name" gorm:"column:assets_name"`                           // 资产名称
	AssetsCode              string    `json:"assets_code" gorm:"column:assets_code"`                           // 资产编码
	AssetsStatus            int       `json:"assets_status" gorm:"column:assets_status"`                       // 资产状态 1正常 2休市 3下架
	PointNumber             int       `json:"point_number" gorm:"column:point_number"`                         // 保留小数点位数
	IsValid                 int       `json:"is_valid" gorm:"column:is_valid"`                                 // 是否有效 1是 0否
	IsShow                  int       `json:"is_show" gorm:"column:is_show"`                                   // 是否展示 1是 0否 删除标识符
	AssetsLogoUrl           string    `json:"assets_logo_url" gorm:"column:assets_logo_url"`                   // 资产logo链接
	Sort                    int       `json:"sort" gorm:"column:sort"`                                         // 排序
	CreateDate              time.Time `json:"create_date" gorm:"column:create_date"`                           // 创建日期
	UpdateDate              time.Time `json:"update_date" gorm:"column:update_date"`                           // 更新日期
	Creator                 string    `json:"creator" gorm:"column:creator"`                                   // 创建人
	Updater                 string    `json:"updater" gorm:"column:updater"`                                   // 更新人
	ReturnRate              float64   `json:"return_rate" gorm:"column:return_rate"`                           // 收益率
	FixedRate               float64   `json:"fixed_rate" gorm:"column:fixed_rate"`                             // 固定利率
	IsFloatRate             int       `json:"is_float_rate" gorm:"column:is_float_rate"`                       // 是否开启浮动利率 0-否 1-是
	SleepTime               string    `json:"sleep_time" gorm:"column:sleep_time"`                             // 休市时间 {start:6-04:00,ts_long:172800} start开始时间，从周几开始几点开始，休息多少秒
	AssetsTypeID            int64     `json:"assets_type_id" gorm:"column:assets_type_id"`                     // 资产类型ID
	AssetsTypeCode          string    `json:"assets_type_code" gorm:"column:assets_type_code"`                 // 资产类型编码
	AssetsTypeName          string    `json:"assets_type_name" gorm:"column:assets_type_name"`                 // 资产类型名称
	PercentageBase          int       `json:"percentage_base" gorm:"column:percentage_base"`                   // 百单赢率基数
	AssetThreshold          float64   `json:"asset_threshold" gorm:"column:asset_threshold"`                   // 资产浮动百单盈利停止阈值
	IsMonitoring            int       `json:"is_monitoring" gorm:"column:is_monitoring"`                       // 是否有效1-是 0-否
	AreaTimer               int       `json:"area_timer" gorm:"column:area_timer"`                             // 设置区域时长
	IntervalTimer           int       `json:"interval_timer" gorm:"column:interval_timer"`                     // 设置间隔触发时间
	InadvanceTimer          int       `json:"inadvance_timer" gorm:"column:inadvance_timer"`                   // 设置提前触发时间（拿这个时间间隔作为判断条件）
	LossThreshold           float64   `json:"loss_threshold" gorm:"column:loss_threshold"`                     // 触发属性：区域平台亏损金额>=xx美金
	EarlyexecutionTimer     int       `json:"earlyexecution_timer" gorm:"column:earlyexecution_timer"`         // 设置插针提前执行时间
	EarlyexecutionFrequency float64   `json:"earlyexecution_frequency" gorm:"column:earlyexecution_frequency"` // 设置插针提前执行频率
	RecoverTimer            int       `json:"recover_timer" gorm:"column:recover_timer"`                       // 设置插针恢复时间
	RecoverFrequency        float64   `json:"recover_frequency" gorm:"column:recover_frequency"`               // 设置插针恢复频率
	ShockAmplitude          float64   `json:"shock_amplitude" gorm:"column:shock_amplitude"`                   // 震荡幅度
	VolatilityAmplitude     float64   `json:"volatility_amplitude" gorm:"column:volatility_amplitude"`         // 波动幅度
	VolatilityFrequency     float64   `json:"volatility_frequency" gorm:"column:volatility_frequency"`         // 波动频率
}

type SleepTime struct {
	Start  string `json:"start"`
	TsLong string `json:"ts_long"`
}

type SetUpTime struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type WeekDay struct {
	TradingTime []SetUpTime `json:"TradingTime"`
	SleepTime   []SetUpTime `json:"SleepTime"`
}

type WeekdayTimeDetail struct {
	Name        string      `json:"name"`
	TradingTime []SetUpTime `json:"TradingTime"`
	SleepTime   []SetUpTime `json:"SleepTime"`
}
