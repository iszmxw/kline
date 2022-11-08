package main

import (
	"fmt"
	"goapi/pkg/xtime"
	"time"
)

func main() {
	SystemTs := time.Now().UnixMilli()
	minute := time.Unix(0, SystemTs*int64(time.Millisecond)).Format("2006-01-02 15:04")
	fmt.Println(minute, xtime.Parse(fmt.Sprintf("%v:00", minute)).UnixMilli())
}
