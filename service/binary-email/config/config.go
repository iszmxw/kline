package config

import (
	"errors"
	"goapi/app/models"
	"goapi/app/response"
	"goapi/pkg/mysql"
)

const SystemEmailRoutingNode = "system_email_routing_node"

// 读取一条配置信息

func GetOne(nation, channel string) (error, response.SystemEmailRoutingNode) {
	var result response.SystemEmailRoutingNode
	//key := fmt.Sprintf("%v:%v:%v", SystemEmailRoutingNode, nation, channel)
	//// 1、先从缓存读取配置信息
	//strConfig, _ := redis.Client.Get(key)
	//if len(strConfig) > 0 {
	//	err := json.Unmarshal([]byte(strConfig), &result)
	//	if err != nil {
	//		return err, result
	//	}
	//	return nil, result
	//}
	// 2、缓存没有直接读取数据库
	mysql.DB.Model(models.SystemEmailRoutingNode{}).Where(map[string]interface{}{
		"valid":   1,
		"nation":  nation,
		"channel": channel,
	}).Find(&result)
	if result.Id == 0 {
		return errors.New("数据库未找到相关渠道配置"), result
	}
	//// 3、数据库读取到了在缓存一遍
	//_, _ = redis.Client.Add(key, "", 0)
	return nil, result
}

// 读取所有配置

func GetAll(name string) {

}
