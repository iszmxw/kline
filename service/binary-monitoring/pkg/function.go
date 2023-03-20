package pkg

import (
	"github.com/shopspring/decimal"
	"math/rand"
)

func Sub(cClose, VolatilityAmplitude, Rand float64) float64 {
	num := decimal.NewFromFloat(cClose).Sub( // 震荡幅度为0到VolatilityAmplitude之间
		decimal.NewFromFloat(Rand).Mul(
			decimal.NewFromFloat(VolatilityAmplitude),
		),
	).InexactFloat64()
	return num
}

func RandFloat(min, max float64) float64 {
	return rand.Float64()*(max-min) + min // 生成0到0.0345之间的随机数
}

func Add(cClose, VolatilityAmplitude, Rand float64) float64 {
	num := decimal.NewFromFloat(cClose).Add( // 震荡幅度为0到VolatilityAmplitude之间
		decimal.NewFromFloat(Rand).Mul(
			decimal.NewFromFloat(VolatilityAmplitude),
		),
	).InexactFloat64()
	return num
}
