package pkg

import (
	"fmt"
	"goapi/pkg/larkbot"
	"goapi/pkg/logger"
)

func Notice(text ...interface{}) {
	go func() {
		err := larkbot.
			NewClient("319b49a2-7167-48fd-8167-616adc851ff6").
			SendMessage(fmt.Sprint(text...))
		if err != nil {
			logger.Error(err)
			return
		}
	}()
}
