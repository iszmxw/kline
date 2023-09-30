package wsdata

import (
	"time"
)

type SubPriceKLineResponse struct {
	Rate      float64   `json:"rate"`
	Precision int       `json:"precision"`
	Repeat    int       `json:"repeat"`
	Ask       float64   `json:"ask"`
	CreatedAt time.Time `json:"created_at"`
	Bid       float64   `json:"bid"`
	Ric       string    `json:"ric"`
}

type DataItem struct {
	Assets []SubPriceKLineResponse `json:"assets"`
	Action string                  `json:"action"`
}

type Res struct {
	Data    []DataItem    `json:"data"`
	Success bool          `json:"success"`
	Errors  []interface{} `json:"errors"`
}
