package main

import (
	"fmt"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone
}

func main() {
	fmt.Println(time.Unix(1647870090, 0).Format("2006-01-02 15:04:05"))
}
