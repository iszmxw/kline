package models

import (
	"goapi/pkg/config"
	"goapi/pkg/helpers"
	"strings"
)

// BaseModel 模型基类
type BaseModel struct {
	Id uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
}

// BaseModelLast 模型基类
type BaseModelLast struct {
	CreatedAt helpers.TimeNormal `gorm:"column:created_at;index;comment:创建时间" json:"created_at"`
	UpdatedAt helpers.TimeNormal `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`
	// 支持 gorm 软删除
	DeletedAt helpers.TimeNormal `gorm:"column:deleted_at;index;comment:删除时间" json:"deleted_at"`
}

// GetStringID 获取 ID 的字符串格式
func (m *BaseModel) GetStringID() string {
	return helpers.Uint64ToString(m.Id)
}

// PageList 分页返回数
type PageList struct {
	CurrentPage int64       `json:"current_page"`
	Total       int64       `json:"total"`
	List        interface{} `json:"list"`
	PageNum     int64       `json:"pageNum"`  // 第几页
	PageSize    int64       `json:"pageSize"` // 每页多少条
}

// InitPageList 设置分页参数
func InitPageList(lists *PageList) {
	// 当前页数
	if lists.CurrentPage <= 0 {
		lists.CurrentPage = 1
	} // 第一页
	if lists.PageNum == 0 {
		lists.PageNum = 1
	}
	// 每页取出多少条数据，默认10条
	if lists.PageSize == 0 {
		lists.PageSize = 10
	}
}

// Prefix 置换sql中的表前缀
func Prefix(str string) string {
	prefix := config.GetString("database.mysql.prefix")
	return strings.Replace(str, "$_", prefix, -1)
}
