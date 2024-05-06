package keyenum

// Code 枚举类型
type Code int

const (
	// 这里保留原始的常量定义
	INIT_NUMBER_POINT Code = iota + 1
	NATION_USER
	USER_OPTION_TYPE
	MOBILE_CONFIG_STATUS
	MOBILE_CONFIG_STATUS_LIST
	INITIALIZATION_PROMPT
	SMALL_SPEAKER_BROADCAST
	BONUS_TRADING_ACTIVITY_NATION
	LOCK_KEY
	APP_SHOW_IMAGES
	SETTING_VIP_RULE
	SETTING_VIP_RULE_SWITCH
	WEB_OPINIONS_TYPE
	SITE_INFORMATION_MAIL_ID
	EMAIL_MESSAGE_QUEUE
	LimitAddress
	PUSH_TOKEN_CHECK
	SYNC_REAL_TRADE_RESULT
	SYNC_DEMO_TRADE_RESULT
	CHARGE_LOCK
	AGENT_TOKEN
	AGENT_RULE
	AGENT_SMS_CODE
	AGENT_USER_NUMBER
	SYSTEM_EMAIL
	AGENT_H5_PROMOTE
	AGENT_LINK_EXPOSURE_NUMBER
	AGENT_USER_INSTALL_AMOUNT
	BLOCK_H5_PAY
	CLIENT_TOKEN
	WEB_TOKEN
	EMAIL_TEMPLATE
	CLIENT_POPUP
	NATION_IP_WHITELIST
	WITHDRAW_WHITELIST
	CLIENT_SMS_CODE
	SMART_START_STOP
	RECHARGE_DELAY_QUEUE
	RECHARGE_READY_QUEUE
	WITHDRAW_DELAY_QUEUE
	WITHDRAW_READY_QUEUE
	RECHARGE_ORDER_RETRY_TIMES
	WITHDRAW_ORDER_RETRY_TIMES
	IOS_VEST_ACCOUNT
	LANGUAGE_CODE
	SERVICE_IP
	SLOW_SQL_TIME_OUT
	SLOW_MAPPER_ID_SQL
	ONLINE_USERS
	INVITE_SUCCESSFUL_USERS
	EMAIL_SCORE
	LOGOUT_TOKEN_REASON
	ONLY_ONE_SIGN_KEY
	RECHARGE_SUCCESS_QUEUE
	NOTICE_ORDER_RESULT_QUEUE
	THIS_WEEK_INVIT_TOP_ONE_INFO
	REGISTRATION_RULE_JUDGMENT
	SYS_ORDER_ID
	SYS_ORDER_ID_QUEUE
	AllowKlineSource
	KlineSourceStatus
	KlineSource
	EMAILLIMITONE
)

// RedisServiceKeyEnum 存储键值对应关系和备注信息的结构体
type RedisServiceKeyEnum struct {
	Key    string // 键值
	Remark string // 备注信息
}

