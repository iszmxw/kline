package utils_socket

import (
	"goapi/pkg/logger"
	"golang.org/x/net/websocket"
	"net/url"
)

type Client struct {
	Host string
	Path string
}

func NewWebsocketClient(host, path string) *Client {
	return &Client{
		Host: host,
		Path: path,
	}
}

func (this *Client) SendMessage(body []byte) error {
	u := url.URL{Scheme: "ws", Host: this.Host, Path: this.Path}
	ws, err := websocket.Dial(u.String(), "", "http://"+this.Host+"/")
	defer ws.Close() //关闭连接
	if err != nil {
		logger.Error(err)
		return err
	}

	_, err = ws.Write(body)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
