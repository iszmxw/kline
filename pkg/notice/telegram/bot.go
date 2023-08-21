package telegram

import (
	"fmt"
	conf "goapi/pkg/config"
	"goapi/pkg/larkbot"
	"goapi/pkg/logger"
	"goapi/pkg/redis"
	"goapi/pkg/utils"
	"strings"
)

const (
	TOKEN      = "5202189634:AAHdLadirD9tujR-djCNbyAyiohJGiWFvAU"
	chatIDPro  = -1001505716949 // 正式环境房间
	chatIDTest = -1001786874645 // 测试环境房间
)

func Notice(info map[string]interface{}) {
	key := utils.Md5(fmt.Sprintf("%s%s", info["title"], info["content"]))
	ok, _ := redis.Client.SelectDbSetNx(conf.GetInt("redis.db"), fmt.Sprintf("lock:errNotice:%s", key), key, 60*3)
	if !ok {
		return
	}
	text := strings.ReplaceAll(fmt.Sprintf(
		`
【标题】：
%v

【内容】：
%v
`, info["title"], info["content"]), "<br>", "")
	// 开启 debug
	if conf.GetString("app.env") == "company_pro" { // 正式环境监控
		err := larkbot.
			NewClient(larkbot.LarkIDPro).
			SendMessage(text)
		if err != nil {
			logger.Error(err)
			return
		}
	} else {
		err := larkbot.
			NewClient(larkbot.LarkIDTest).
			SendMessage(text)
		if err != nil {
			logger.Error(err)
			return
		}
	}
}
