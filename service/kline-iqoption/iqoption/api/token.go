package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type IQToken struct {
	Code string `json:"code"`
	Ssid string `json:"ssid"`
}

type Credentials struct {
	Username string
	Password string
}

func GetIqToken(username, password string) string {
	var IqToken IQToken
	url := "https://api.iqoption.com/v2/login"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{
    "identifier": "%s",
    "password": "%s"
}`, username, password))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("content-type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	err = json.Unmarshal(body, &IqToken)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return IqToken.Ssid
}
