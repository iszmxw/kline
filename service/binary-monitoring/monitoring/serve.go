package monitoring

import (
	"encoding/json"
	"errors"
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
	"goapi/service/binary-monitoring/queue"
	"math"
	"math/rand"
	"strings"
	"time"
)

type Params struct {
	// 最后交易的时间节点
	LastTradeBeginTimestamp int64
	// 买涨的人数
	upNum int64
	// 买涨的金额
	upAmounts float64
	// 买跌的人数
	downNum int64
	// 买跌的金额
	downAmounts float64
	// HClose 最大 close
	HClose float64
	// LClose 最小 close
	LClose float64
	// avgClose 最近一分钟平均 close
	avgClose float64
	// diffClose 差异 close
	diffClose float64
	winLoss   float64
	WinType   string
	Peoples   []string
	OrderIds  []string
	config    response.TradingAssetsMonitoringConfig
}

// 监控

func Start() {
	// 1、查询需要插针的资产
	var list []response.TradingAssetsMonitoring
	mysql.DB.
		Model(models.TradingAssets{}).
		Where(models.Prefix("$_trading_assets.is_valid"), 1).
		Where(models.Prefix("$_trading_assets.is_monitoring"), 1).
		Select(models.Prefix("$_trading_assets.id,$_trading_assets.assets_code,$_trading_assets.point_number,$_trading_assets.monitoring_number,$_trading_assets.is_monitoring,$_trading_assets.shock_amplitude,$_trading_assets.volatility_amplitude,$_trading_assets.volatility_frequency,$_system_data_dictionary.value as assets_insert_frame")).
		Joins(models.Prefix("right join $_system_data_dictionary on $_system_data_dictionary.name='SMO_ASSETS_INSERT_FRAME'")).
		Find(&list)
	for _, item := range list {
		OK, _ := redis.Client.Get(fmt.Sprintf("Monitoring:%s", item.AssetsCode))
		if len(OK) > 0 {
			logger.Info(OK)
			continue
		}
		var AssetsInsertFrame response.AssetsInsertFrame
		err := json.Unmarshal([]byte(item.AssetsInsertFrame), &AssetsInsertFrame)
		if err != nil {
			logger.Info(err)
			return
		}
		currentTimeMillis := time.Now().UnixMilli()
		if 60-AssetsInsertFrame.InAdvanceTimer <= 0 {
			// 设置提前触发时间,如果设置不合理，就跳过，同时打印日志，或者发送通知
			logger.Info(currentTimeMillis, "设置提前触发时间不合理")
			continue
		}
		if 60-AssetsInsertFrame.EarlyExecutionTimer <= 10 {
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
		if AssetsInsertFrame.AreaTimer > AssetsInsertFrame.IntervalTimer {
			logger.Info(currentTimeMillis, "区域时长大于间隔触发时间")
			continue
		}
		// 小数点波动位数
		if item.MonitoringNumber == 0 {
			logger.Info(currentTimeMillis, "请设置小数点波动位数")
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
			Where(models.Prefix("$_trading_order.trade_end_timestamp >= ?"), currentTimeMillis-(AssetsInsertFrame.InAdvanceTimer+AssetsInsertFrame.EarlyExecutionTimer+int64(AssetsInsertFrame.EarlyExecutionFrequency))*1000).
			Select(models.Prefix("$_trading_order.*,$_app_currency_rate.b_num as currency_rate")).
			Joins(models.Prefix("right join $_app_currency_rate on $_trading_order.currency_type=$_app_currency_rate.currency_b")).
			Find(&tradingOrder)
		num := len(tradingOrder)
		//logger.Info(num)
		if num > 0 {
			go new(Params).Calculate(currentTimeMillis-int64(AssetsInsertFrame.EarlyExecutionFrequency)*1000, item, tradingOrder)
		}
	}
}

// 初始化参数

func (p *Params) Init(config response.TradingAssetsMonitoring) bool {
	// 最后交易的时间节点
	p.LastTradeBeginTimestamp = 0
	// 买涨的人数
	p.upNum = 0
	// 买涨的金额
	p.upAmounts = 0.00
	// 买跌的人数
	p.downNum = 0
	// 买跌的金额
	p.downAmounts = 0.00
	// HClose 最大 close
	p.HClose = 1000000
	// LClose 最小 close
	p.LClose = 0.00
	p.winLoss = 0.00
	p.WinType = ""
	p.Peoples = make([]string, 0)
	p.OrderIds = make([]string, 0)
	// 插针配置
	var AssetsInsertFrame response.AssetsInsertFrame
	err := json.Unmarshal([]byte(config.AssetsInsertFrame), &AssetsInsertFrame)
	if err != nil {
		logger.Info(err)
		return false
	}
	p.config = response.TradingAssetsMonitoringConfig{
		ID:                  config.ID,
		AssetsCode:          config.AssetsCode,
		PointNumber:         config.PointNumber,
		MonitoringNumber:    config.MonitoringNumber,
		IsMonitoring:        config.IsMonitoring,
		ShockAmplitude:      config.ShockAmplitude,
		VolatilityAmplitude: config.VolatilityAmplitude,
		VolatilityFrequency: config.VolatilityFrequency,
		AssetsInsertFrame: response.AssetsInsertFrame{
			AreaTimer:               AssetsInsertFrame.AreaTimer,
			IntervalTimer:           AssetsInsertFrame.IntervalTimer,
			InAdvanceTimer:          AssetsInsertFrame.InAdvanceTimer,
			LossThreshold:           AssetsInsertFrame.LossThreshold,
			EarlyExecutionTimer:     AssetsInsertFrame.EarlyExecutionTimer,
			EarlyExecutionFrequency: AssetsInsertFrame.EarlyExecutionFrequency,
			RecoverTimer:            AssetsInsertFrame.RecoverTimer,
			RecoverFrequency:        AssetsInsertFrame.RecoverFrequency,
		},
	}
	return true
}

// 进行计算分析是否需要插针，以及按照规则生成插针数据

func (p *Params) Calculate(currentTimeMillis int64, conf response.TradingAssetsMonitoring, orders []response.TradingOrder) {
	if !p.Init(conf) {
		logger.Info("初始化参数失败")
		return
	}
	if len(orders) == 1 {
		order := orders[0]
		p.Peoples = append(p.Peoples, order.AppUserID)
		p.OrderIds = append(p.OrderIds, order.TradeOrderID)
		if strings.EqualFold("up", order.UpDownType) {
			p.upNum++
			// 换算美元
			p.upAmounts = decimal.NewFromFloat(p.upAmounts).Add(
				decimal.NewFromFloat(order.Amounts).Mul(decimal.NewFromFloat(order.ReturnRate)).
					Div(decimal.NewFromFloat(order.CurrencyRate)),
			).InexactFloat64()
		}
		if strings.EqualFold("down", order.UpDownType) {
			p.downNum++
			// 换算美元
			p.downAmounts = decimal.NewFromFloat(p.downAmounts).Add(
				decimal.NewFromFloat(order.Amounts).Mul(decimal.NewFromFloat(order.ReturnRate)).
					Div(decimal.NewFromFloat(order.CurrencyRate)),
			).InexactFloat64()
		}
		p.HClose = order.BeginAssetsPrice
		p.LClose = order.BeginAssetsPrice
		p.LastTradeBeginTimestamp = order.TradeBeginTimestamp
	} else {
		for _, order := range orders {
			p.Peoples = append(p.Peoples, order.AppUserID)
			p.OrderIds = append(p.OrderIds, order.TradeOrderID)
			p.HClose = order.BeginAssetsPrice
			p.LClose = order.BeginAssetsPrice
			if strings.EqualFold("up", order.UpDownType) {
				p.upNum++
				// 换算美元
				p.upAmounts = decimal.NewFromFloat(p.upAmounts).Add(
					decimal.NewFromFloat(order.Amounts).Mul(decimal.NewFromFloat(order.ReturnRate)).
						Div(decimal.NewFromFloat(order.CurrencyRate)),
				).InexactFloat64()
			}
			if strings.EqualFold("down", order.UpDownType) {
				p.downNum++
				// 换算美元
				p.downAmounts = decimal.NewFromFloat(p.downAmounts).Add(
					decimal.NewFromFloat(order.Amounts).Mul(decimal.NewFromFloat(order.ReturnRate)).
						Div(decimal.NewFromFloat(order.CurrencyRate)),
				).InexactFloat64()
			}
			// 最大close
			if p.HClose <= order.BeginAssetsPrice {
				p.HClose = order.BeginAssetsPrice
			} else {
				logger.Info(p.HClose, order.BeginAssetsPrice)
			}
			// 最小close
			if p.LClose >= order.BeginAssetsPrice {
				p.LClose = order.BeginAssetsPrice
			} else {
				logger.Info(p.LClose, order.BeginAssetsPrice)
			}
			if order.TradeBeginTimestamp > p.LastTradeBeginTimestamp {
				p.LastTradeBeginTimestamp = order.TradeBeginTimestamp
			}
		}
	}
	// 检测是否需要进行干预，以及干预的方向
	ok, Close := p.finalResult()
	if !ok {
		return
	}
	// 检测是否安全期内
	yes, err := redis.Client.SelectDbSetNx(15, fmt.Sprintf("Monitoring:%s", p.config.AssetsCode), "suo", p.config.AssetsInsertFrame.IntervalTimer)
	if err != nil {
		return
	}
	if !yes {
		return
	}
	Close, err = helpers.FormatFloat(Close, p.config.PointNumber)
	if err != nil {
		return
	}
	if p.LClose == 0 || p.HClose == 1000000 {
		logger.Info("最大最小值初始化失败", orders)
		ok = redis.Client.Delete(fmt.Sprintf("Monitoring:%s", p.config.AssetsCode))
		logger.Info(ok)
		return
	}

	logger.Info("========================trade_end_timestamp > ", time.Unix(0, (currentTimeMillis-p.config.AssetsInsertFrame.InAdvanceTimer*1000)*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999"), "========================")
	logger.Info("========================trade_end_timestamp <= ", time.Unix(0, (currentTimeMillis+(60-p.config.AssetsInsertFrame.InAdvanceTimer)*1000)*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999"), "========================")
	logger.Info("========================", orders[0].KlineTime, "========================")
	//logger.Info("平均close", p.avgClose)
	logger.Info("最大close", p.HClose)
	logger.Info("最小close", p.LClose)
	logger.Info("diffClose--", p.diffClose)
	go queue.Consume(fmt.Sprintf("Queue:interferenceNeedle:%s", p.config.AssetsCode))
	go p.CreateInterferenceNeedle(currentTimeMillis, Close)

}

// 最终结果
func (p *Params) finalResult() (bool, float64) {
	OK, _ := redis.Client.Get(fmt.Sprintf("Monitoring:%s", p.config.AssetsCode))
	if len(OK) > 0 {
		logger.Info(OK)
		return false, 0
	}
	var avgArr map[string]interface{}
	//var floatArr []float64
	currentTimeMillis := time.Now().UnixMilli()
	mysql.DB.Model(models.KlineLog{}).
		Where("ts >= ?", currentTimeMillis-60000).
		Where("ts <= ?", currentTimeMillis).
		Where("contract_code", p.config.AssetsCode).
		Select("MAX(c_close) - MIN(c_close) as Close").
		Find(&avgArr)
	p.diffClose = helpers.StrToFloat64(string(avgArr["Close"].([]uint8)))
	p.diffClose = helpers.Decimal(p.diffClose, p.config.PointNumber)
	// 区域内用户盈利金额
	p.winLoss = decimal.NewFromFloat(p.upAmounts).Sub(decimal.NewFromFloat(p.downAmounts)).InexactFloat64()
	if p.winLoss == 0 { // 不用干预的情况：盈亏平了
		return false, 0
	}
	if p.winLoss > 0 {
		p.WinType = "winDown" // 低价位平台盈利
	}
	if p.winLoss < 0 {
		p.winLoss = -p.winLoss
		p.WinType = "winUp" // 高价位平台盈利
	}
	// 不用干预的情况：区域平台亏损金额<xx美金
	if p.winLoss < p.config.AssetsInsertFrame.LossThreshold {
		return false, 0
	}
	//res1 := math.Abs(decimal.NewFromFloat(p.HClose).Sub(decimal.NewFromFloat(p.avgClose)).InexactFloat64())
	//res2 := math.Abs(decimal.NewFromFloat(p.LClose).Sub(decimal.NewFromFloat(p.avgClose)).InexactFloat64())
	//if res1 != 0 {
	//	floatArr = append(floatArr, res1)
	//}
	//if res2 != 0 {
	//	floatArr = append(floatArr, res2)
	//}
	//if len(floatArr) == 0 {
	//	logger.Info("p.diffClose", "获取失败")
	//	return false, 0
	//}
	//p.diffClose = randSliceValue(floatArr)
	//p.diffClose = getNum(p.config.MonitoringNumber, p.avgClose)
	// 第几位在波动
	//p.config.MonitoringNumber = len(strings.Split(fmt.Sprintf("%v", p.diffClose), ".")[1])
	// 同一个时间点有人买涨，有人买跌
	if p.HClose == p.LClose {
		if strings.EqualFold("winUp", p.WinType) { // 高价位平台盈利，算出一个高价位
			//xs := helpers.MinNum(p.HClose)
			Close := p.HClose + p.diffClose*decimal.NewFromFloat(1).Add(decimal.NewFromFloat(p.config.VolatilityAmplitude)).InexactFloat64()
			//logger.Info(Close)
			return true, Close
		}
		if strings.EqualFold("winDown", p.WinType) { // 低价位平台盈利，算出一个低价位
			//xs := helpers.MinNum(p.LClose)
			Close := p.LClose - p.diffClose*decimal.NewFromFloat(1).Add(decimal.NewFromFloat(p.config.VolatilityAmplitude)).InexactFloat64()
			//logger.Info(Close)
			return true, Close
		}
	}
	if strings.EqualFold("winUp", p.WinType) { // 高价位平台盈利
		return true, p.HClose
	}
	if strings.EqualFold("winDown", p.WinType) { // 低价位平台盈利
		return true, p.LClose
	}
	if p.LastTradeBeginTimestamp == 0 {
		return false, 0
	}
	return false, 0
}

func randSliceValue(xs []float64) float64 {
	return xs[rand.Intn(len(xs))]
}

func Sub(cClose, VolatilityAmplitude, Rand float64) float64 {
	num := decimal.NewFromFloat(cClose).Sub( // 震荡幅度为0到VolatilityAmplitude之间
		decimal.NewFromFloat(Rand).Mul(
			decimal.NewFromFloat(VolatilityAmplitude),
		),
	).InexactFloat64()
	return num
}

func Rand(cClose, VolatilityAmplitude, Rand float64) float64 {
	num := decimal.NewFromFloat(cClose).Add( // 震荡幅度为-VolatilityAmplitude到VolatilityAmplitude之间
		decimal.NewFromFloat(Rand).Mul(
			decimal.NewFromFloat(VolatilityAmplitude).Sub(decimal.NewFromFloat(VolatilityAmplitude).Mul(decimal.NewFromFloat(0.5))),
		),
	).InexactFloat64()
	return num
}

func RandFloat(min, max float64) float64 {
	return rand.Float64()*(max-min) + min // 生成0到0.0345之间的随机数
}

func Add(cClose, VolatilityAmplitude, Rand float64) float64 {
	num := decimal.NewFromFloat(cClose).Add( // 震荡幅度为0到VolatilityAmplitude之间
		decimal.NewFromFloat(Rand).Mul(
			decimal.NewFromFloat(VolatilityAmplitude),
		),
	).InexactFloat64()
	return num
}

// 创建干扰针

func (p *Params) CreateInterferenceNeedle(currentTimeMillis int64, Close float64) {
	// 提前插针时间
	EarlyExecutionTimer := currentTimeMillis
	if EarlyExecutionTimer < p.LastTradeBeginTimestamp {
		//logger.Info("插针时间比当前订单的时间还要早", EarlyExecutionTimer-p.LastTradeBeginTimestamp)
		//return
	}
	EarlyExecutionTimerEnd := currentTimeMillis + (60-p.config.AssetsInsertFrame.EarlyExecutionTimer)*1000
	if EarlyExecutionTimerEnd < p.LastTradeBeginTimestamp {
		EarlyExecutionTimerEnd = p.LastTradeBeginTimestamp
		logger.Info("EarlyExecutionTimerEnd--", EarlyExecutionTimerEnd)
	}
	RecoverTs := EarlyExecutionTimerEnd
	RecoverTimer := currentTimeMillis + (60-p.config.AssetsInsertFrame.RecoverTimer)*1000
	// 到达提前执行插针时间 ~ 区间开始时间内往 最低价格*(1-资产波动幅度)位置按设置频率次数下挫
	cClose := p.HClose
	//diffClose := p.diffClose
	//logger.Info("diffClose-:", fmt.Sprintf("%v", diffClose))
	logger.Info("p.config.AssetsInsertFrame.EarlyExecutionFrequency", p.config.AssetsInsertFrame.EarlyExecutionFrequency)
	redis.Client.ZAdd("Queue:interferenceNeedle:list", goRedis.Z{
		Score:  float64(EarlyExecutionTimer),
		Member: fmt.Sprintf("Queue:interferenceNeedle:%s", p.config.AssetsCode),
	})
	for i := 0; i < p.config.AssetsInsertFrame.EarlyExecutionFrequency; i++ {
		Ts := EarlyExecutionTimer + (int64(i))*1000
		// 划分粒度，每隔1秒执行一次，一秒内300毫秒一次
		for l := 0; l < 2; l++ {
			Ts = Ts + (int64(l)+1)*300
			// 低价位平台盈利
			if strings.EqualFold("winDown", p.WinType) {
				cClose = Sub(cClose, p.config.VolatilityAmplitude, RandFloat(0, p.diffClose))
			}
			// 高价位平台盈利
			if strings.EqualFold("winUp", p.WinType) {
				cClose = Add(cClose, p.config.VolatilityAmplitude, RandFloat(0, p.diffClose))
			}
			cClose, _ = helpers.FormatFloat(cClose, p.config.PointNumber)
			logger.Info("cClose-:", time.Unix(0, Ts*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999"), "---------", cClose)
			data := redisData(p.config.AssetsCode, config.KlineSource, Ts, cClose)
			if len(data) > 0 {
				_, err := redis.Client.Lpush(fmt.Sprintf("Queue:interferenceNeedle:%s", p.config.AssetsCode), data)
				if err != nil {
					panic(err)
				}
			}
		}
	}
	// 资产在最低价格~最低价格*(1-资产震荡幅度）之间按设置震荡频率随机震荡，直到区间结束时间
	cClose = Close
	timeNow := EarlyExecutionTimer + int64(p.config.AssetsInsertFrame.EarlyExecutionFrequency)*int64(1000)
	i := 0
	var SystemTs int64
	for SystemTs = timeNow; SystemTs <= EarlyExecutionTimerEnd; SystemTs = SystemTs + 1000 {
		i++
		for l := 0; l < 2; l++ {
			SystemTs = SystemTs + (int64(l)+1)*300
			// 求余，每隔xx秒震荡一次，随机的是震荡方向
			ok := decimal.NewFromInt(int64(i)).Mod(decimal.NewFromInt(p.config.VolatilityFrequency)).IsZero()
			if ok {
				if strings.EqualFold("winDown", p.WinType) && cClose > Close { // 低价位平台盈利
					cClose = Sub(Close, p.config.VolatilityAmplitude, RandFloat(0, p.diffClose))
				} else if strings.EqualFold("winUp", p.WinType) && cClose < Close { // 高价位平台盈利
					cClose = Add(Close, p.config.VolatilityAmplitude, RandFloat(0, p.diffClose))
				} else {
					cClose = Rand(cClose, p.config.VolatilityAmplitude, RandFloat(0, p.diffClose))
					//if rand.Float64() < 0.8 { // 80%的概率出现随机震荡
					//}
					logger.Info(cClose)
					logger.Info(cClose)
				}
			}
			cClose, _ = helpers.FormatFloat(cClose, p.config.PointNumber)
			// 获取随机数 close 值
			logger.Info("cClose-:", time.Unix(0, SystemTs*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999"), "---------", cClose)
			data := redisData(p.config.AssetsCode, config.KlineSource, SystemTs, cClose)
			if len(data) > 0 {
				_, err := redis.Client.Lpush(fmt.Sprintf("Queue:interferenceNeedle:%s", p.config.AssetsCode), data)
				if err != nil {
					panic(err)
				}
			}

		}
	}
	// 追加结束信号，后面好根据结束信号来生成恢复数据插针逻辑
	logger.Info("cClose-:", time.Unix(0, SystemTs*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999"), "---------", cClose)
	data := redisData(p.config.AssetsCode, "End", SystemTs, 0)
	if len(data) > 0 {
		_, err := redis.Client.Lpush(fmt.Sprintf("Queue:interferenceNeedle:%s", p.config.AssetsCode), data)
		if err != nil {
			panic(err)
		}
	}
	logger.Info("currentTimeMillis========================", time.Unix(0, currentTimeMillis*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999"), "========================")
	logger.Info("EarlyExecutionTimer========================", time.Unix(0, EarlyExecutionTimer*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999"), "========================")
	logger.Info("SystemTs========================", time.Unix(0, SystemTs*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999"), "========================")

	// 插针明细表
	p.Peoples = helpers.RemoveRepeatedElement(p.Peoples)
	var NeedleDetails models.NeedleDetails
	NeedleDetails.AssetsCode = p.config.AssetsCode
	NeedleDetails.StartTimer = EarlyExecutionTimer
	NeedleDetails.StartDate = time.Unix(0, EarlyExecutionTimer*int64(time.Millisecond))
	NeedleDetails.EndTimer = EarlyExecutionTimerEnd
	NeedleDetails.EndDate = time.Unix(0, EarlyExecutionTimerEnd*int64(time.Millisecond))
	NeedleDetails.OrderNum = p.upNum + p.downNum
	NeedleDetails.PeopleNum = len(p.Peoples)
	//
	NeedleDetails.MarketQuoteOdds = 0
	NeedleDetails.MarketQuoteWinAmount = 0
	NeedleDetails.NeedleAfterOdds = 0
	NeedleDetails.NeedleAfterWinAmount = 0
	//
	NeedleDetails.CreateDate = time.Now()
	NeedleDetails.UpdateDate = time.Now()
	err := mysql.DB.Model(models.NeedleDetails{}).Create(&NeedleDetails).Error
	if err != nil {
		logger.Error(err)
	}
	err = mysql.DB.Model(models.TradingOrder{}).Where("trade_order_id in ?", p.OrderIds).Updates(map[string]interface{}{
		"interference_id": NeedleDetails.ID,
	}).Error
	if err != nil {
		logger.Error(err)
	}
	// 频次
	logger.Info(time.Now().UnixMilli(), EarlyExecutionTimer, time.Now().UnixMilli()-EarlyExecutionTimer, RecoverTimer, RecoverTs, EarlyExecutionTimerEnd)
}

func getNum(MonitoringNumber int, num float64) float64 {
	defer func() {
		if r := recover(); r != nil {
			logger.Info(r)
		}
	}()
	arr := strings.Split(fmt.Sprintf("%v", num), ".")
	xs := arr[1]
	res := strings.Trim(xs, "0")
	if len(res) < MonitoringNumber {
		logger.Error(errors.New("监控位数错误"))
		return 0
	}
	ok := res[len(res)-MonitoringNumber:]
	pf := math.Pow(10, float64(len(arr[1])))
	return helpers.StrToFloat64(ok) / pf
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
