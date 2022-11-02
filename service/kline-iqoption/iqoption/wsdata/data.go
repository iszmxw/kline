package wsdata

type SubPriceKLineResponse struct {
	ActiveId int     `json:"active_id"`
	Size     int     `json:"size"`
	At       int64   `json:"at"`
	From     int     `json:"from"`
	To       int     `json:"to"`
	Id       int     `json:"id"`
	Open     float64 `json:"open"`
	Close    float64 `json:"close"`
	Min      float64 `json:"min"`
	Max      float64 `json:"max"`
	Ask      float64 `json:"ask"`
	Bid      float64 `json:"bid"`
	Volume   int     `json:"volume"`
	Phase    string  `json:"phase"`
}

// {"name":"sendMessage","request_id":"319","local_time":176414,"msg":{"name":"get-first-candles","version":"1.0","body":{"active_id":1,"split_normalization":true}}}
// {"name":"sendMessage","request_id":"404","local_time":248116,"msg":{"name":"get-first-candles","version":"1.0","body":{"active_id":3,"split_normalization":true}}}
// 订阅消息

type Resp struct {
	Name             string      `json:"name"`
	MicroserviceName string      `json:"microserviceName"`
	Msg              interface{} `json:"msg"`
}

type SendMessage struct {
	Name      string `json:"name"`
	RequestId string `json:"request_id"`
	LocalTime int    `json:"local_time"`
	Msg       Sub    `json:"msg"`
}

type Sub struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Body    Body   `json:"body"`
}

type Body struct {
	ActiveId           int  `json:"active_id"`
	SplitNormalization bool `json:"split_normalization"`
}

type SendPing struct {
	Name      string `json:"name"`
	RequestId string `json:"request_id"`
	LocalTime int    `json:"local_time"`
	Msg       Msg    `json:"msg"`
}

type Msg struct {
	UserTime      string `json:"userTime"`
	HeartbeatTime string `json:"heartbeatTime"`
}

type SendLogin struct {
	Name      string `json:"name"`
	RequestId string `json:"request_id"`
	LocalTime int    `json:"local_time"`
	Msg       Login  `json:"msg"`
}

type Login struct {
	Ssid            string `json:"ssid"`
	Protocol        int    `json:"protocol"`
	SessionId       string `json:"session_id"`
	ClientSessionId string `json:"client_session_id"`
}

type Authenticated struct {
	Name            string `json:"name"`
	Msg             bool   `json:"msg"`
	ClientSessionId string `json:"client_session_id"`
	RequestId       string `json:"request_id"`
}
