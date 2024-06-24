package decimal

import (
	"fmt"
	"github.com/leekchan/accounting"
	"math"
	"strconv"
	"strings"
)

func FormatDouble(userTradingRank float64, numberPoint int) string {
	// 四舍五入保留指定小数位数
	scaleFactor := math.Pow10(numberPoint)
	formattedValue := math.Round(userTradingRank*scaleFactor) / scaleFactor

	// 将结果转换为字符串
	result := strconv.FormatFloat(formattedValue, 'f', numberPoint, 64)

	return result
}

// 格式化整数部分，每三位用"."分隔
func formatInt(num int, sep string) string {
	str := fmt.Sprintf("%d", num)
	var formatted []string
	for len(str) > 3 {
		formatted = append([]string{str[len(str)-3:]}, formatted...)
		str = str[:len(str)-3]
	}
	if len(str) > 0 {
		formatted = append([]string{str}, formatted...)
	}
	return strings.Join(formatted, sep)
}

// 格式化小数部分，保留指定位数的小数并向下取整
func formatDecimal(decimalPart float64, decimalPlaces int) string {
	if decimalPlaces <= 0 {
		return ""
	}
	// 将小数部分乘以10的指定位数次方，然后取整，再转成字符串
	integerPart := int(math.Floor(decimalPart * math.Pow(10, float64(decimalPlaces))))
	return fmt.Sprintf("%0*d", decimalPlaces, integerPart)
}

func PreserveDecimalPlaces(value float64, numberPoint int) string {
	factor := math.Pow(10, float64(numberPoint))
	truncatedValue := math.Floor(value*factor) / factor
	res := fmt.Sprintf("%.*f", numberPoint, truncatedValue)
	return res
}

func SplitDisplay(num float64, nation string, decimalPlaces int) string {
	// 分离整数部分和小数部分
	intPart := int(num)
	decimalPart := num - float64(intPart)
	sep := ","
	if nation == "Indonesia" {
		sep = "."
	}

	// 格式化整数部分
	intPartStr := formatInt(intPart, sep)

	// 格式化小数部分
	decimalStr := formatDecimal(decimalPart, decimalPlaces)
	if len(decimalStr) == 0 {
		return intPartStr
	}
	// 拼接结果
	return intPartStr + "." + decimalStr
}

func SplitDisplayInteger(num float64, nation string) string {
	var ac accounting.Accounting
	if nation == "India" {
		ac = accounting.Accounting{
			Symbol:    "",
			Precision: 0,
			Thousand:  ",",
			Decimal:   ".",
		}
	}
	if nation == "Indonesia" {
		ac = accounting.Accounting{
			Symbol:    "",
			Precision: 0,
			Thousand:  ".",
			Decimal:   ".",
		}
	}
	// 使用 accounting 进行格式化
	formattedNumber := ac.FormatMoney(num)
	return formattedNumber
}
