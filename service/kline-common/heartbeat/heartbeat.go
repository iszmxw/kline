package heartbeat

import (
	"errors"
	"fmt"
	"goapi/app/models"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
	"goapi/service/kline-common"
	"goapi/service/kline-common/source"
	"strings"
	"time"
)

// 心跳 heartbeat 健康

func Start(num time.Duration, Source string, CodeData []KlineCommon.Info) {
	logger.Info("等待当前数据订阅5分钟后开始心跳检测")
	time.Sleep(time.Minute * 5)
	ticker := time.NewTicker(num * time.Second) // 每隔 num 秒检查一次
	for range ticker.C {
		for _, info := range CodeData {
			if strings.EqualFold(Source, "iqoption") {
				// 进行新的检测方法来检测
				CheckListData(Source, info)
			} else {
				CheckOutStatus(Source, info)
			}
		}
	}
}

func CheckOutStatus(Source string, info KlineCommon.Info) {
	status := source.CheckStatus(Source, info.ContractCode)
	logger.Info(fmt.Sprintf("数据源【%v】对应资产【%v】健康状态：【%v】", Source, info.ContractCode, status))
}

// 检测列表数据

func CheckListData(Source string, info KlineCommon.Info) {
	health := "4" // 警告，延迟过高
	var KlineData models.KlineLog
	err := mysql.DB.Debug().
		Model(models.KlineLog{}).
		Select("log_id,ts").
		Where("kline_source", Source).
		Where("contract_code", info.ContractCode).
		Order("ts desc").
		Limit(1).
		Find(&KlineData).Error
	if err != nil {
		logger.Error(err)
	}
	KeyStatus := fmt.Sprintf("status:%v:market.%v.index.1min", Source, info.ContractCode)
	if KlineData.ID > 0 {
		// 差
		SystemTs := time.Now().UnixMilli()
		diff := (SystemTs - KlineData.Ts) / 1000
		switch true {
		case diff < 60*1:
			health = "1" // 正常，绿色
			break
		case (diff >= 60) && (diff < 60*2):
			health = "2" // 延迟，黄色
			break
		case (diff >= 60*2) && (diff < 60*3):
			health = "3" // 警告，延迟过高
			break
		case diff >= 60*3:
			health = "4" // 警告，延迟过高
			break
		}
	} else {
		logger.Error(errors.New("查询数据为空"))
	}
	_, _ = redis.Client.Add(KeyStatus, health, 0) // 警告，延迟过高
}
