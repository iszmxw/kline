package requests

type GetSettingVipParams struct {
	Base
	DeviceType string `json:"deviceType"` // 设备类型
}
