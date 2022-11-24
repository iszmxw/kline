package wsdata

type Info struct {
	Symbol       string `json:"symbol"`
	Desc         string `json:"desc"`
	ContractCode string `json:"contract_code"`
}

type SubPriceKLineResponse struct {
	Id     float64 `json:"id"`
	Symbol string  `json:"symbol"`
	Open   float64 `json:"open"`
	Close  float64 `json:"close"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
}
