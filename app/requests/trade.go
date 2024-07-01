package requests

type GetAssetsListParams struct {
	Base
	AssetsName string `json:"assetsName"` // 资产名称
	PageInfo
}
