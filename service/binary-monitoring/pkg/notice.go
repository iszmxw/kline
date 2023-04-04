package pkg

import (
	"fmt"
	"goapi/pkg/larkbot"
	"goapi/pkg/logger"
)

func Notice(text ...interface{}) {
	go func() {
		err := larkbot.
			NewClient(larkbot.LarkNeedle).
			SendMessage(fmt.Sprint(text...))
		if err != nil {
			logger.Error(err)
			return
		}
	}()
}

func NoticeLanguage(text ...interface{}) {
	go func() {
		err := larkbot.
			NewClient(larkbot.LarkLanguage).
			SendMessage(fmt.Sprint(text...))
		if err != nil {
			logger.Error(err)
			return
		}
	}()
}
