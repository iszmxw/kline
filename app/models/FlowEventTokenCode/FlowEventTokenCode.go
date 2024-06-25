package FlowEventTokenCode

// Code 枚举类型
type Code int

const (
	Register Code = iota
	FirstPurchase
	Purchase
	Trade
	FirstWithdraw
	Withdraw
	CompleteTutorial
	SkipTutorial
	SupplementarySignature
)

// FlowEventTokenCode 存储键值对应关系和备注信息的结构体
type FlowEventTokenCode struct {
	Code  string // 键值
	Alias string // 备注信息
}

// Map 存储键值对应关系和备注信息的映射
var Map = map[Code]FlowEventTokenCode{
	Register:               {"register", "注册"},
	FirstPurchase:          {"firstPurchase", "首次充值"},
	Purchase:               {"purchase", "充值"},
	Trade:                  {"trade", "首次IDR下单"},
	FirstWithdraw:          {"firstWithdraw", "首次提现"},
	Withdraw:               {"withdraw", "提现"},
	CompleteTutorial:       {"completeTutorial", "完成教程"},
	SkipTutorial:           {"skipTutorial", "跳过教程"},
	SupplementarySignature: {"supplementarySignature", "补签签到"},
}

// Info 根据键代码获取键值
func Info(Code Code) FlowEventTokenCode {
	return Map[Code]
}
