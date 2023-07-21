package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

}

func delAdId(adId string) {
	url := "https://api.adjust.com/device_service/api/v1/forget_device"
	method := "POST"
	payload := strings.NewReader("app_token=tiwyto5q38jk&adid=" + adId)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Token token=ygMaBirkbgfB97ADsFMn")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
