package config

import "goapi/pkg/config"

// email 参数
func init() {
	config.Add("mongo", config.StrMap{
		"host":     config.Env("MONGODB_HOST", "127.0.0.1"), // 启用或禁用嗅探器（默认启用）。
		"port":     config.Env("MONGODB_PORT", 27017),       // elasticsearch 服务地址，多个服务地址使用逗号分隔
		"database": config.Env("MONGODB_DATABASE", ""),      // 启用gzip压缩
		"username": config.Env("MONGODB_USERNAME", ""),      // 启用gzip压缩
		"password": config.Env("MONGODB_PASSWORD", ""),      // 启用gzip压缩
	})
}
