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
	Gzip      string `json:"gzip"`
}

type SubUserInfo struct {
	AppUserId string `json:"appUserId"`
}

type SubStationLetter struct {
	Sub SubToken `json:"sub"`
	Id  string   `json:"id"`
}

type SubToken struct {
	Token       string `json:"token"`
	DeciveBrand string `json:"decive_brand" mapstructure:"decive_brand"`
}

type SubNationInfo struct {
	ItemCode string `json:"itemCode"`
	Nation   string `json:"nation"`
}

type SubAdjustCallBack struct {
	Sub AdjustCallBack `json:"sub"`
	Id  string         `json:"id"`
}

type AdjustCallBack struct {
	Adid string `json:"adid"`
}

type DayTime struct {
	Num string `json:"num" form:"num" validate:"required"`
}

type SubParams struct {
	Sub interface{} `json:"sub"`
	Id  string      `json:"id"`
}

type DelayTimestamp struct {
	DelayTimestamp   int `json:"delayTimestamp"`   // 延迟时间戳
	SendTimestamp    int `json:"sendTimestamp"`    // 发生时间戳
	ReceiveTimestamp int `json:"receiveTimestamp"` // 接收时间
}

type DelayInfo struct {
	CreateTimestamp  int64  `bson:"createTimestamp" json:"createTimestamp"`
	DelayTimestamp   int    `bson:"delayTimestamp" json:"delayTimestamp"`
	SendTimestamp    int    `bson:"sendTimestamp" json:"sendTimestamp"`       // 发生时间戳
	ReceiveTimestamp int    `bson:"receiveTimestamp" json:"receiveTimestamp"` // 接收时间
	ItemCode         string `bson:"itemCode" json:"itemCode"`
	AppUserId        string `bson:"appUserId" json:"appUserId"`
	Ip               string `bson:"ip" json:"ip"`
	IpInfo           string `bson:"ipInfo" json:"ipInfo"`
	Url              string `bson:"url" json:"url"`
	Origin           string `bson:"origin" json:"origin"`
	SecWebSocketKey  string `bson:"secWebSocketKey" json:"secWebSocketKey"`
	UserAgent        string `bson:"userAgent" json:"userAgent"`
	Header           string `bson:"header" json:"header"`
}
