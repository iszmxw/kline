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
