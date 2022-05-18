package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"goapi/pkg/logger"
	"strings"
)

const (
	TOKEN  = "5202189634:AAHdLadirD9tujR-djCNbyAyiohJGiWFvAU"
	chatID = -685844165
)

func Notice(info map[string]interface{}) {
	text := strings.ReplaceAll(fmt.Sprintf(
		`
【标题】：
%v

【内容】：
%v
`, info["title"], info["content"]), "<br>", "")
	bot, err := tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		logger.Error(err)
		return
	}
	// 开启 debug
	//bot.Debug = true
	logger.Info(fmt.Sprintf("【%s】登录成功", bot.Self.UserName))
	msg := tgbotapi.NewMessage(chatID, text)
	send, err := bot.Send(msg)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(send)
}
