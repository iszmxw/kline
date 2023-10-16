package activity

import (
	"fmt"
	"goapi/pkg/logger"
	"goapi/service/binary-monitoring/pkg"
	"os/exec"
	"strings"
	"time"
)

func Serve(ip string) []string {
	checkServe := make(map[string][]string)
	// 香港测试服务器
	checkServe["47.243.109.20"] = []string{
		"smt-web.jar",
		"smt-client.jar",
		"smt-paid.jar",
		"email.jar",
		"paid.jar",
		"iqoption-web.jar",
		"iqoption-agent.jar",
		"iqoption-client.jar",
		"xxl-job-admin.jar",
		"xxl-job-executor.jar",
	}
	checkServe["149.129.239.244"] = []string{
		"client.jar",
		"file.jar",
	}
	checkServe["147.139.191.241"] = []string{
		"client.jar",
	}
	checkServe["108.136.140.75"] = []string{
		"client.jar",
	}
	checkServe["108.136.242.5"] = []string{
		"web.jar",
		"agent.jar",
		"xxl-job-admin.jar",
		"xxl-job-executor.jar",
		"paid.jar",
		"callback.jar",
		"timer.jar",
	}
	checkServe["147.139.174.207"] = []string{
		"client.jar",
		"web.jar",
		"agent.jar",
		"xxl-job-admin.jar",
		"xxl-job-executor.jar",
		"paid.jar",
		"callback.jar",
		"timer.jar",
	}
	checkServe["147.139.191.215"] = []string{
		"web.jar",
		"web-second.jar",
		"agent.jar",
		"xxl-job-admin.jar",
		"xxl-job-executor.jar",
		"paid.jar",
		"email.jar",
		"file.jar",
		"callback.jar",
		"timer.jar",
	}
	checkServe["38.55.214.117"] = []string{
		"binary-go-client",
	}
	return checkServe[ip]
}

func Start(ip string) {
	for true {
		arr := Serve(ip)
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
