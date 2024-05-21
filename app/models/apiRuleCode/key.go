package apiRuleCode

// Code 枚举类型
type Code int

const (
	APP_AES_PARAMS Code = iota + 1
	WEB_SUBMIT_FILE_HOST
	APP_TIPS_INFO
	EMAIL_OTP_CONTENT_SET
	PHONE_OTP_CONTENT_SET
	SEND_ADJUST_HIDDEN_RULE
	FLOW_CONFIGURATION
	NATION_SET_RULE
	MOBILE_POWER_CONFIG
	USEING_CREDIT_RULE
	OS_SHOW_ASSETS_RULE
)

// ApiRuleCode 存储键值对应关系和备注信息的结构体
type ApiRuleCode struct {
	Code  string // 键值
	Alias string // 别名
}

// CodeMap 存储键值对应关系和备注信息的映射
var CodeMap = map[Code]ApiRuleCode{
	APP_AES_PARAMS:          {"APP_AES_PARAMS", "移动端接口加密设置"},
	WEB_SUBMIT_FILE_HOST:    {"WEB_SUBMIT_FILE_HOST", "提交文件配置"},
	APP_TIPS_INFO:           {"APP_TIPS_INFO", "APP提示信息配置"},
	EMAIL_OTP_CONTENT_SET:   {"EMAIL_OTP_CONTENT_SET", "发送邮件验证码配置"},
	PHONE_OTP_CONTENT_SET:   {"PHONE_OTP_CONTENT_SET", "发送手机验证码配置"},
	SEND_ADJUST_HIDDEN_RULE: {"SEND_ADJUST_HIDDEN_RULE", "adjust埋点事件发送事件配置"},
	FLOW_CONFIGURATION:      {"FLOW_CONFIGURATION", "埋点事件第三方配置"},
	NATION_SET_RULE:         {"NATION_SET_RULE", "获取国家对应的货币保留小数位数和单位"},
	MOBILE_POWER_CONFIG:     {"MOBILE_POWER_CONFIG", "获取手机号国家启用状态规则"},
	USEING_CREDIT_RULE:      {"USEING_CREDIT_RULE", "认证项配置规则"},
	OS_SHOW_ASSETS_RULE:     {"OS_SHOW_ASSETS_RULE", "操作系统展示规则"},
}

// ApiRuleInfo 根据键代码获取键值
func ApiRuleInfo(Code Code) ApiRuleCode {
	return CodeMap[Code]
}
