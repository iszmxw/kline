package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	chttp := new(http.Client)
	for true {
		go curl(chttp)
		time.Sleep(time.Second * 8)
	}
}

func curl(chttp *http.Client) {
	get, err := chttp.Get("http://47.242.148.57:7104/app/v1/trade/maintenanceTableData.json?num=14")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(get)
}
