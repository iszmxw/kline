package response

type GetUserWithdrawList struct {
	BankList []BankList `json:"bankList"`
	Total    int64      `json:"total"`
}

type BankList struct {
	BankId      int64  `json:"bankId"`      // 银行默认设置表ID
	BankName    string `json:"bankName"`    // 银行名称
	IsRecommend int64  `json:"isRecommend"` // 是否推荐
	Sort        int64  `json:"sort"`        // 排序
}
