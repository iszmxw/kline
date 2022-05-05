package request

import (
	"net/http"
	"strings"
)

func Post(url string, Header http.Header, body string) (map[string]interface{}, http.Header, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		return nil, nil, err
	}
	req.Header = Header
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	// 解析Response
	return ParseResponse(response)

}
