package heartbeat

import (
	"fmt"
	"goapi/pkg/logger"
	"goapi/pkg/redis"
	"goapi/service/kline-common"
	"goapi/service/kline-common/source"
	"goapi/service/kline-iqoption/iqoption/api"
	kconfig "goapi/service/kline-iqoption/iqoption/config"
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
	// iqoption 新的数据源活跃度检测方法
	floatList := api.GetListData(kconfig.ListKey(info.ContractCode))
	fmt.Println(floatList)
	if len(floatList) < 100 {
		// 少于100个数据点，就不进行检测了
		return
	}
	firstNum := floatList[0]
	// 健康
	health := false
	for _, item := range floatList {
		if item != firstNum {
			health = true
			break
		}
	}
	KeyStatus := fmt.Sprintf("status:%v:market.%v.index.1min", Source, info.ContractCode)
	// 不健康
	if health == false {
		_, _ = redis.Client.Add(KeyStatus, "4", 0) // 警告，延迟过高
	} else {
		_, _ = redis.Client.Add(KeyStatus, "1", 0) // 正常，绿色
	}
	// 大于 3000个点的话也要删除掉
	if len(floatList) > 3000 {
		redis.Client.Delete(kconfig.ListKey(info.ContractCode))
	}
	// 第二波检测
	//CheckOutStatus(Source, info)
}
