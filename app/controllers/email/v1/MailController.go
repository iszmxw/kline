package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/app/models"
	"goapi/app/response"
	"goapi/pkg/echo"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	mail_notice "goapi/pkg/notice/mail"
	"goapi/pkg/output"
	"strings"
	"time"
)

// k线图服务

type MailController struct {
	BaseController
}

/**接口编号:【U01】
 * .发送邮件【POST】
 */
// MailController
// 历史行情 mysql 查询

func (h *MailController) SendEmail(c *gin.Context) {
	//beginDate := time.Now().Format("2006-01-02 15:04:05")
	symbol := c.Query("symbol")
	period := c.Query("period")
	beginDate := c.Query("beginDate")
	endDate := c.Query("endDate")
	if len(beginDate) <= 0 || helpers.StringToInt(beginDate) <= 0 {
		echo.Error(c, "FAILD", "：beginDate 起始交易日期必填")
		return
	}
	if helpers.StringToInt64(beginDate) > time.Now().UnixMilli() {
		logger.Error(errors.New("越界查询"))
		echo.Error(c, "FAILD", "：beginDate error 越界查询")
		return
	}
	if len(endDate) <= 0 || helpers.StringToInt(endDate) <= 0 {
		echo.Error(c, "FAILD", "：endDate 截止交易日期必填")
		return
	}
	tips := fmt.Sprintf(
		"当前系统时间为【%v】 查询范围从【%v】，到【%v】",
		time.Now().Format("2006-01-02 15:04:05"),
		time.Unix(0, helpers.StringToInt64(beginDate)*int64(time.Millisecond)).Format("2006-01-02 15:04:05"),
		time.Unix(0, helpers.StringToInt64(endDate)*int64(time.Millisecond)).Format("2006-01-02 15:04:05"),
	)
	logger.Info(tips)
	if len(symbol) <= 0 || len(period) <= 0 {
		symbol = "BTC-USD"
		period = "1min"
	}
	klineSource := response.KlineSource(symbol)
	// 没有读取到 klineSource
	if len(klineSource) == 0 {
		logger.Error(errors.New("没有读取到 klineSource"))
		output.Rjson(c, nil, "没有读取到 klineSource", 0)
		return
	}
	var data []response.KlineLog
	ch := fmt.Sprintf("market.%s.index.%s", strings.ToUpper(symbol), period)
	DB := mysql.DB.Debug()
	// 一次取出历史记录最新三百条
	KlineDB := DB.Model(models.KlineLog{}).
		Where("ch", ch).
		Where("ts >= ?", beginDate).
		Where("ts <= ?", endDate).
		Order("ts desc")
	// 查询来源
	KlineDB = KlineDB.Where("kline_source", klineSource)
	var Count int64 // 查询总条数
	KlineDB.Count(&Count)
	if Count > 3000 {
		Count = 3000
	}
	err := KlineDB.Limit(int(Count)).Find(&data).Error
	if err != nil {
		logger.Error(err)
		echo.Error(c, "FAILD", err.Error())
		return
	}
	// 侦测到数据小于 150 条的时候，发送邮件预警
	if len(data) <= 150 {
		mail_notice.Notice("msgDataAbnormal", map[string]interface{}{
			"c":    c,
			"tips": tips,
		})
	}
	// 监控K线数据时间段的数据波动，实时调整赔率
	echo.Success(c, gin.H{
		"HistKlineLogs": data,
		"beginDate":     time.Unix(helpers.StringToInt64(beginDate)/1000, 0).Format("2006-01-02 15:04:05"),
		"endDate":       time.Unix(helpers.StringToInt64(endDate)/1000, 0).Format("2006-01-02 15:04:05"),
		"klineSource":   klineSource,
		"totals":        Count,
		"contractCode":  symbol,
	}, "")
}
