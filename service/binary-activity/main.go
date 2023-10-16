package main

import (
	"fmt"
	"goapi/pkg/logger"
	"goapi/service/binary-activity/activity"
	"goapi/service/binary-activity/monitor"
	"os/exec"
	"strings"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	// 定义日志目录
	logger.Init("activity")
}

func main() {
	// 查询当前服务器ip
	res, err := exec.Command("curl", "ip.sb").Output()
	if err != nil {
		fmt.Println("执行命令出错：", err)
		return
	}
	ip := strings.Trim(string(res), "\n\n")
	logger.Info("ip：", ip)
	// 检测服务存活状态
	go activity.Start(ip)
	// 检测磁盘是否占用高
	go monitor.Start(ip)
	select {}
}
