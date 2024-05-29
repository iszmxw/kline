package response

type WebLoginUser struct {
	IsWorkSubmit       int    `json:"isWorkSubmit"`
	ItemCode           string `json:"itemCode"`
	LastWorkSubmitDate string `json:"lastWorkSubmitDate"`
	Nation             string `json:"nation"`
	RoleCode           string `json:"roleCode"`
	SystemUserId       string `json:"systemUserId"`
	UserNick           string `json:"userNick"`
}

// 登录成功返回信息

type SuccessUserLogin struct {
	AppItems          ItemNationMessageParams `json:"appItems"`          // appItems 信息
	Avatar            string                  `json:"avatar"`            // 头像
	IsExistEmail      int                     `json:"isExistEmail"`      // 是否绑定了邮箱
	IsHiddenIosAssets int                     `json:"isHiddenIosAssets"` // 是否隐藏 ios 资产
	IsVisit           int                     `json:"isVisit"`           // 是否游客
	Token             string                  `json:"token"`             // 用户的 token
	UserId            string                  `json:"userId"`            // 用户的 userID
}
