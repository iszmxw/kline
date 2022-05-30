package requests

type Sub struct {
	Sub string `json:"sub"`
	Id  string `json:"id"`
}

type DayTime struct {
	Num int `json:"num" form:"num" validate:"required"`
}
