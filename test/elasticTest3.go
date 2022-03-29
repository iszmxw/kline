package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
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

func main() {
	// 执行ES请求需要提供一个上下文对象
	//// 创建term查询条件，用于精确查询
	//where := elastic.NewTermQuery("Author", "tizi")
	searchResult, err := esdk.Client().Search().
		Index("kline_log"). // 设置索引名
		Query(elastic.NewMatchAllQuery()). // 设置查询条件
		//Sort("Created", true). // 设置排序字段，根据Created字段升序排序，第二个参数false表示逆序
		From(0). // 设置分页参数 - 起始偏移量，从第0行记录开始
		Size(10). // 设置分页参数 - 每页大小
		Pretty(true). // 查询结果返回可读性较好的JSON格式
		Do(context.Background()) // 执行请求
	if err != nil {
		logger.Error(err)
	}
	fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits())

	if searchResult.TotalHits() > 0 {
		// 查询结果不为空，则遍历结果
		var b1 Articles
		// 通过Each方法，将es结果的json结构转换成struct对象
		for _, item := range searchResult.Each(reflect.TypeOf(b1)) {
			// 转换成Article对象
			if t, ok := item.(Articles); ok {
				fmt.Println(t.Title)
			}
		}
	}
}
