package serve

import (
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
	"goapi/service/binary-monitoring/pkg"
	"math/rand"
	"strings"
	"time"
)

// 最终结果
func (p *Service) finalResult() (bool, float64) {
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
		Select("MAX(c_close) as maxClose,MIN(c_close) as minClose").
		Find(&avgArr)
	p.maxClose = helpers.StrToFloat64(string(avgArr["maxClose"].([]uint8)))
	p.minClose = helpers.StrToFloat64(string(avgArr["minClose"].([]uint8)))
	p.diffClose = helpers.Decimal(p.maxClose-p.minClose, p.config.PointNumber)
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
	// 同一个时间点有人买涨，有人买跌
	if p.HClose == p.LClose {
		if strings.EqualFold("winUp", p.WinType) { // 高价位平台盈利，算出一个高价位
			Close := pkg.Add(p.HClose, p.config.VolatilityAmplitude*0.5, pkg.RandFloat(0, p.diffClose))
			return true, Close
		}
		if strings.EqualFold("winDown", p.WinType) { // 低价位平台盈利，算出一个低价位
			Close := pkg.Sub(p.LClose, p.config.VolatilityAmplitude*0.5, pkg.RandFloat(0, p.diffClose))
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

// 进行计算分析是否需要插针，以及按照规则生成插针数据

func (p *Service) Calculate(currentTimeMillis int64, conf response.TradingAssetsMonitoring, orders []response.TradingOrder) {
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
		p.LastTradeEndTimestamp = order.TradeEndTimestamp
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
			if p.LastTradeBeginTimestamp < order.TradeBeginTimestamp {
				p.LastTradeBeginTimestamp = order.TradeBeginTimestamp
			}
			if p.LastTradeEndTimestamp < order.TradeEndTimestamp {
				p.LastTradeEndTimestamp = order.TradeEndTimestamp
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
	go p.StartConsume(fmt.Sprintf("Queue:interferenceNeedle:%s", p.config.AssetsCode))
	go p.CreateInterferenceNeedle(currentTimeMillis, Close)

}

// 创建干扰针

func (p *Service) CreateInterferenceNeedle(currentTimeMillis int64, Close float64) {
	// 提前插针时间
	EarlyExecutionTimer := currentTimeMillis
	if EarlyExecutionTimer < p.LastTradeBeginTimestamp {
		//logger.Info("插针时间比当前订单的时间还要早", EarlyExecutionTimer-p.LastTradeBeginTimestamp)
		//return
	}
	EarlyExecutionTimerEnd := currentTimeMillis + (60-p.config.AssetsInsertFrame.EarlyExecutionTimer)*1000
	if EarlyExecutionTimerEnd < p.LastTradeEndTimestamp {
		EarlyExecutionTimerEnd = p.LastTradeEndTimestamp
		logger.Info("EarlyExecutionTimerEnd--", EarlyExecutionTimerEnd)
	}
	RecoverTs := EarlyExecutionTimerEnd
	RecoverTimer := currentTimeMillis + (60-p.config.AssetsInsertFrame.RecoverTimer)*1000
	// 到达提前执行插针时间 ~ 区间开始时间内往 最低价格*(1-资产波动幅度)位置按设置频率次数下挫
	cClose := Close
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
				cClose = pkg.Sub(cClose, p.config.VolatilityAmplitude, pkg.RandFloat(0, p.diffClose))
			}
			// 高价位平台盈利
			if strings.EqualFold("winUp", p.WinType) {
				cClose = pkg.Add(cClose, p.config.VolatilityAmplitude, pkg.RandFloat(0, p.diffClose))
			}
			cClose, _ = helpers.FormatFloat(cClose, p.config.PointNumber)
			logger.Info("cClose-:", time.Unix(0, Ts*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999"), "---------", cClose, "---", p.diffClose)
			data := redisData(p.config.AssetsCode, config.KlineSource, Ts, cClose)
			if len(data) > 0 {
				_, err := redis.Client.Lpush(fmt.Sprintf("Queue:interferenceNeedle:%s", p.config.AssetsCode), data)
				if err != nil {
					logger.Info(data)
					logger.Info(fmt.Sprintf("Queue:interferenceNeedle:%s", p.config.AssetsCode))
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
				// 计算价格波动方向
				randFloat64 := rand.Float64()
				if randFloat64 < 0.66666 { // 66.666%的概率出现随机震荡
					if randFloat64 < 0.33333 {
						cClose = pkg.Add(cClose, p.config.VolatilityAmplitude, pkg.RandFloat(0, p.diffClose))
					} else {
						cClose = pkg.Sub(cClose, p.config.VolatilityAmplitude, pkg.RandFloat(0, p.diffClose))
					}
				}
				if strings.EqualFold("winDown", p.WinType) && cClose > Close { // 低价位平台盈利
					cClose = pkg.Sub(Close, p.config.VolatilityAmplitude, pkg.RandFloat(0, p.diffClose))
				} else if strings.EqualFold("winUp", p.WinType) && cClose < Close { // 高价位平台盈利
					cClose = pkg.Add(Close, p.config.VolatilityAmplitude, pkg.RandFloat(0, p.diffClose))
				}
			}
			cClose, _ = helpers.FormatFloat(cClose, p.config.PointNumber)
			// 获取随机数 close 值
			logger.Info("cClose-:", time.Unix(0, SystemTs*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999"), "---------", cClose, "---", p.diffClose)
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
