package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"goapi/app/models"
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/esdk"
	"goapi/pkg/logger"
	"reflect"
	"time"
)

type Articles struct {
	_id     string    // 文章标题
	Title   string    // 文章标题
	Content string    // 文章内容
	Author  string    // 作者
	Created time.Time // 发布时间
}

func init() {
	//var cstZone = time.FixedZone("CST", 8*3600) // 东八
	//time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("test")
	// 初始化 SQL
	bootstrap.SetupDB()
}

func main() {
	// 执行ES请求需要提供一个上下文对象
	var Count = 3000
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Filter(
		elastic.NewRangeQuery("ts").Gte("1648540103000").Lte("1648541003000"),
		elastic.NewQueryStringQuery(fmt.Sprintf("*%s*", "market.ETH-USD.index.1min")))
	source, err2 := boolQuery.Source()
	logger.Info(err2)
	logger.Info(source)
	esClient, err := esdk.Client()
	if err != nil {
		logger.Error(err)
		return
	}
	searchResult, err := esClient.Search().
		Index("kline_log"). // 设置索引名
		//Query(elastic.NewTermQuery("ch", "market.ETH-USD.index.1min")). // 设置查询条件
		//Type("doc"). // 设置查询条件
		Query(boolQuery). // 设置查询条件
		Sort("ts", true). // 设置排序字段，根据 ts 字段升序排序，第二个参数false表示逆序
		From(0). // 设置分页参数 - 起始偏移量，从第0行记录开始
		Size(Count). // 设置分页参数 - 每页大小
		Pretty(true). // 查询结果返回可读性较好的JSON格式
		Do(context.Background()) // 执行请求
	if err != nil {
		logger.Error(err)
	}
	fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits())
	if searchResult.TotalHits() > 0 {
		// 查询结果不为空，则遍历结果
		var resp models.KlineLog
		var respArr []models.KlineLog
		// 通过Each方法，将es结果的json结构转换成struct对象
		for _, item := range searchResult.Each(reflect.TypeOf(resp)) {
			// 转换成Article对象
			if t, ok := item.(models.KlineLog); ok {
				respArr = append(respArr, t)
			}
		}
		//logger.Info(respArr)
	}
}
