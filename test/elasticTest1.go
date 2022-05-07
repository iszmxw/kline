package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
)

var client *elastic.Client

// 索引mapping定义，这里仿微博消息结构定义
const mapping = `
{
  "mappings": {
    "properties": {
      "user": {
        "type": "keyword"
      },
      "message": {
        "type": "text"
      },
      "image": {
        "type": "keyword"
      },
      "created": {
        "type": "date"
      },
      "tags": {
        "type": "keyword"
      },
      "location": {
        "type": "geo_point"
      },
      "suggest_field": {
        "type": "completion"
      }
    }
  }
}`

func main() {
	var err error
	var exists bool
	// 设置ES服务地址，支持多个地址
	urlOpt := elastic.SetURL("http://10.10.10.10:9200")
	sniffOpt := elastic.SetSniff(false)
	// 设置基于http base auth验证的账号和密码
	//elastic.SetBasicAuth("user", "secret"))
	// 创建ES client用于后续操作ES
	client, err = elastic.NewClient(urlOpt, sniffOpt)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Println("connect to es success")

	// 执行ES请求需要提供一个上下文对象
	ctx := context.Background()
	// 首先检测下weibo索引是否存在
	exists, err = client.IndexExists("weibo1").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// weibo索引不存在，则创建一个
		_, err = client.CreateIndex("weibo1").BodyString(mapping).Do(ctx)
		if err != nil {
			// Handle error
			fmt.Println(err.Error())
		}
	}
}
