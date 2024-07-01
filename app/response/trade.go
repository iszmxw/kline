package response

type UserTradeMessage struct {
	AppUserId            string  `json:"appUserId"`            // 用户ID
	IsNew                int     `json:"isNew"`                // 是否新客户 1是 0否
	BlackList            int     `json:"blackList"`            // 是否黑名单 1是 0否
	TradingCount         int     `json:"tradingCount"`         // 当日交易次数
	TradingAmounts       float64 `json:"tradingAmounts"`       // 当日交易金额
	RechargeAmounts      float64 `json:"rechargeAmounts"`      // 历史充值金额
	RechargeCount        int     `json:"rechargeCount"`        // 历史充值次数
	RechargeTodayAmounts float64 `json:"rechargeTodayAmounts"` // 当天充值金额
	RechargeTodayCount   int     `json:"rechargeTodayCount"`   // 当天充值次数
	WithdrawAmounts      float64 `json:"withdrawAmounts"`      // 历史提现金额
	WithdrawCount        int     `json:"withdrawCount"`        // 历史提现次数
	WithdrawTodayAmounts float64 `json:"withdrawTodayAmounts"` // 当天提现金额
	WithdrawTodayCount   int     `json:"withdrawTodayCount"`   // 当天提现次数
	HisTradingCount      int     `json:"hisTradingCount"`      // 历史交易次数
	HisTradingAmounts    float64 `json:"hisTradingAmounts"`    // 历史交易金额
}

type GetTradingAssetsResult struct {
	TradingAssetsList []TradingAssetsBean `json:"tradingAssetsList"` // 用户ID
	Total             int                 `json:"total"`             // 是否新客户 1是 0否
}

// TradingAssetsBean 表示交易资产的结构体
type TradingAssetsBean struct {
	Id             int64   `json:"id"`             // 资产表主键ID
	AssetsName     string  `json:"assetsName"`     // 资产名称
	AssetsCode     string  `json:"assetsCode"`     // 资产编码
	AssetsLogoUrl  string  `json:"assetsLogoUrl"`  // 资产logo链接
	ReturnRate     float64 `json:"returnRate"`     // 收益率
	AssetsStatus   int     `json:"assetsStatus"`   // 资产状态 1正常 2休市 3下架
	IsValid        int     `json:"isValid"`        // 是否有效 1是 0否
	IsShow         int     `json:"isShow"`         // 是否展示 1是 0否
	PointNumber    int     `json:"pointNumber"`    // 保留小数点位数
	AssetsTypeCode string  `json:"assetsTypeCode"` // 资产类型编码
	AssetsTypeName string  `json:"assetsTypeName"` // 资产类型名称
}
