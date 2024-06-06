package models

// TSystemBank models
type TSystemBank struct {
	ID       int64  `json:"id" gorm:"id"`               // 系统银行映射表
	BankId   string `json:"bank_id" gorm:"bank_id"`     // 银行唯一id
	BankName string `json:"bank_name" gorm:"bank_name"` // 银行名称
	Nation   string `json:"nation" gorm:"nation"`       // 国家
}

// TableName 表名称
func (*TSystemBank) TableName() string {
	return "t_system_bank"
}
