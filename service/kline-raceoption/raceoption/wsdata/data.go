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

type SubPriceKLineResponse struct {
	Symbol     string  `json:"symbol"`
	TimeId     string  `json:"timeId"`
	Ask        float64 `json:"ask"`
	Bid        float64 `json:"bid"`
	Strike     float64 `json:"strike"`
	AskVolume  string  `json:"askVolume"`
	BidVolume  string  `json:"bidVolume"`
	MarketTime string  `json:"marketTime"`
}
