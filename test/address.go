package main

import (
	"fmt"
	"goapi/pkg/address"
)

func main() {
	fmt.Println(address.IpDomainByIp("183.237.196.214:8080", 2))
	//if strings.Contains("183.237.196.214", "47.242.148.57") {
	//	fmt.Println(123456)
	//}
}
