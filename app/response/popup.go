package response

type PopupMsg struct {
	Body         string `json:"body"`
	Confirm      string `json:"confirm"`
	CreatDate    int64  `json:"creatDate"`
	Creator      string `json:"creator"`
	Id           int    `json:"id"`
	Img          string `json:"img"`
	ItemCode     string `json:"itemCode"`
	LanguageCode string `json:"languageCode"`
	Name         string `json:"name"`
	Remarks      string `json:"remarks"`
	Title        string `json:"title"`
	Type         int    `json:"type"`
	Redirect     string `json:"redirect"`
	UpdateDate   int64  `json:"updateDate"`
	Updater      string `json:"updater"`
}
