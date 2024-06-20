package response

import "time"

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
	PageNum             int                   `json:"pageNum"`             // 页数
	PageSize            int                   `json:"pageSize"`            // 条数
	Page                int                   `json:"page"`                // 第几页
	Total               int                   `json:"total"`               // 共有多少条
	NextMondayTime      int                   `json:"nextMondayTime"`      // 下周一时间
	UserTradingRankList []UserTradingRankList `json:"userTradingRankList"` // 用户交易排行榜列表
	MyRank              UserTradingRankList
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

type WeekDataV2 struct {
	List           []WeekDataList `json:"list"`
	FirebaseNotify int32          `json:"firebaseNotify"` // 获取是否开启firebase通知 1-开启 0-关闭
}

type WeekDataList struct {
	Bonus  string `json:"bonus"`  // 奖励的 bonus
	Week   string `json:"week"`   // 展示的星期几
	Date   int64  `json:"date"`   // 本周七天日期，时间戳
	Desc   string `json:"desc"`   // 描述
	Imgurl string `json:"imgurl"` // 签到状态图标
	Status int    `json:"status"` // 签到状态 0-未签到待补签，1-已签到 2-签到日等待签到 3-未到签到日
}

// AppSignIn 是签到信息结构体
type AppSignIn struct {
	ID            int       `json:"id"`            // 主键id
	AppUserID     string    `json:"appUserId"`     // 用户ID
	ItemCode      string    `json:"itemCode"`      // 关联项目编码
	BonusBalance  float64   `json:"bonusBalance"`  // 奖励余额
	WhichDay      string    `json:"whichDay"`      // 周几
	GiveAwayBonus float64   `json:"giveAwayBonus"` // 赠送bonus
	DeviceModel   string    `json:"deviceModel"`   // 设备型号
	SigninType    int       `json:"signinType"`    // 签到类型 1-正常签到，2-补签签到
	Nation        string    `json:"nation"`        // 绑定国家
	SigninStatus  int       `json:"signinStatus"`  // 签到状态 0-未签到待补签，1-已签到 2-签到日等待签到 3-未到签到日
	SigninIP      string    `json:"signinIp"`      // 签到 IP
	SigninDate    time.Time `json:"signinDate"`    // 签到日期
	CreateDate    time.Time `json:"createDate"`    // 创建日期
	UpdateDate    time.Time `json:"updateDate"`    // 更新日期
}

type SignInRuleBeans struct {
	Rule         map[string]string           `json:"rule"`         // 规则
	MonetaryUnit string                      `json:"monetaryUnit"` // 货币符号
	EnumStatus   []SignInRuleEnumStatusBeans // 枚举
}

type SignInRuleEnumStatusBeans struct {
	Status int    `json:"status"` // 状态
	Desc   string `json:"desc"`   // 描述
	Imgurl string `json:"imgurl"` // 图片地址
}

type ActivitySetJson struct {
	Min          float64 `json:"min"`          // 最小值
	Max          float64 `json:"max"`          // 最大值
	DiscountRate float64 `json:"discountRate"` // 百分比
}
