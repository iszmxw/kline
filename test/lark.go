package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"goapi/pkg/helpers"
	"goapi/service/binary-monitoring/pkg"
)

func main() {

	Rand := pkg.RandFloat(0, 0.00008)

	res := decimal.NewFromFloat(0.4).Sub(decimal.NewFromFloat(0.2).Mul(decimal.NewFromFloat(0.5))).InexactFloat64()

	fromString, err := decimal.NewFromString(fmt.Sprintf("%v", helpers.Decimal(Rand, 6)))
	if err != nil {
		return
	}

	fmt.Println(fromString, res)
}
