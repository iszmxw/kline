package requests

type Sub struct {
	Sub string `json:"sub"`
	Id  string `json:"id"`
}

type SubHistory struct {
	Symbol    string `json:"symbol"`
	BeginDate string `json:"beginDate"`
	EndDate   string `json:"endDate"`
}

type SubUserInfo struct {
	AppUserId    string `json:"appUserId"`
}

type DayTime struct {
	Num string `json:"num" form:"num" validate:"required"`
}
