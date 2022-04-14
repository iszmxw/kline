package raceoption

type Info struct {
	Symbol  string `json:"symbol"`
	Desc    string `json:"desc"`
	Referer string `json:"Referer"`
}

type Response struct {
	Symbol     string  `json:"symbol"`
	TimeId     string  `json:"timeId"`
	Ask        float64 `json:"ask"`
	Bid        float64 `json:"bid"`
	Strike     float64 `json:"strike"`
	AskVolume  string  `json:"askVolume"`
	BidVolume  string  `json:"bidVolume"`
	MarketTime string  `json:"marketTime"`
}
