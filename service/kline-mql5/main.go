package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/pprof"
	KlineCommon "goapi/service/kline-common"
	kconfig "goapi/service/kline-mql5/mql5/config"
	"io"
	"net"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init(kconfig.ServiceLog)
	// 初始化 SQL
	bootstrap.SetupDB()
}

type Mql5 struct {
	Name       string      `json:"name"`
	Bid        float64     `json:"bid"`
	Ask        float64     `json:"ask"`
	Flags      int         `json:"flags"`
	Last       interface{} `json:"last"`
	Time       string      `json:"time"`
	TimeMsc    int64       `json:"timeMsc"`
	Volume     interface{} `json:"volume"`
	VolumeReal interface{} `json:"volumeReal"`
}

func main() {
	// pprof 分析
	pprof.Debug(":6071")
	bootstrap.SetupRedis(15)
	defer bootstrap.RedisClose()
	// 创建TCP监听
	listener, err := net.Listen("tcp", ":7104")
	if err != nil {
		logger.Error(err)
		return
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			logger.Error(err)
			return
		}
	}(listener)

	fmt.Println("Server started. Listening on port 7104...")

	// 处理接收到的连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error(err)
			return
		}
		// 在新的goroutine中处理连接
		go handleConnection(conn)
	}
}

// 处理Socket连接
func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			logger.Error(err)
			return
		}
	}(conn)
	reader := bufio.NewReader(conn)
	for {
		data, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed by client")
				break
			}
			logger.Error(err)
			break
		}

		go SocketMql5(data)
	}
}

func SocketMql5(data []byte) {
	var mql5 Mql5
	err := json.Unmarshal(data, &mql5)
	if err != nil {
		fmt.Println("Error unmarshalling data:", err)
		return
	}
	CodeData := KlineCommon.GetData(kconfig.KlineSource)
	for _, obj := range CodeData {
		if obj.Symbol == mql5.Name {
			logger.Info("缓存 " + obj.Symbol)
			tm := time.Unix(mql5.TimeMsc/1000, 0)
			logger.Info("时间 " + tm.Format("2006-01-02 15:04:05.999"))
			fmt.Println(helpers.Decimal((mql5.Bid+mql5.Ask)/2, 5), mql5)
		}
	}
}
