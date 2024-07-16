package requests

type CashFlowParams struct {
	Base
	PageInfo
	Today     string `json:"today"`
	OrderType string `json:"orderType"`
}

type OrderDetailParams struct {
	Base
	PageInfo
	OrderId string `json:"orderId" validate:"required"` // 订ID
}