// KeyMap 存储键值对应关系和备注信息的映射
var KeyMap = map[Code]RedisServiceKeyEnum{
	INIT_NUMBER_POINT:             {"init:numberPoint", "初始化系统小数位数"},
	NATION_USER:                   {"nation:", "系统使用国家"},
	USER_OPTION_TYPE:              {"mine:opinionsType:", "客户反馈问题类型"},
	MOBILE_CONFIG_STATUS:          {"international:mobile:config:", "电话启用规则判断"},
	MOBILE_CONFIG_STATUS_LIST:     {"international:mobile:config:list", "国家电话启用规则列表"},
	INITIALIZATION_PROMPT:         {"init:API_RESULT", "初始化提示语"},
	SMALL_SPEAKER_BROADCAST:       {"init:trade:", "小喇叭播报"},
	BONUS_TRADING_ACTIVITY_NATION: {"time:lossReturn:", "bonus交易活动"},
	LOCK_KEY:                      {"redis:lock:", "redis锁key"},
	APP_SHOW_IMAGES:               {"config:image:", "app界面展示图片"},
	SETTING_VIP_RULE:              {"config:setting_vip_rule:", "VIP配置表信息"},
	SETTING_VIP_RULE_SWITCH:       {"config:setting_vip_rule_switch", "VIP配置表信息 开关，是否开启关闭"},
	WEB_OPINIONS_TYPE:             {"web:v1:mine:opinionsType:", "后台意见反馈类型"},
	SITE_INFORMATION_MAIL_ID:      {"SiteInformation:mailId:%s", "站内信"},
	EMAIL_MESSAGE_QUEUE:           {"EmailMessageQueue", "电子邮件消息队列"},
	LimitAddress:                  {"LimitAddress:", "站内信"},
	PUSH_TOKEN_CHECK:              {"check:pushToken", "pushTokenIsCheck"},
	SYNC_REAL_TRADE_RESULT:        {"RealOrder:Result:", "真实交易订单结果更新锁"},
	SYNC_DEMO_TRADE_RESULT:        {"DemoOrder:Result:", "虚拟交易订单结果更新锁"},
	CHARGE_LOCK:                   {"Charge:Order", "充值订单创建锁"},
	AGENT_TOKEN:                   {"Agent:Token:", "代理登录"},
	AGENT_RULE:                    {"Agent:rule:", "代理规则"},
	AGENT_SMS_CODE:                {"Agent:Sms:", "代理验证码"},
	AGENT_USER_NUMBER:             {"Agent:user:", "代理现在有用户"},
	SYSTEM_EMAIL:                  {"System:Email", "充提订单失败发送系统邮箱"},
	AGENT_H5_PROMOTE:              {"Agent:H5PROMOTE:", "代理H5推广地址"},
	AGENT_LINK_EXPOSURE_NUMBER:    {"Agent:LinkExposure:", "代理链接曝光数"},
	AGENT_USER_INSTALL_AMOUNT:     {"Agent:Install:", "代理安装数"},
	BLOCK_H5_PAY:                  {"Block:Pay:", "数字货币支付页面落地页"},
	CLIENT_TOKEN:                  {"Client:Token:", "Client端token分组"},
	WEB_TOKEN:                     {"Web:Token:", "web后台token分组"},
	EMAIL_TEMPLATE:                {"email:template:", "邮箱模板"},
	CLIENT_POPUP:                  {"Client:Popup:", "Client端弹窗"},
	NATION_IP_WHITELIST:           {"Client:NationIpWhitelist:", "IP 白名单"},
	WITHDRAW_WHITELIST:            {"Client:submitWithdrawWhitelist:", "提现白名单"},
	CLIENT_SMS_CODE:               {"Client:Sms:", "代理验证码"},
	SMART_START_STOP:              {"SmartStartStop:", "智能启停（最近1m数据稳定或停止启停资产）"},
	RECHARGE_DELAY_QUEUE:          {"Recharge:DelayQueue", "支付延迟队列"},
	RECHARGE_READY_QUEUE:          {"Recharge:ReadyQueue", "支付立即执行队列"},
	WITHDRAW_DELAY_QUEUE:          {"Withdraw:DelayQueue", "提现延迟队列"},
	WITHDRAW_READY_QUEUE:          {"Withdraw:ReadyQueue", "提现立即执行队列"},
	RECHARGE_ORDER_RETRY_TIMES:    {"Recharge:RetryTimes:", "支付订单重试次数"},
	WITHDRAW_ORDER_RETRY_TIMES:    {"Withdraw:RetryTimes:", "提现订单重试次数"},
	IOS_VEST_ACCOUNT:              {"Client:iosVestAccount:", "ios特定马甲账号"},
	LANGUAGE_CODE:                 {"System:languageCode:", "获取语言编码"},
	SERVICE_IP:                    {"System:IP:", "服务器IP"},
	SLOW_SQL_TIME_OUT:             {"System:SlowSqlTimeOut", "慢SQL超时时间限制 毫秒"},
	SLOW_MAPPER_ID_SQL:            {"System:SlowMapperIdSql", "慢SQL超时时间限制 毫秒"},
	ONLINE_USERS:                  {"onlineUsers", "在线用户"},
	INVITE_SUCCESSFUL_USERS:       {"System:Invite:", "邀请成功用户"},
	EMAIL_SCORE:                   {"Email:config:userNumber:", "EMAIL_SCORE"},
	LOGOUT_TOKEN_REASON:           {"Logout:Token:Reason:", "用户token过期原因"},
	ONLY_ONE_SIGN_KEY:             {"Sign:KeyOne:", "用户签到时间锁"},
	RECHARGE_SUCCESS_QUEUE:        {"Toast:Queue:", "充值成功通知消息队列"},
	NOTICE_ORDER_RESULT_QUEUE:     {"Notice:Order:Result:Queue:", "推送订单结果通知"},
	THIS_WEEK_INVIT_TOP_ONE_INFO:  {"Invit:Top1:Info:", "本周邀请用户排行榜信息"},
	REGISTRATION_RULE_JUDGMENT:    {"registrationRuleJudgment:Num", "注册设备规则限制的数量"},
	SYS_ORDER_ID:                  {"Sys:OrderNumber", "自动递增订单号，利用redis来实现"},
	SYS_ORDER_ID_QUEUE:            {"Sys:OrderNumber:queueList", "队列桶，里面是生成的订单号，没次快用完了，就批量生成20个进去"},
	AllowKlineSource:              {"AllowKlineSource:1", "允许的数据源"},
	KlineSourceStatus:             {"status:", "数据源状态"},
	KlineSource:                   {"kline_source:", "当前的数据源"},
	EMAILLIMITONE:                 {"User:Email:ExchangeList:", "邮箱更换限制次数"},
}

// KeyInfo 根据键代码获取键值
func KeyInfo(key Code) RedisServiceKeyEnum {
	return KeyMap[key]
}
