package response

type GetUserLossReturn struct {
	TodayBonusReturn float64 `json:"todayBonusReturn"` // 计算得出用户待返还Bonus
	TotalBonusReturn float64 `json:"totalBonusReturn"` // 用户累计获得返还Bonus
	VipRate          float64 `json:"vipRate"`          // 亏损返还比例
	VipIconUrl       string  `json:"vipIconUrl"`       // 会员等级图标URL
}

type GetBonusTypes struct {
	List  []BonusTypes      `json:"list"`  // BonusTypes
	Types map[string]string `json:"types"` // types 键值对 key->value
}

type BonusTypes struct {
	ActivityCode string `json:"activityCode"` // 活动代码
	ActivityName string `json:"activityName"` // 活动名称
}

type GetBonusRecords struct {
	ActivityLogUrl string `json:"activityLogUrl"` // 活动log链接
	ActivityName   string `json:"activityName"`   // 活动名称
	CreateDate     int64  `json:"createDate"`     // 赠送时间
	Amounts        int    `json:"amounts"`        // Bouns赠送金额
	MonetaryUnit   string `json:"monetaryUnit"`   // bonus赠送货币单位
	ShowAmounts    string `json:"showAmounts"`    // 展示Bouns赠送金额
	IsValid        int    `json:"isValid"`        // 是否过期
}

// BannerList 结构体用于表示 BannerParams数据
type BannerList struct {
	BannerName   string `json:"bannerName"`   // Banner名称
	Sort         int    `json:"sort"`         // 排序顺序
	Position     int    `json:"position"`     // 广告位置（1：首页轮播）
	PhotoUrl     string `json:"photoUrl"`     // 图片链接URL
	Description  string `json:"description"`  // 文字描述信息
	ActionUrl    string `json:"actionUrl"`    // 跳转链接URL
	ShowTime     int    `json:"showTime"`     // 显示时间（1：长期展示，2：未登录时，3：首充前，4：有充值活动次数时候）
	JumpPosition int    `json:"jumpPosition"` // 跳转位置
}

type FAQManagerList struct {
	Type     string `json:"type"`     // 问题分类
	Question string `json:"question"` // 问题
	Answer   string `json:"answer"`   // 答案
}

type UserRank struct {
	UserTradingRankList []UserTradingRankList `json:"userTradingRankList"` // 用户交易排行榜列表
	MyRank              UserTradingRankList   `json:"myRank"`              // 我的排行信息
	PageNum             int                   `json:"pageNum"`             // 页数
	PageSize            int                   `json:"pageSize"`            // 条数
	Page                int                   `json:"page"`                // 第几页
	Total               int                   `json:"total"`               // 共有多少条
	NextMondayTime      int                   `json:"nextMondayTime"`      // 下周一时间
}

type UserTradingRankList struct {
	Avatar             string      `json:"avatar"`             // 头像
	AppUserId          string      `json:"appUserId"`          // 用户ID
	TradingAmounts     string      `json:"tradingAmounts"`     // 交易金额
	BonusAmounts       string      `json:"bonusAmounts"`       // bonus金额
	No                 int         `json:"no"`                 // 排名
	ShowTradingAmounts string      `json:"showTradingAmounts"` // 展示交易金额
	ShowBonusAmounts   string      `json:"showBonusAmounts"`   // 展示 bonus 金额
	TradeContent       interface{} `json:"tradeContent"`       // 交易说明文案
}
