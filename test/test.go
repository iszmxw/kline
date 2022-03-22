package main

import (
	"fmt"
	"strings"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone
}

func main() {
	fmt.Println(time.Unix(1647870090, 0).Format("2006-01-02 15:04:05"))
	fmt.Println(strings.Replace("BTC-USD-uasasas-asasas", "-", "/", -1)) // 合约代码 BTC-USD
}
