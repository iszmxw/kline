package esdk

import (
	"context"
	"goapi/pkg/logger"
	"time"
)

type Article struct {
	_id     string    // 文章标题
	Title   string    // 文章标题
	Content string    // 文章内容
	Author  string    // 作者
	Created time.Time // 发布时间
}

func IndexPut(id string, body interface{}) bool {
	var err error
	// 执行ES请求需要提供一个上下文对象
	ctx := context.Background()
	// 使用client创建一个新的文档
	_, err = Client().Index().
		Index("kline_log"). // 设置索引名称
		Type("doc").
		Id(id). // 设置文档id
		BodyJson(body). // 指定前面声明struct对象
		Refresh("true"). // 执行操作后刷新索引。
		Do(ctx) // 执行请求，需要传入一个上下文对象
	if err != nil {
		logger.Error(err)
		return false
	}
	return true
}
