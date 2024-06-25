package BonusActivityTypeCode

// Code 枚举类型
type Code int

const (
	// 定义奖金活动类型的常量
	FirstRecharge Code = iota
	RegistFirst
	TransactionLoss
	InvitePolitely
	SupplementarySignature
	RankReward
)

// BonusActivityTypeCode 存储键值对应关系和备注信息的结构体
type BonusActivityTypeCode struct {
	Code   string // 键值
	Remark string // 备注信息
}

// KeyMap 存储键值对应关系和备注信息的映射
var KeyMap = map[Code]BonusActivityTypeCode{
	FirstRecharge:          {"first_recharge", "Bonus Hadiah Isi Ulang Pertama."},
	RegistFirst:            {"regist_first", "Bonus hadiah pendaftaran."},
	TransactionLoss:        {"transaction_loss", "Kerugian kembali bonus hadiah."},
	InvitePolitely:         {"invite_politely", "Undang Bonus Hadiah Transaksi."},
	SupplementarySignature: {"supplementary_signature", "Masuk ulang dan bonus masuk."},
	RankReward:             {"rank_reward", "Jumlah bonus peringkat trading"},
}

// KeyInfo 根据键代码获取键值
func KeyInfo(code Code) BonusActivityTypeCode {
	return KeyMap[code]
}
