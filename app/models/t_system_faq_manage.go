package models

// TSystemFaqManage undefined
type TSystemFaqManage struct {
	ID         int64  `json:"id" gorm:"id"`                   // faq表管理主键id
	Type       string `json:"type" gorm:"type"`               // 问题分类
	Question   string `json:"question" gorm:"question"`       // 问题
	Answer     string `json:"answer" gorm:"answer"`           // 答案
	Status     int64  `json:"status" gorm:"status"`           // 状态
	ItemCode   string `json:"item_code" gorm:"item_code"`     // 项目编码
	Nation     string `json:"nation" gorm:"nation"`           // 国家
	IsDelete   int64  `json:"is_delete" gorm:"is_delete"`     // 是否逻辑删除
	CreateTime int64  `json:"create_time" gorm:"create_time"` // 创建时间
	UpdateTime int64  `json:"update_time" gorm:"update_time"` // 更新时间
	Sort       int64  `json:"sort" gorm:"sort"`               // 排序
}

// TableName 表名称
func (*TSystemFaqManage) TableName() string {
	return "t_system_faq_manage"
}
