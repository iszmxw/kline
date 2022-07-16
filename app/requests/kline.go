package requests

type Sub struct {
	Sub string `json:"sub"`
	Id  string `json:"id"`
}

type SubHistory struct {
	Symbol    string `json:"symbol"`
	BeginDate string `json:"beginDate"`
	EndDate   string `json:"endDate"`
	Limit     string `json:"limit"`
}

type SubUserInfo struct {
	AppUserId string `json:"appUserId"`
}

type TradeOrderInfo struct {
	AppUserId           string `json:"appUserId"`
	AssetsCode          string `json:"assetsCode"`
	TradeBeginTimestamp string `json:"tradeBeginTimestamp"`
	BeginAssetsPrice    string `json:"beginAssetsPrice"`
	EndAssetsPrice      string `json:"endAssetsPrice"`
	UpDownType          string `json:"upDownType"`
	Amounts             string `json:"amounts"`
	IsDemo              string `json:"isDemo"`
}

type DayTime struct {
	Num string `json:"num" form:"num" validate:"required"`
}
