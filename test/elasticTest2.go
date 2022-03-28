package main

import (
	"fmt"
	"goapi/pkg/elastic"
	"time"
)

type Article struct {
	_id     string    // 文章标题
	Title   string    // 文章标题
	Content string    // 文章内容
	Author  string    // 作者
	Created time.Time // 发布时间
}

func main() {
	// 定义一篇博客
	data := Article{_id: "666", Title: "golang es教程", Content: "go如何操作ES", Author: "tizi", Created: time.Now()}
	fmt.Println(elastic.IndexPut(data._id, data))
}
