package main

import (
	"encoding/json"
	"fmt"
	"goapi/pkg/logger"
	"goapi/service/binary-monitoring/pkg"
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
	check := `{
    "47.242.235.4": [
        "client.jar",
        "web.jar",
        "agent.jar",
        "xxl-job-admin.jar",
        "xxl-job-executor.jar",
        "paid.jar",
        "email.jar"
    ],
    "149.129.239.244": [
        "client.jar",
        "file.jar"
    ],
    "147.139.191.241": [
        "client.jar"
    ],
    "108.136.140.75": [
        "client.jar"
    ],
    "108.136.242.5": [
        "web.jar",
        "agent.jar",
        "xxl-job-admin.jar",
        "xxl-job-executor.jar",
        "paid.jar",
        "callback.jar",
        "timer.jar"
    ],
    "147.139.174.207": [
        "client.jar",
        "web.jar",
        "agent.jar",
        "xxl-job-admin.jar",
        "xxl-job-executor.jar",
        "paid.jar",
        "callback.jar",
        "timer.jar"
    ],
    "147.139.191.215": [
        "web.jar",
        "web-second.jar",
        "agent.jar",
        "xxl-job-admin.jar",
        "xxl-job-executor.jar",
        "paid.jar",
        "email.jar",
        "file.jar",
        "callback.jar",
        "timer.jar"
    ],
    "38.55.214.117": [
        "binary-go-client"
    ]
}`
	for true {
		var checkData map[string][]string
		err = json.Unmarshal([]byte(check), &checkData)
		if err != nil {
			return
		}
		arr := checkData[ip]
		if len(arr) == 0 {
			return
		}
		// 定义要检测的程序名称
		for _, programName := range arr {
			// 检测输出中是否包含指定程序名称
			if isProcessRunning(programName) {
				logger.Info(fmt.Sprintf("程序 %s 正在运行。\n", programName))
			} else {
				logger.Info(ip, "=============", programName)
				pkg.NoticeServeKill(fmt.Sprintf("程序 %s 未启动。请在服务器【%s】上检查\n", programName, ip))
				time.Sleep(time.Second * 2)
			}
		}
		time.Sleep(time.Second * 30)
	}
}
func isProcessRunning(programName string) bool {
	cmd := exec.Command("ps", "aux")
	output, err := cmd.Output()
	if err != nil {
		logger.Error(err)
		return false
	}
	logger.Info(string(output))
	// 检查输出中是否包含程序名
	return strings.Contains(string(output), programName)
}
