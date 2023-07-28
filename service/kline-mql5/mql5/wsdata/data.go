package wsdata

type Mql5 struct {
	Name       string      `json:"name"`
	Bid        float64     `json:"bid"`
	Ask        float64     `json:"ask"`
	Flags      int         `json:"flags"`
	Last       interface{} `json:"last"`
	Time       string      `json:"time"`
	TimeMsc    int64       `json:"timeMsc"`
	Volume     interface{} `json:"volume"`
	VolumeReal interface{} `json:"volumeReal"`
}
