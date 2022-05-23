package requests

type SetKline struct {
	AssetsCode  string `json:"assets_code" form:"assets_code" validate:"required"`
	KlineSource string `json:"kline_source" form:"kline_source" validate:"required"`
}

type KlineSourceList struct {
	AssetsCode string `json:"assets_code" form:"assets_code"`
}
