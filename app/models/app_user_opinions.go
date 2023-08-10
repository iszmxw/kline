package models

import "time"

type AppUserOpinions struct {
	ID                 int64     `json:"id"`                  // 用户反馈表id
	AppUserID          string    `json:"app_user_id"`         // 关联客户ID
	UserName           string    `json:"user_name"`           // 用户姓名
	Phone              string    `json:"phone"`               // 电话号码
	ProblemDescription string    `json:"problem_description"` // 问题描述
	Email              string    `json:"email"`               // 用户邮箱
	QuestionType       int       `json:"question_type"`       // 问题类型1充值，2提现，3交易，4其他
	OrderID            string    `json:"order_id"`            // 订单ID
	ProblemPhotoUrl    string    `json:"problem_photo_url"`   // 问题展示图片
	IsRead             int       `json:"is_read"`             // 用户是否已读，0-未读，1已读
	IsSolve            int       `json:"is_solve"`            // 是否需要处理，0未处理，1已处理
	IsDelete           int       `json:"is_delete"`           // 是否删除 1删除 0不删除
	Handler            string    `json:"handler"`             // 处理人
	CreatTime          time.Time `json:"creat_time"`          // 创建时间
	HandleDate         time.Time `json:"handle_date"`         // 处理日期
	Remark             string    `json:"remark"`              // 备注
	NewsExpand         string    `json:"news_expand"`         // 扩展字段
	SubmitIp           string    `json:"submit_ip"`           // 提交时网络IP
	Nation             string    `json:"nation"`              // 绑定国家
}

func (m *AppUserOpinions) TableName() string {
	return "t_app_user_opinions"
}
