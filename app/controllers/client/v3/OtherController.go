package v3

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/pkg/config"
	"goapi/pkg/echo"
	"goapi/pkg/logger"
	"goapi/pkg/mongodb"
	"goapi/pkg/redis"
	"time"
)

// k线图服务

type OtherController struct {
	BaseController
}

// 接口信息

func (h *OtherController) Info(c *gin.Context) {
	data := cmap.New().Items()
	data["name"] = "SmartOption3.0"
	data["url"] = "http://192.168.192.95:7104/doc/swagger.json"
	data["location"] = "http://192.168.192.95:7104/doc/swagger.json"
	data["swaggerVersion"] = "3.0"
	var result []map[string]interface{}
	result = append(result, data)
	c.JSON(200, result)
}

// 获取 15 库的所有键值对

func (h *OtherController) PushList(c *gin.Context) {
	index := c.Query("index")
	item := c.Query("item")
	if len(item) > 0 && len(index) > 0 {
		ok, err := redis.Client.LPushDb(15, fmt.Sprintf("list-mql5:%s", index), item)
		if err != nil {
			echo.Error(c, "Failed", err.Error())
			return
		}
		echo.Success(c, gin.H{"result": ok}, "")
		return
	}
	echo.Error(c, "Failed", "")
}

type Test struct {
	ItemCode        string `bson:"itemCode" json:"itemCode"`
	Name            string `bson:"name" json:"name"`
	Value           string `bson:"value" json:"value"`
	CreateTimestamp int64  `bson:"createTimestamp" json:"createTimestamp"`
}

func (h *OtherController) Mongo(c *gin.Context) {
	name := c.Query("name")
	value := c.Query("value")
	if len(name) > 0 && len(value) > 0 {
		delayInfo := Test{
			ItemCode:        config.GetString("app.item_code"),
			Name:            name,
			Value:           value,
			CreateTimestamp: time.Now().UnixMilli(),
		}
		collection := mongodb.MgoCli.Database("smo").Collection("tb_test")
		//插入某一条数据
		inst, err := collection.InsertOne(context.TODO(), delayInfo)
		if err != nil {
			logger.Error(err)
			echo.Error(c, "Failed", err.Error())
			return
		}
		echo.Success(c, gin.H{"result": inst.InsertedID}, "")
		return
	}
	echo.Error(c, "Failed", "")
}
