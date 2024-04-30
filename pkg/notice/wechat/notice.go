package wechat

import (
	"fmt"
	"goapi/pkg/logger"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func Notice(title, bodys string) {
	// 组装模板消息
	msg := title + "：\n\n"
	msg += fmt.Sprintf("当前时间：%v\n\n", time.Now().Format(time.DateTime))
	msg += fmt.Sprintf("详情：\n\n%v", bodys)
	go SendMsg("", msg)
}

// 发送微信消息通知

func SendMsg(id, msg string) {
	params := url.Values{}
	if len(id) > 0 {
		params.Set("id", id)
	}
	params.Set("msg", msg)
	encodedParams := params.Encode()
	baseURL := "https://wx-bot.vercel.app/wechat/v1/bot/SendMsg.json?bot_id=1"
	fullURL := baseURL + "&" + encodedParams
	res, err := http.Get(fullURL)
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	logger.Info(string(body))
}
