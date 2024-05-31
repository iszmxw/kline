package requests

type GetLanguagePageParams struct {
	Base
	Type string `json:"type" validate:"required"`
	Uuid string `json:"uuid"`
	Adid string `json:"adid"`
}

type ExistingCountryParams struct {
	Base
	SystemLanguage string `json:"systemLanguage"` // 用户手机默认系统语言
}

type PageShow struct {
	ShowColumn string `json:"show_column"` // 语言位置 key
	ShowWord   string `json:"show_word"`   // 展示内容
}
