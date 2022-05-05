package request

import (
	"net/http"
)

func Get(url string, Header http.Header) (map[string]interface{}, http.Header, error) {
	req, err := http.NewRequest("GET", url, nil)
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
