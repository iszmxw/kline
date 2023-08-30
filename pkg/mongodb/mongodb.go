package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goapi/pkg/config"
	"goapi/pkg/logger"
)

var MgoCli *mongo.Client

func initEngine() {
	var err error
	credential := options.Credential{
		AuthSource: config.GetString("mongo.database"),
		Username:   config.GetString("mongo.username"),
		Password:   config.GetString("mongo.password"),
	}
	applyURI := fmt.Sprintf("mongodb://%s:%s", config.GetString("mongo.host"), config.GetString("mongo.port"))
	clientOptions := options.Client().
		ApplyURI(applyURI).
		SetAuth(credential)

	// 连接到MongoDB
	MgoCli, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logger.Error(err)
	}
	// 检查连接
	err = MgoCli.Ping(context.TODO(), nil)
	if err != nil {
		logger.Error(err)
	}
}

func GetMgoCli() *mongo.Client {
	if MgoCli == nil {
		initEngine()
	}
	return MgoCli
}
