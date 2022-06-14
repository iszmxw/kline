package models

import "time"

type AppUser struct {
	AppUserId            string    `json:"app_user_id"`           //app用户表
	Mobile               string    `json:"mobile"`                //手机号
	Email                string    `json:"email"`                 //邮箱
	Password             string    `json:"password"`              //密码
	CreateDate           time.Time `json:"create_date"`           //创建日期
	UpdateDate           time.Time `json:"update_date"`           //更新日期
	ItemCode             string    `json:"item_code"`             //关联项目编码
	BlackList            int       `json:"black_list"`            //是否黑名单 1是 0否
	Balance              float64   `json:"balance"`               //余额
	VirtualBalance       float64   `json:"virtual_balance"`       //虚拟余额
	BonusBalance         float64   `json:"bonus_balance"`         //奖励余额
	CurrencyType         string    `json:"currency_type"`         //货币类型 IDR USD CNY
	UserLevel            int       `json:"user_level"`            //用户等级
	IsNew                int       `json:"is_new"`                //是否新用户 1是 0否 首充后变成0
	Avatar               string    `json:"avatar"`                //头像链接
	ReferralCode         string    `json:"referral_code"`         //邀请码
	Gender               int       `json:"gender"`                //性别
	CreditEmail          string    `json:"credit_email"`          //认证邮箱
	IdBackUrl            string    `json:"id_back_url"`           //证件反面链接
	UserName             string    `json:"user_name"`             //用户姓名
	ParentCode           string    `json:"parent_code"`           //父级邀请码
	InvitDate            time.Time `json:"invit_date"`            //受邀请日期
	LastLoginIp          string    `json:"last_login_ip"`         //最后的登录IP
	Token                string    `json:"token"`                 //登录凭证
	LastToken            string    `json:"last_token"`            //最后的登录凭证
	DeciveId             string    `json:"decive_id"`             //设备ID
	DeciveBrand          string    `json:"decive_brand"`          //设备品牌 多个用,号隔开 当前为最后一个
	HadChangeDecive      int       `json:"had_change_decive"`     //是否更换设备
	DeviceModel          string    `json:"device_model"`          //设备型号
	AppPackage           string    `json:"app_package"`           //应用包名
	LastLoginDate        time.Time `json:"last_login_date"`       //最后登录日期
	PullBlackDate        time.Time `json:"pull_black_date"`       //拉黑日期
	BlackRemarks         string    `json:"black_remarks"`         //拉黑备注
	RegistIp             string    `json:"regist_ip"`             //注册时IP
	RegistType           string    `json:"regist_type"`           //客户来源方式
	IsLogout             int       `json:"is_logout"`             //是否注销 1是 0否
	AppVersion           string    `json:"app_version"`           //注册时的app版本号
	IdFrontUrl           string    `json:"id_front_url"`          //证件正面链接
	CreditMobile         string    `json:"credit_mobile"`         //认证手机号
	IdentityNo           string    `json:"identity_no"`           //认证证件身份ID
	Nation               string    `json:"nation"`                //绑定国家
	Religion             string    `json:"religion"`              //宗教
	City                 string    `json:"city"`                  //城市
	IsNeedKyc            int       `json:"is_need_kyc"`           //是否需要去认证KYC 1是 0否 3完成KYC认证
	OptimizationDivision string    `json:"optimization_division"` //优化师
	TrackChannel         string    `json:"track_channel"`         //投放方式:自投,代投,代投公司
	Network              string    `json:"network"`               //投放渠道
	FlowCompanyCode      string    `json:"flow_company_code"`     //整合流量方公司编码
	FlowBindId           string    `json:"flow_bind_id"`          //整合流量方公司绑定ID
	Campaign             string    `json:"campaign"`              //投放渠道活动名称
	UserSource           string    `json:"user_source"`           //用户来源，关联表t_app_channel
}
