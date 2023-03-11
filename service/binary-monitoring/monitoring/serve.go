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

type Params struct {
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
	LClose   float64
	winLoss  float64
	WinType  string
	Peoples  []string
	OrderIds []string
	config   response.TradingAssetsMonitoring
}

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
			new(Params).Calculate(currentTimeMillis, item, tradingOrder)
		}
	}
}

// 初始化参数

func (p *Params) Init(config response.TradingAssetsMonitoring) bool {
	// 买涨的人数
	p.upNum = 0
	// 买涨的金额
	p.upAmounts = 0.00
	// 买跌的人数
	p.downNum = 0
	// 买跌的金额
	p.downAmounts = 0.00
	// HClose 最大 close
	p.HClose = 0.00
	// LClose 最小 close
	p.LClose = 0.00
	p.winLoss = 0.00
	p.WinType = ""
	p.Peoples = make([]string, 0)
	p.OrderIds = make([]string, 0)
	// 插针配置
	p.config = config
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
	} else {
		for _, order := range orders {
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
			// 最大close
			if order.BeginAssetsPrice > p.HClose {
				p.HClose = order.BeginAssetsPrice
			}
			// 最大close
			if order.BeginAssetsPrice < p.LClose {
				p.LClose = order.BeginAssetsPrice
			}
		}
	}
	// 检测是否需要进行干预，以及干预的方向
	ok, Close := p.finalResult()
	if !ok {
		return
	}
	// 检测是否安全期内
	ok, err := redis.Client.SelectDbSetNx(15, fmt.Sprintf("Monitoring:%s", orders[0].AssetsCode), "suo", p.config.IntervalTimer)
	if err != nil {
		return
	}
	if !ok {
		return
	}
	err = mysql.DB.Model(models.TradingOrder{}).Where("trade_order_id in ?", p.OrderIds).Updates(map[string]interface{}{
		"is_interference": 1,
	}).Error
	if err != nil {
		logger.Error(err)
	}
	// 插针明细表
	p.Peoples = helpers.RemoveRepeatedElement(p.Peoples)
	var NeedleDetails models.NeedleDetails
	NeedleDetails.AssetsCode = p.config.AssetsCode
	NeedleDetails.StartTimer = 0
	NeedleDetails.StartDate = time.Unix(0, NeedleDetails.StartTimer*int64(time.Millisecond))
	NeedleDetails.EndTimer = 0
	NeedleDetails.EndDate = time.Unix(0, NeedleDetails.EndTimer*int64(time.Millisecond))
	NeedleDetails.OrderNum = len(orders)
	NeedleDetails.PeopleNum = len(p.Peoples)
	//
	NeedleDetails.MarketQuoteOdds = 0
	NeedleDetails.MarketQuoteWinAmount = 0
	NeedleDetails.NeedleAfterOdds = 0
	NeedleDetails.NeedleAfterWinAmount = 0
	//
	NeedleDetails.CreateDate = time.Now()
	NeedleDetails.UpdateDate = time.Now()
	go p.CreateInterferenceNeedle(currentTimeMillis, Close)

}

// 最终结果
func (p *Params) finalResult() (bool, float64) {
	logger.Info("最大close", p.HClose)
	logger.Info("最小close", p.LClose)
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
	if p.winLoss < p.config.LossThreshold {
		return false, 0
	}

	// 同一个时间点有人买涨，有人买跌
	if p.HClose == p.LClose {
		if strings.EqualFold("winUp", p.WinType) { // 高价位平台盈利，算出一个高价位
			Close := p.HClose * decimal.NewFromFloat(1).Add(decimal.NewFromFloat(p.config.VolatilityAmplitude)).InexactFloat64()
			return true, Close
		}
		if strings.EqualFold("winDown", p.WinType) { // 低价位平台盈利，算出一个低价位
			Close := p.LClose * decimal.NewFromFloat(1).Sub(decimal.NewFromFloat(p.config.VolatilityAmplitude)).InexactFloat64()
			return true, Close
		}
	}
	if strings.EqualFold("winUp", p.WinType) { // 高价位平台盈利
		return true, p.HClose
	}
	if strings.EqualFold("winDown", p.WinType) { // 低价位平台盈利
		return true, p.LClose
	}
	return false, 0
}

// 创建干扰针

