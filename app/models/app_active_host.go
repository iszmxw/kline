package models

import "time"

// AppActiveHost app活跃域名表
type AppActiveHost struct {
	Id         int64     `gorm:"column:id; PRIMARY_KEY" json:"id"` //app活跃域名表主键自增ID
	ActiveHost string    `json:"active_host"`                      //活跃的域名
	CreateDate time.Time `json:"create_date"`                      //创建日期
	UpdateDate time.Time `json:"update_date"`                      //修改日期
	Creator    string    `json:"creator"`                          //添加人ID
	Updater    string    `json:"updater"`                          //更新人ID
	ItemCode   string    `json:"item_code"`                        //关联项目编码
	IsShow     int       `json:"is_show"`                          //是否展示 1是 0否
	IsDelete   int       `json:"is_delete"`                        //是否删除 1是 0否
	Level      int       `json:"level"`                            //活跃等级
	Nation     string    `json:"nation"`                           //绑定国家
}
