package requests

type RegisterParams struct {
	Nation   string `json:"nation" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SubStatistics struct {
	Sub SubBody `json:"sub"`
	Id  string  `json:"id"`
}

type SubBody struct {
	Nation string `json:"nation"`
}
