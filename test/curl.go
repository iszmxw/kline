package main

import (
	"fmt"
	"goapi/pkg/helpers"
)

func main() {
	var data = []string{
		"47.242.148.57",
		"183.237",
		"183.240.8.89",
		"103.144",
		"183.240",
		"240e:3ba:3022:1600:f969:ea8:399c:7ecf",
	}
	if helpers.InArray("240e:3ba:3022:1600:f969:ea8:399c:7ecf", data) {
		fmt.Println("111")
		fmt.Println("111")
		fmt.Println("111")
		fmt.Println("111")
	}
}
