package main

import (
	"encoding/json"
	"fmt"
	"goapi/app/models"
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/address"
	conf "goapi/pkg/config"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("binary-test")
}

func main() {
	// 初始化 SQL
	// 定义日志目录
	logger.Info("初始化 SQL")
	bootstrap.SetupDB()
	// 初始化 Redis
	logger.Info("初始化 Redis")
	db := conf.GetInt("redis.db")
	bootstrap.SetupRedis(db)
	defer bootstrap.RedisClose()
	// 查询ip
	var list []models.AppUser
	mysql.DB.Debug().Model(models.AppUser{}).Where("mobile IS NOT NULL").Where("item_code", conf.GetString("app.item_code")).Limit(100000).Find(&list)
	if len(list) == 0 {
		return
	}
	fmt.Println(len(list))
	for index, item := range list {
		getNation(index, item)
	}

	//data := address.GetMessageByIp("140.213.70.48")
	fmt.Println("data")
	time.Sleep(1500000 * time.Second)
}

func getNation(index int, item models.AppUser) {
	ok, _ := redis.Client.Get(fmt.Sprintf("Nation:Indonesia:%s", item.RegistIp))
	if strings.EqualFold("ok", ok) {
		//fmt.Println(index, "已缓存")
		return
	}
	var data address.Address
	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("http://ip-api.com/json/%v", item.RegistIp), nil)
	if err != nil {
		getNation(index, item)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		getNation(index, item)
		return
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			getNation(index, item)
			return
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		getNation(index, item)
		return
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		getNation(index, item)
		return
	}
	if strings.EqualFold(data.Country, "Indonesia") {
		_, _ = redis.Client.Add(fmt.Sprintf("Nation:Indonesia:%s", item.RegistIp), "ok", 0)
		fmt.Println(index, "ok", item.AppUserId)
	} else {
		fmt.Println(index, item.AppUserId, item.RegistIp, data.Country)
	}
}
