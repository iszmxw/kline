package response

import (
	"goapi/pkg/helpers"
)

// 资产表

type TradingAssets struct {
	Id            int64              `json:"id"`              //资产表主键ID
	ItemCode      string             `json:"item_code"`       //应用编码
	AssetsName    string             `json:"assets_name"`     //资产名称
	AssetsCode    string             `json:"assets_code"`     //资产编码
	IsValid       int                `json:"is_valid"`        //是否有效 1是 0否
	IsShow        int                `json:"is_show"`         //是否展示 1是 0否 删除标识符
	AssetsLogoUrl string             `json:"assets_logo_url"` //资产logo链接
	Sort          int                `json:"sort"`            //排序
	CreateDate    helpers.TimeNormal `json:"create_date"`     //创建日期
	UpdateDate    helpers.TimeNormal `json:"update_date"`     //更新日期
	Creator       string             `json:"creator"`         //创建人
	Updater       string             `json:"updater"`         //更新人
	ReturnRate    float64            `json:"return_rate"`     //收益率
}

type TradingAssetsMonitoring struct {
	ID                      int64   `json:"id" gorm:"column:id"`                                             // 资产表主键ID
	AssetsCode              string  `json:"assets_code" gorm:"column:assets_code"`                           // 资产编码
	PointNumber             int     `json:"point_number" gorm:"column:point_number"`                         // 保留小数点位数
	IsMonitoring            int     `json:"is_monitoring" gorm:"column:is_monitoring"`                       // 是否有效1-是 0-否
	AreaTimer               int     `json:"area_timer" gorm:"column:area_timer"`                             // 设置区域时长
	IntervalTimer           int     `json:"interval_timer" gorm:"column:interval_timer"`                     // 设置间隔触发时间
	InadvanceTimer          int64   `json:"inadvance_timer" gorm:"column:inadvance_timer"`                   // 设置提前触发时间（拿这个时间间隔作为判断条件）
	LossThreshold           float64 `json:"loss_threshold" gorm:"column:loss_threshold"`                     // 触发属性：区域平台亏损金额>=xx美金
	EarlyexecutionTimer     int64   `json:"earlyexecution_timer" gorm:"column:earlyexecution_timer"`         // 设置插针提前执行时间
	EarlyexecutionFrequency int     `json:"earlyexecution_frequency" gorm:"column:earlyexecution_frequency"` // 设置插针提前执行频率
	RecoverTimer            int64   `json:"recover_timer" gorm:"column:recover_timer"`                       // 设置插针恢复时间
	RecoverFrequency        float64 `json:"recover_frequency" gorm:"column:recover_frequency"`               // 设置插针恢复频率
	ShockAmplitude          float64 `json:"shock_amplitude" gorm:"column:shock_amplitude"`                   // 震荡幅度
	VolatilityAmplitude     float64 `json:"volatility_amplitude" gorm:"column:volatility_amplitude"`         // 波动幅度
	VolatilityFrequency     int64   `json:"volatility_frequency" gorm:"column:volatility_frequency"`         // 震荡频率
}

type AssetsCodeArray struct {
	AssetsCode  string `json:"assets_code"`  //资产编码
	IsWhitelist string `json:"is_whitelist"` //是否白名单1是 0否，白名单的资产不跟随自动变化赔率
}
