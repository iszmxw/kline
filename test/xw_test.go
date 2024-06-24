package test

import (
	"fmt"
	"goapi/pkg/decimal"
	"testing"
)

func Test_xw(t *testing.T) {
	res := decimal.SplitDisplayInteger(123123456789.88, "Indonesia")
	fmt.Println(res)
}
