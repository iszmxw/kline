package larkbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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
