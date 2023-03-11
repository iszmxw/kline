package wsdata

//type SubPriceKLineResponse struct {
//	E  string `json:"e"` // 事件类型
//	E1 int64  `json:"E"` // 事件时间
//	I  string `json:"i"` // 标的交易对
//	P  string `json:"p"` // 指数价格
//}

type SubPriceKLineResponse struct {
	E  string `json:"e"`  // 事件类型
	E1 int64  `json:"E"`  // 事件时间
	Ps string `json:"ps"` // 标的交易对
	K  struct {
		T  int64  `json:"t"` // 这根K线的起始时间
		T1 int64  `json:"T"` // 这根K线的结束时间
		S  string `json:"s"` // 无意义字段
		I  string `json:"i"` // K线间隔
		F  int64  `json:"f"` // 无意义字段
		L  int64  `json:"L"` // 无意义字段
		O  string `json:"o"` // 这根K线期间第一笔成交价
		C  string `json:"c"` // 这根K线期间末一笔成交价
		H  string `json:"h"` // 这根K线期间最高成交价
		L1 string `json:"l"` // 这根K线期间最低成交价
		V  string `json:"v"` // 无意义字段
		N  int    `json:"n"` // 这根K线更新 期间 数据数量
		X  bool   `json:"x"` // 这根K线是否完结(是否已经开始下一根K线)
		Q  string `json:"q"` // 无意义字段
		V1 string `json:"V"` // 无意义字段
		Q1 string `json:"Q"` // 无意义字段
		B  string `json:"B"` // 无意义字段
	} `json:"k"`
}
