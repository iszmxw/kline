package requests

type RegisterParams struct {
	LanguageCode string `json:"languageCode" validate:"required"`
	Pid          string `json:"pid"`
	InviteCode   string `json:"inviteCode"`
	Nation       string `json:"nation"`
	Email        string `json:"email" validate:"required,email"`
	Code         string `json:"code" validate:"required"`
	Password     string `json:"password" validate:"required"`
}

type SubStatistics struct {
	Sub interface{} `json:"sub"`
	Id  string      `json:"id"`
}

type SubBody struct {
	Nation string `json:"nation"`
}
