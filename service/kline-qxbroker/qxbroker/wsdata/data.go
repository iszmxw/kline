package wsdata

type Info struct {
	Symbol       string `json:"symbol"`
	Desc         string `json:"desc"`
	ContractCode string `json:"contract_code"`
}

type SubPriceKLineResponse struct {
	Symbol string  `json:"symbol"`
	TimeId float64 `json:"timeId"`
	Close  float64 `json:"close"`
	Ping   float64 `json:"ping"`
}
