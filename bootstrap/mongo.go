package bootstrap

import (
	"goapi/pkg/mongodb"
)

// SetupMongoDB 初始化MongoDB
func SetupMongoDB() {
	mongodb.MgoCli = mongodb.GetMgoCli()
}
