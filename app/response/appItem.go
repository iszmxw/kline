package response

type ItemNationMessageParams struct {
	ItemCode         string  `json:"itemCode"`
	LanguageCode     string  `json:"languageCode"`
	ItemNationalName string  `json:"itemNationalName"`
	ItemNationalFlag string  `json:"itemNationalFlag"`
	Nation           string  `json:"nation"`
	IsStart          string  `json:"isStart"`
	TimeZone         string  `json:"timeZone"`
	DemoBalance      float64 `json:"DemoBalance"`
	CurrencyType     string  `json:"CurrencyType"`
	MonetaryUnit     string  `json:"monetaryUnit"`
}
