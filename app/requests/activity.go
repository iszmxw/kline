package requests

type GetBonusRecordsParams struct {
	Base
	ActivityCode string `json:"activityCode" validate:"required"` // 项目编码
	IsToday      string `json:"isToday" validate:"required"`      // 国家编码
	PageNum      string `json:"pageNum" validate:"required"`      // 语言编码
	PageSize     string `json:"pageSize"`                         // 系统类型 1安卓 2IOS 3html
}
