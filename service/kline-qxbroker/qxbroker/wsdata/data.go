package wsdata

type SubPriceKLineResponse struct {
	Symbol string  `json:"symbol"`
	TimeId float64 `json:"timeId"`
	Close  float64 `json:"close"`
	Ping   float64 `json:"ping"`
}
