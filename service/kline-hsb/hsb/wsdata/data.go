package wsdata

type Info struct {
	ID           int    `json:"id"`
	Symbol       string `json:"symbol"`
	Desc         string `json:"desc"`
	ContractCode string `json:"contract_code"`
}

type SubPriceKLineResponse struct {
	Cmd float64 `json:"cmd"`
	Pb  struct {
		Tick struct {
			Id       float64 `json:"id"`
			Open     float64 `json:"open"`
			Price    float64 `json:"price"`
			Seq      int     `json:"seq"`
			Time     int     `json:"time"`
			TickDeep []struct {
				PriceBid  int     `json:"price_bid"`
				VolumeBid int     `json:"volume_bid"`
				PriceAsk  float64 `json:"price_ask"`
				VolumeAsk int     `json:"volume_ask"`
			} `json:"tick_deep"`
			PriceClose float64 `json:"price_close"`
			Digits     float64 `json:"digits"`
		} `json:"tick"`
	} `json:"pb"`
}
