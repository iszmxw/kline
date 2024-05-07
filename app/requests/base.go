package requests

type Base struct {
	ItemCode     string `json:"itemCode" validate:"required"`     // 项目编码
	Nation       string `json:"nation" validate:"required"`       // 国家编码
	LanguageCode string `json:"languageCode" validate:"required"` // 语言编码
	AppOsType    string `json:"appOsType"`                        // 系统类型 1安卓 2IOS 3html
}
