package wsdata

type Info struct {
	Symbol       string `json:"symbol"`
	Desc         string `json:"desc"`
	ContractCode string `json:"contract_code"`
}

type WSSubData struct {
	Operation string `json:"operation"`
	Id        string `json:"id"`
	Param     string `json:"param"`
	DataType  string `json:"data_type,omitempty"`
}

//type SubPriceKLineResponse struct {
//	Symbol     string  `json:"symbol"`
//	TimeId     string  `json:"timeId"`
//	Ask        float64 `json:"ask"`
//	Bid        float64 `json:"bid"`
//	Strike     float64 `json:"strike"`
//	AskVolume  string  `json:"askVolume"`
//	BidVolume  string  `json:"bidVolume"`
//	MarketTime string  `json:"marketTime"`
//}

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
