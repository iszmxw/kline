package response

type RespUploadBase64Str struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Data    struct {
		Url string `json:"url"`
	} `json:"data"`
}
