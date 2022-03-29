package esdk

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"goapi/pkg/logger"
	"reflect"
)

func IndexSearch(where *elastic.TermQuery) {
	// 执行ES请求需要提供一个上下文对象
	//// 创建term查询条件，用于精确查询
	//where := elastic.NewTermQuery("Author", "tizi")
	searchResult, err := Client().Search().
		Index("blogs"). // 设置索引名
		Query(where). // 设置查询条件
		Sort("Created", true). // 设置排序字段，根据Created字段升序排序，第二个参数false表示逆序
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
		var b1 Article
		// 通过Each方法，将es结果的json结构转换成struct对象
		for _, item := range searchResult.Each(reflect.TypeOf(b1)) {
			// 转换成Article对象
			if t, ok := item.(Article); ok {
				fmt.Println(t.Title)
			}
		}
	}
}
