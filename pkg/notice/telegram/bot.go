package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	conf "goapi/pkg/config"
	"goapi/pkg/logger"
	"strings"
)

const (
	TOKEN      = "5202189634:AAHdLadirD9tujR-djCNbyAyiohJGiWFvAU"
	chatIDPro  = -1001505716949     // 正式环境房间
	chatIDTest = -1001786874645 // 测试环境房间
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
	if conf.GetString("app.env") == "company_pro" { // 正式环境监控
		msg := tgbotapi.NewMessage(chatIDPro, text)
		_, err = bot.Send(msg)
		if err != nil {
			logger.Error(err)
			return
		}
	} else {
		msg := tgbotapi.NewMessage(chatIDPro, text)
		_, err = bot.Send(msg)
		if err != nil {
			logger.Error(err)
			return
		}
	}
}
