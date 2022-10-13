package models

import "time"

// 代理用户表

type AgentUser struct {
	AgentUserId              string    `json:"agent_user_id"`               //agent用户表
	ItemCode                 string    `json:"item_code"`                   //关联项目编码
	Nation                   string    `json:"nation"`                      //绑定国家
	Email                    string    `json:"email"`                       //邮箱
	Password                 string    `json:"password"`                    //密码
	Mobile                   string    `json:"mobile"`                      //手机号
	Telegram                 string    `json:"telegram"`                    //telegram
	WhatsApp                 string    `json:"whats_app"`                   //whats_app
	Balance                  float64   `json:"balance"`                     //余额
	CurrencyType             string    `json:"currency_type"`               //货币类型 IDR USD CNY
	UserLevel                int       `json:"user_level"`                  //用户等级
	UserRate                 float64   `json:"user_rate"`                   //当前分红比例
	UserStatus               int       `json:"user_status"`                 //代理账号状态1正常2停用
	Avatar                   string    `json:"avatar"`                      //头像链接
	Token                    string    `json:"token"`                       //登录凭证
	ReferralCode             uint64    `json:"referral_code"`               //邀请码
	LastLoginDate            time.Time `json:"last_login_date"`             //最后登录日期
	BlackList                int       `json:"black_list"`                  //是否黑名单 1是 0否
	PullBlackDate            time.Time `json:"pull_black_date"`             //拉黑日期
	BlackRemarks             string    `json:"black_remarks"`               //拉黑备注
	RegistIp                 string    `json:"regist_ip"`                   //注册时IP
	RegistType               string    `json:"regist_type"`                 //客户来源方式
	RegistNation             string    `json:"regist_nation"`               //注册时国家
	CreateDate               time.Time `json:"create_date"`                 //创建日期
	UpdateDate               time.Time `json:"update_date"`                 //更新日期
	LevelStartSettlementDate time.Time `json:"level_start_settlement_date"` //等级开始结算周期,每30天为一周期
	LevelEndSettlementDate   time.Time `json:"level_end_settlement_date"`   //等级结算结束日，下一等级开始日
}
