package main

import (
	"encoding/json"
	"fmt"
	"goapi/config"
	"goapi/pkg/logger"
	"os"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("test")
}

type S struct {
	ChangeType    int     `json:"changeType"`
	ChangeBalance float64 `json:"changeBalance"`
	SourceIndex   int64   `json:"sourceIndex"`
	TimerBalance  float64 `json:"timerBalance"`
}

func main() {
	filePtr, err := os.Open("./test/xw.json")
	if err != nil {
		fmt.Printf("文件打开失败 [Err:%s]\n", err.Error())
		return
	}
	defer func(filePtr *os.File) {
		err := filePtr.Close()
		if err != nil {
			logger.Error(err)
		}
	}(filePtr)
	var info []S
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&info)
	if err != nil {
		fmt.Println("解码失败", err.Error())
	} else {
		for index, item := range info {
			fmt.Println(index, item)
		}
	}

}
