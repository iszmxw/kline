package models

import "time"

// TSystemDataDictionary undefined
type TSystemDataDictionary struct {
	ID         int64     `json:"id"`          // 数据字典主键
	Name       string    `json:"name"`        // 数据key
	Value      string    `json:"value"`       // 数据value
	Nation     string    `json:"nation"`      // 绑定国家
	CreateDate time.Time `json:"create_date"` // 创建时间
	UpdateDate time.Time `json:"update_date"` // 修改时间
	Remark     string    `json:"remark"`      // 备注
}

// TableName 表名称
func (*TSystemDataDictionary) TableName() string {
	return "t_system_data_dictionary"
}
