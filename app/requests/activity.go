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

// SiteInformationCacheResult 站内信缓存结果
type SiteInformationCacheResult struct {
	ItemCode     string // 应用编码
	Nation       string // 国家 Indonesia、India
	PushPlatForm string // 推送平台
	TopicName    string // 主题名称
	ImageUrl     string // 将在通知中显示的图像的URL
	Url          string // 跳转url
	DeciveBrand  string // 设备品牌 apple(android)
	MailId       int64  // 站内信发送表id
	MessageId    int64  // 站内信信息模板表id
	SendTime     int64  // 发送时间
	AppUserIds   string // 将要批量发送的用户 当为 ALL 时发送给所有用户
}
