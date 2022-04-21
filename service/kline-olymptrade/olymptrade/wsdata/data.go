package wsdata

type Info struct {
	Symbol       string `json:"symbol"`
	Desc         string `json:"desc"`
	ContractCode string `json:"contract_code"`
}

type SubPriceKLineResponse []struct {
	D []struct {
		P     string  `json:"p"`
		T     float64 `json:"t"`
		Q     float64 `json:"q"`
		Open  float64 `json:"open"`
		Low   float64 `json:"low"`
		High  float64 `json:"high"`
		Close float64 `json:"close"`
	} `json:"d"`
	E int64 `json:"e"`
}
