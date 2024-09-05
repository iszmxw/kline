package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	apiURL      = "https://api.binomo.com/passport/v2/sign_in?locale=en"
	methodPost  = "POST"
	contentType = "application/json"
)

var defaultHeaders = map[string]string{
	"accept-language":    "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7,vi;q=0.6",
	"cache-control":      "no-cache, no-store, must-revalidate",
	"device-type":        "web",
	"origin":             "https://binomo.com",
	"pragma":             "no-cache",
	"priority":           "u=1, i",
	"referer":            "https://binomo.com/",
	"sec-ch-ua":          `"Chromium";v="128", "Not;A=Brand";v="24", "Google Chrome";v="128"`,
	"sec-ch-ua-mobile":   "?0",
	"sec-ch-ua-platform": `"macOS"`,
	"sec-fetch-dest":     "empty",
	"sec-fetch-mode":     "cors",
	"sec-fetch-site":     "same-site",
	"user-agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36",
	"user-timezone":      "Asia/Shanghai",
	"Host":               "api.binomo.com",
	"Connection":         "keep-alive",
	"content-type":       contentType,
	"Cookie":             "__exponea_etc__=999e8a3b-ba06-480a-bc53-54631b486281; _gcl_au=1.1.1286575530.1724145099; _ga=GA1.1.1608523047.1724145100; _fbp=fb.1.1724145100540.33319079185029721; a=567ea2283542; ac=email_confirmation; sa=deposit_coupon; device_type=web; device_id=a5ecdabfa0fcd473edcefae02b4013c4; _ga_6TNVMCB3BC=GS1.1.1724148745.2.1.1724149229.60.0.0",
}

type Resp struct {
	Data struct {
		Authtoken string `json:"authtoken"`
		UserId    string `json:"user_id"`
	} `json:"data"`
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Token(email, password, deviceId string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := &http.Client{}
	req, err := createRequest(ctx, email, password, deviceId)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer res.Body.Close()

	return parseResponse(res)
}

func createRequest(ctx context.Context, email, password, deviceId string) (*http.Request, error) {
	payloadData := LoginPayload{
		Email:    email,
		Password: password,
	}

	payloadBytes, err := json.Marshal(payloadData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, methodPost, apiURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Add("device-id", deviceId)
	for key, value := range defaultHeaders {
		req.Header.Add(key, value)
	}

	return req, nil
}

func parseResponse(res *http.Response) (string, error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var response Resp
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return response.Data.Authtoken, nil
}
