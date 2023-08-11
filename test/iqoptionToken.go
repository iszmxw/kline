package main

import (
	"fmt"
	"goapi/service/kline-iqoption/iqoption/api"
	"math/rand"
	"time"
)

func main() {
	credentials := []api.Credentials{
		{Username: "iszmxw@yopmail.com", Password: "y9cmiRA8dN6XvBt"},
		{Username: "iqoptionsmo@yopmail.com", Password: "Snz_WK4BifM/_7W"},
		{Username: "iqoptionsmt@yopmail.com", Password: "m!Q5zp2yxh%.RDV"},
		{Username: "iqoptionios@yopmail.com", Password: "6y*Z@2@YuSN!G2G"},
		{Username: "iqoptiontest@yopmail.com", Password: "C%9CFgz4cV#Qqu%"},
	}
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(credentials))
	randomCredentials := credentials[randomIndex]
	fmt.Println(randomCredentials)
	token := api.GetIqToken(randomCredentials.Username, randomCredentials.Password)
	fmt.Println(token)
}
