package mine

import (
	"math/rand"
	"time"
)

// GetDummyBonus 计算本周虚拟用户邀请赠送奖金
func GetDummyBonus(bonusAmount, thisWeekBonus float64) float64 {
	var bonusAmounts float64
	dayOfWeek := time.Now().Weekday()

	switch dayOfWeek {
	// 周五至周日，每日随机增长0-1
	case time.Friday, time.Saturday, time.Sunday:
		randomNumber := rand.Intn(2)
		bonusAmounts = float64(randomNumber)*bonusAmount + thisWeekBonus
	// 周一默认为1个
	case time.Monday:
	// 周二为2个
	case time.Tuesday:
	// 周四为3人
	case time.Thursday:
		bonusAmounts = bonusAmount + thisWeekBonus
	default:
		bonusAmounts = thisWeekBonus
	}

	return bonusAmounts
}
