package requests

type CashFlowParams struct {
	Base
	PageInfo
	Today     string `json:"today"`
	OrderType string `json:"orderType"`
}
