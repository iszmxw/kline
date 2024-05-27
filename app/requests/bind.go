package requests

// BindTokenParams 结构体
type BindTokenParams struct {
	Base
	PushToken    string `json:"pushToken"`    // push令牌
	PushPlatForm string `json:"pushPlatForm"` // push平台 firebase
	SubmitType   int    `json:"submitType"`   // 类型 1安装 2注册
	DeviceBrand  string `json:"deviceBrand"`  // 设备品牌
	DeviceId     string `json:"deviceId"`     // 设备ID
	AppUserId    string `json:"app_user_id"`  // 用户ID
}

// UnbindTokenParams 结构体
type UnbindTokenParams struct {
	Base
	PushToken    string `json:"pushToken"`    // push令牌
	PushPlatForm string `json:"pushPlatForm"` // push平台 firebase
	DeviceBrand  string `json:"deviceBrand"`  // 设备品牌
	DeviceId     string `json:"deviceId"`     // 设备ID
	AppUserId    string `json:"app_user_id"`  // 用户ID
}

// BindEmailParams 结构体
type BindEmailParams struct {
	Base
	Email     string `json:"email" validate:"required"` // 邮箱
	AppUserId string `json:"app_user_id"`               // 用户ID
}
