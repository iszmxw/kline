package wsdata

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

type ITicker interface {
	Ticker(exchange string, pair string, value TickerValue)
}

type TickerValue struct {
	High   float64
	Low    float64
	Volume float64
	Last   float64
	Time   string
	Period string // 合约周期
}

// 数据流

type DataStream struct {
	Updated int64     `json:"Updated"`
	Rates   []float64 `json:"Rates"` // [Bid,Ask,High,Low]
	Symbol  string    `json:"Symbol"`
}
