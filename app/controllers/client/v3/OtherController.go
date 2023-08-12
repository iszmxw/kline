package v3

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
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
	Name            string `bson:"name" json:"name"`
	Value           string `bson:"value" json:"value"`
	CreateTimestamp int64  `bson:"createTimestamp" json:"createTimestamp"`
}

func (h *OtherController) Mongo(c *gin.Context) {
	name := c.Query("name")
	value := c.Query("value")
	if len(value) > 0 && len(name) > 0 {
		delayInfo := Test{
			Name:            name,
			Value:           value,
			CreateTimestamp: time.Now().UnixMilli(),
		}
		collection := mongodb.MgoCli.Database("smo").Collection("tb_delay_statistics")
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
