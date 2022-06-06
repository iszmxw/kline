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
	get, err := chttp.Get("http://8.215.38.97:8080/app/v1/trade/maintenanceTableData.json?num=30")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(get)
}
