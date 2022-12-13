package models

import "time"

// 应用表

type AppItem struct {
	Id                    int64     `json:"id"`                    //应用表主键ID
	TimeZone              string    `json:"time_zone"`             //时区
	CreateDate            time.Time `json:"create_date"`           //创建日期
	Operater              string    `json:"operater"`              //操作人
	ItemCode              string    `json:"item_code"`             //应用编码
	ItemName              string    `json:"item_name"`             //应用名称
	IsStart               int       `json:"is_start"`              //是否有效 1是 0否
	Corporation           string    `json:"corporation"`           //法人代表
	ConnetionMethod       string    `json:"connetion_method"`      //联系方式
	AppLogoUrl            string    `json:"app_logo_url"`          //应用logo
	AppVersion            string    `json:"app_version"`           //应用版本号
	AppGroom              string    `json:"app_groom"`             //APP推荐理由
	Score                 string    `json:"score"`                 //评分
	Describes             string    `json:"describes"`             //应用描述
	ProductName           string    `json:"product_name"`          //产品名称
	Sort                  int       `json:"sort"`                  //排序
	PageVersion           int       `json:"page_version"`          //页面版本
	TermsAgreement        string    `json:"terms_agreement"`       //条款协议URL
	PrivacyAgreement      string    `json:"privacy_agreement"`     //隐私协议 URL
	Indexweburl           string    `json:"indexWebUrl"`           //appH5网页-首页URL
	Transactionweburl     string    `json:"transactionWebUrl"`     //appH5网页-交易URL
	Courseurl             string    `json:"courseUrl"`             //appH5网页-教程URL
	Commonproblemurl      string    `json:"commonProblemUrl"`      //appH5网页-常见问题URL
	Rechargeguidelinesurl string    `json:"rechargeGuidelinesUrl"` //appH5网页-充值指引URL
	Invitationurl         string    `json:"invitationUrl"`         //appH5网页-邀请页面URL
	TradeLossReturnUrl    string    `json:"trade_loss_return_url"` //appH5网页-亏损返还URL
	DownloadIos           string    `json:"download_ios"`          //应用下载链接ios
	DownloadAndriod       string    `json:"download_andriod"`      //应用下载链接andriod
	ServicePhone          string    `json:"service_phone"`         //服务电话
	GoMail                string    `json:"go_mail"`               //邮件服务地址
	ItemType              int       `json:"item_type"`             //项目类型 1正式应用 2测试应用 3贷超应用 4外包出去的app
	UpdateDate            time.Time `json:"update_date"`           //更新日期
	AgentLocation         string    `json:"agent_location"`        //代理公司地址
	Email                 string    `json:"email"`                 //公司邮箱
	Adress                string    `json:"adress"`                //公司地址
	NumberPoint           string    `json:"number_point"`          //保留小数点位数
	DemoBalance           float64   `json:"demo_balance"`          //初始化demo余额
	ItemNationalName      string    `json:"item_national_name"`    //项目国家名称
	ItemNationalFlag      string    `json:"item_national_flag"`    //项目国国旗
	LanguageCode          string    `json:"language_code"`         //语言编码
	WebUrl                string    `json:"web_url"`               //网页链接
	WebUrlRemark          string    `json:"web_url_remark"`        //网页链接参数说明
	MonetaryUnit          string    `json:"monetary_unit"`         //货币单位
	CurrencyType          string    `json:"currency_type"`         //货币名称 IDR USD CNY
}
