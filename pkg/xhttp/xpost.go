package xhttp

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
)

func FormData(url string, params map[string]string) ([]byte, error) {
	// 创建一个缓冲区和multipart写入器
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	// 添加参数字段
	for key, value := range params {
		writer.WriteField(key, value)
	}
	// 关闭multipart写入器
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	// 创建请求
	req, err := http.NewRequest("POST", url, &buffer)
	if err != nil {
		return nil, err
	}

	// 设置Content-Type头部
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
