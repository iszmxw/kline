package models

import "time"

type TAppUserAdjustCallbackData struct {
	ID                   int64     `json:"id"`                    // adjust回传数据主键自增id
	AppUserID            string    `json:"app_user_id"`           // 关联客户id
	ItemCode             string    `json:"item_code"`             // 项目编码
	IsOrganic            string    `json:"is_organic"`            // 是否自然客 1是 0否
	EventName            string    `json:"event_name"`            // 事件名称
	GpsAdid              string    `json:"gps_adid"`              // Google广告ID
	CreateDate           time.Time `json:"create_date"`           // 创建时间
	CreateTs             int64     `json:"create_ts"`             // 创建时间戳
	UpdateDate           time.Time `json:"update_date"`           // 更新时间
	AppName              string    `json:"app_name"`              // 应用名称
	Network              string    `json:"network"`               // 网络名称
	Campaign             string    `json:"campaign"`              // 活动名称
	Adid                 string    `json:"adid"`                  // Adjust 设备 ID
	ActivityKind         string    `json:"activity_kind"`         // 活动类型
	DataJson             string    `json:"data_json"`             // 扩展信息JSON
	Remarks              string    `json:"remarks"`               // 备注
	OptimizationDivision string    `json:"optimization_division"` // 优化师
	TrackChannel         string    `json:"track_channel"`         // 投放方式
	AppVersion           string    `json:"app_version"`           // 注册时的app版本号
	Country              string    `json:"country"`               // 国家
	Tracker              string    `json:"tracker"`               // 投放渠道跟踪用户ID
}

func (m *TAppUserAdjustCallbackData) TableName() string {
	return "t_app_user_adjust_callback_data"
}
