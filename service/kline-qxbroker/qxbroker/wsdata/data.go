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
	Symbol string  `json:"symbol"`
	TimeId float64 `json:"timeId"`
	Close  float64 `json:"close"`
	Ping   float64 `json:"ping"`
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

type Chart struct {
	ChartId  string `json:"chart_id"`
	Settings string `json:"settings"`
}

type Settings struct {
	ChartId               string `json:"chartId"`
	ChartType             int    `json:"chartType"`
	CurrentExpirationTime int    `json:"currentExpirationTime"`
	IsFastOption          bool   `json:"isFastOption"`
	IsFastAmountOption    bool   `json:"isFastAmountOption"`
	IsIndicatorsMinimized bool   `json:"isIndicatorsMinimized"`
	IsIndicatorsShowing   bool   `json:"isIndicatorsShowing"`
	IsShortBetElement     bool   `json:"isShortBetElement"`
	ChartPeriod           int    `json:"chartPeriod"`
	CurrentAsset          struct {
		Symbol string `json:"symbol"`
	} `json:"currentAsset"`
	DealValue        int    `json:"dealValue"`
	DealPercentValue int    `json:"dealPercentValue"`
	IsVisible        bool   `json:"isVisible"`
	TimePeriod       int    `json:"timePeriod"`
	GridOpacity      int    `json:"gridOpacity"`
	IsAutoScrolling  int    `json:"isAutoScrolling"`
	IsOneClickTrade  bool   `json:"isOneClickTrade"`
	UpColor          string `json:"upColor"`
	DownColor        string `json:"downColor"`
}
