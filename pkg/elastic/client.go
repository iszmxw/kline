package elastic

import (
	"github.com/olivere/elastic"
	"goapi/pkg/logger"
	"log"
	"os"
	"time"
)

var EsClient *elastic.Client

func Client() *elastic.Client {
	var err error
	// 创建client连接ES
	EsClient, err = elastic.NewClient(
		elastic.SetSniff(false),
		// elasticsearch 服务地址，多个服务地址使用逗号分隔
		elastic.SetURL("http://192.168.0.221:9200"),
		// 启用gzip压缩
		elastic.SetGzip(true),
		// 设置监控检查时间间隔
		elastic.SetHealthcheckInterval(10*time.Second),
		// 设置请求失败最大重试次数
		elastic.SetMaxRetries(5),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)
	if err != nil {
		logger.Info(err)
	}

	return EsClient
}
