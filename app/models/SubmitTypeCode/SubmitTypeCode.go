package SubmitTypeCode

// Code 枚举类型
type Code int

const (
	MobileCredit Code = iota
	EmailCredit
	IdentityCredit
	BankCredit
	InvitCredit
)

// SubmitTypeCode 存储键值对应关系和备注信息的结构体
type SubmitTypeCode struct {
	Code  int    // 键值
	Alias string // 备注信息
}

// Map 存储键值对应关系和备注信息的映射
var Map = map[Code]SubmitTypeCode{
	MobileCredit:   {1, "验证手机号"},
	EmailCredit:    {2, "验证邮箱"},
	IdentityCredit: {3, "验证证件"},
	BankCredit:     {4, "验证银行卡"},
	InvitCredit:    {5, "提交邀请码"},
}

// Info 根据键代码获取键值
func Info(Code Code) SubmitTypeCode {
	return Map[Code]
}
