package main

import (
	"fmt"
	"github.com/opesun/goquery"
	"strings"
)

func main() {
	x, err := goquery.ParseUrl("https://www.xe.com/zh-CN/currencyconverter/convert/?Amount=1&From=USD&To=INR")
	if err != nil {
		panic(err)
	}
	result := x.Find(".result__BigRate-sc-1bsijpp-1.iGrAod").Text()
	if len(result) > 0 {
		arr := strings.Split(result, " ")
		fmt.Println(arr[0])
	}
}
