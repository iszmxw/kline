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

type Option struct {
	Success bool `json:"success"`
	Data    []struct {
		Symbol        string  `json:"symbol"`
		DisplayName   string  `json:"displayName"`
		PayoutCorr    float64 `json:"payoutCorr"`
		Active        bool    `json:"active"`
		Weekend       bool    `json:"weekend"`
		OptionType    string  `json:"optionType"`
		MinAmount     int     `json:"minAmount"`
		MaxAmount     int     `json:"maxAmount"`
		SequenceNr    string  `json:"sequenceNr"`
		Dynamic       bool    `json:"dynamic"`
		MinPayout     float64 `json:"minPayout"`
		MaxPayout     float64 `json:"maxPayout"`
		Profitability float64 `json:"profitability"`
		Coeff         int     `json:"coeff"`
		DynPayout     float64 `json:"dynPayout"`
		RoundingDigit string  `json:"roundingDigit"`
		MinTradeTime  string  `json:"minTradeTime"`
		AgScanP       string  `json:"agScanP"`
		AgPnl         int     `json:"agPnl"`
		AgWinRate     int     `json:"agWinRate"`
		AgDown        int     `json:"agDown"`
	} `json:"data"`
	Constraints []interface{} `json:"constraints"`
}