func (p *Params) CreateInterferenceNeedle(currentTimeMillis int64, Close float64) {
	// 提前插针时间
	EarlyExecutionTimer := currentTimeMillis + (60-p.config.EarlyexecutionTimer)*1000
	EarlyExecutionTimerEnd := currentTimeMillis + 60*1000
	RecoverTs := EarlyExecutionTimerEnd
	RecoverTimer := currentTimeMillis + (60-p.config.RecoverTimer)*1000
	// 到达提前执行插针时间 ~ 区间开始时间内往 最低价格*(1-资产波动幅度)位置按设置频率次数下挫
	cClose := p.HClose
	var diffClose float64 = 0
	// 低价位平台盈利
	if strings.EqualFold("winDown", p.WinType) {
		diffClose = decimal.NewFromFloat(p.HClose).Sub(decimal.NewFromFloat(Close)).InexactFloat64()
	}
	// 高价位平台盈利
	if strings.EqualFold("winUp", p.WinType) {
		diffClose = decimal.NewFromFloat(Close).Sub(decimal.NewFromFloat(p.LClose)).InexactFloat64()
	}
	for i := 0; i < p.config.EarlyexecutionFrequency; i++ {
		// 低价位平台盈利
		if strings.EqualFold("winDown", p.WinType) {
			diffClose = diffClose * decimal.NewFromFloat(1).Sub(decimal.NewFromFloat(p.config.VolatilityAmplitude)).InexactFloat64()
			cClose = cClose - diffClose
		}
		// 高价位平台盈利
		if strings.EqualFold("winUp", p.WinType) {
			diffClose = diffClose * decimal.NewFromFloat(1).Add(decimal.NewFromFloat(p.config.VolatilityAmplitude)).InexactFloat64()
			cClose = cClose + diffClose
		}
		cClose, _ = helpers.FormatFloat(cClose, p.config.PointNumber)
		data := redisData(p.config.AssetsCode, config.KlineSource, EarlyExecutionTimer+(int64(i))*1000, cClose)
		logger.Info(data)
		if len(data) > 0 {
			_, err := redis.Client.Lpush(fmt.Sprintf("Queue:interferenceNeedle:%s", p.config.AssetsCode), data)
			if err != nil {
				panic(err)
			}
		}
	}
	// 资产在最低价格~最低价格*(1-资产震荡幅度）之间按设置震荡频率随机震荡，直到区间结束时间
	cClose = Close
	timeNow := EarlyExecutionTimer + int64(p.config.EarlyexecutionFrequency)*1000
	i := 0
	for SystemTs := timeNow; SystemTs <= EarlyExecutionTimerEnd; SystemTs = SystemTs + 1000 {
		i++
		// 求余，每隔xx秒震荡一次，随机的是震荡方向
		ok := decimal.NewFromInt(int64(i)).Mod(decimal.NewFromInt(p.config.VolatilityFrequency)).IsZero()
		if ok {
			// 随机震荡
			if rand.Intn(2) == 0 {
				cClose = cClose * decimal.NewFromFloat(1).Sub(decimal.NewFromFloat(p.config.ShockAmplitude)).InexactFloat64()
			} else {
				cClose = cClose * decimal.NewFromFloat(1).Add(decimal.NewFromFloat(p.config.ShockAmplitude)).InexactFloat64()
			}
		}
		cClose, _ = helpers.FormatFloat(cClose, p.config.PointNumber)
		if cClose > p.LClose {
			cClose = p.LClose
		}
		result := cClose
		logger.Info(result)
		// 获取随机数 close 值
		data := redisData(p.config.AssetsCode, config.KlineSource, SystemTs, cClose)
		logger.Info(data)
		if len(data) > 0 {
			_, err := redis.Client.Lpush(fmt.Sprintf("Queue:interferenceNeedle:%s", p.config.AssetsCode), data)
			if err != nil {
				panic(err)
			}
		}
	}
	// 追加结束信号，后面好根据结束信号来生成恢复数据插针逻辑
	data := redisData(p.config.AssetsCode, "End", 0, 0)
	logger.Info(data)
	if len(data) > 0 {
		_, err := redis.Client.Lpush(fmt.Sprintf("Queue:interferenceNeedle:%s", p.config.AssetsCode), data)
		if err != nil {
			panic(err)
		}
	}
	redis.Client.ZAdd("Queue:interferenceNeedle:list", goRedis.Z{
		Score:  float64(EarlyExecutionTimer),
		Member: fmt.Sprintf("Queue:interferenceNeedle:%s", p.config.AssetsCode),
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
