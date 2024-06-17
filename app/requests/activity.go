package requests

type GetBonusRecordsParams struct {
	Base
	ActivityCode string `json:"activityCode"`                // 活动编码 eg:first_recharge
	IsToday      string `json:"isToday" validate:"required"` // 是否查询今天 ，1，是，0-查询所有
	PageNum      int64  `json:"pageNum" validate:"required"` // 页码
	PageSize     int64  `json:"pageSize"`                    // 单页条数
}

type FAQManagerParams struct {
	Base
	Type string `json:"type" validate:"required"` // FAQ类型
}

type GetUserRankParams struct {
	Base
	WhichWeek string `json:"whichWeek" validate:"required"` // 哪一周 1本周 2上周
	PageNum   int    `json:"pageNum"`                       // 第几页
	PageSize  int    `json:"pageSize"`                      // 每页多少条
}

type GetUserSignInParams struct {
	Base
	SignInTime     int64 `json:"signInTime"`     // 签到时间
	FirebaseNotify int64 `json:"firebaseNotify"` // 是否开启firebase签到推送 1-开启 0-关闭
}
