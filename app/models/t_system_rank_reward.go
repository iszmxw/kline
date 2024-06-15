package models

import "time"

// TSystemRankReward undefined
type TSystemRankReward struct {
	ID             int64     `json:"id" gorm:"id"`                             // 排行榜奖励主键id
	RewardType     int64     `json:"reward_type" gorm:"reward_type"`           // 奖励类型 1-bonus奖励 2-充值赠送次数奖励
	RewardAmount   float64   `json:"reward_amount" gorm:"reward_amount"`       // 奖励数额
	LowTradeAmount float64   `json:"low_trade_amount" gorm:"low_trade_amount"` // 最低交易金额
	LowTradeCount  int64     `json:"low_trade_count" gorm:"low_trade_count"`   // 最低交易次数
	MaxRank        int64     `json:"max_rank" gorm:"max_rank"`                 // 最高排名
	MinRank        int64     `json:"min_rank" gorm:"min_rank"`                 // 最低排名
	ItemCode       string    `json:"item_code" gorm:"item_code"`               // 关联项目
	Nation         string    `json:"nation" gorm:"nation"`                     // 绑定国家
	CreateDate     time.Time `json:"create_date" gorm:"create_date"`           // 创建时间
	UpdateDate     time.Time `json:"update_date" gorm:"update_date"`           // 修改时间
	Creator        string    `json:"creator" gorm:"creator"`                   // 创建人
	Updater        string    `json:"updater" gorm:"updater"`                   // 修改人
	IsDeleted      int64     `json:"is_deleted" gorm:"is_deleted"`             // 逻辑删除
}

// TableName 表名称
func (*TSystemRankReward) TableName() string {
	return "t_system_rank_reward"
}
