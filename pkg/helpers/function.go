package helpers

// 全局函数助手

import (
	"github.com/google/uuid"
	"goapi/pkg/logger"
	"os/exec"
	"strings"
)

// GetUUID 获取UUID
func GetUUID() string {
	return uuid.New().String()
}

func GetUUID16() string {
	res := strings.ToUpper(uuid.New().String())
	arr := strings.Split(res, "-")
	return arr[0] + arr[1] + arr[2]
}

func GetUUID11() string {
	res := strings.ToUpper(uuid.New().String())
	arr := strings.Split(res, "-")
	str := arr[0] + arr[1]
	return string([]rune(str)[:11])
}

func IP() string {
	// 查询当前服务器ip
	res, err := exec.Command("curl", "ip.sb").Output()
	if err != nil {
		logger.Info("执行命令出错：", err.Error())
		// 未知
		return "Unknown"
	}
	ip := strings.Trim(string(res), "\n\n")
	return ip
}
