package requests

type GetLanguagePageParams struct {
	Base
	Nation string `json:"nation" validate:"required"`
	Type   string `json:"type" validate:"required"`
	Uuid   string `json:"uuid"`
	Adid   string `json:"adid"`
}
