package larkbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goapi/pkg/utils"
	"net/http"
)

const (
	LarkIDTest          = "8f95cc4a-72bc-472a-93e3-3584975675ed" // 测试环境房间
	LarkIDPro           = "40ea8c9f-1fa5-4549-8359-bb4aeb311575" // 测试环境房间
	LarkNeedle          = "319b49a2-7167-48fd-8167-616adc851ff6" // 插针ID
	LarkLanguage        = "9dd8ccce-8e1f-4c1b-abc5-49ef64a00dc5" // 测试环境房间
	LarkServeMonitoring = "c83007c5-d8b3-4104-bbd2-d47aa9cac8e3" // 服务监控机器人
)

type Client struct {
	WebhookURL string
}

type Message struct {
	MsgType string `json:"msg_type"`
	Content struct {
		Text string `json:"text"`
	} `json:"content"`
}

func NewClient(token string) *Client {
	return &Client{
		WebhookURL: fmt.Sprintf("https://open.larksuite.com/open-apis/bot/v2/hook/%s", token),
	}
}

func (c *Client) SendMessage(text string) error {
	Serve := fmt.Sprintf("当前的服务器IP：【%s】\n\n\n", utils.ServeIp())
	text = Serve + text
	msg := &Message{
		MsgType: "text",
		Content: struct {
			Text string `json:"text"`
		}(struct{ Text string }{Text: text}),
	}
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	resp, err := http.Post(c.WebhookURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("send message failed with status code: %d", resp.StatusCode)
	}
	return nil
}
