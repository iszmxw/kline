package models

import "time"

// TAppItemVersion 结构体
type TAppItemVersion struct {
	ID             int64     `json:"id" gorm:"id"`                             // 应用关联版本号
	ItemCode       string    `json:"item_code" gorm:"item_code"`               // 关联应用编码
	VersionString  string    `json:"version_string" gorm:"version_string"`     // 版本号字符串
	VersionValue   int64     `json:"version_value" gorm:"version_value"`       // 版本号整数值
	CreateDate     time.Time `json:"create_date" gorm:"create_date"`           // 创建日期
	IsUsed         int64     `json:"is_used" gorm:"is_used"`                   // 是否在使用 1是 0否
	IsReinstall    int64     `json:"is_reinstall" gorm:"is_reinstall"`         // 是否为必须更新的版本
	IsNew          int64     `json:"is_new" gorm:"is_new"`                     // 是否为最新版本
	UpdateDate     time.Time `json:"update_date" gorm:"update_date"`           // 更新日期
	AppType        int64     `json:"app_type" gorm:"app_type"`                 // 应用类型 1安卓 2IOS 3html
	AppDownloadUrl string    `json:"app_download_url" gorm:"app_download_url"` // 应用下载链接
	Nation         string    `json:"nation" gorm:"nation"`                     // 国家
	Google         int64     `json:"google" gorm:"google"`                     // 是否开启谷歌登录1是，2否
	Facebook       int64     `json:"facebook" gorm:"facebook"`                 // 是否开启fackbook登录，1是2否
	Mobile         int64     `json:"mobile" gorm:"mobile"`                     // 是否开启手机登录，1是2否
	Email          int64     `json:"email" gorm:"email"`                       // 是否开启邮箱登录，1是2否
	Title          string    `json:"title" gorm:"title"`                       // 版本弹窗标题
	Description    string    `json:"description" gorm:"description"`           // 版本弹窗详情
	IsDeleted      int64     `json:"is_deleted" gorm:"is_deleted"`             // 逻辑删除
}

// TableName 表名称
func (*TAppItemVersion) TableName() string {
	return "t_app_item_version"
}
