package main

import "C"
import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"goapi/bootstrap"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/logger"
	"goapi/pkg/pprof"
	"goapi/pkg/redis"
	KlineCommon "goapi/service/kline-common"
	"goapi/service/kline-mql5/mql5"
	kconfig "goapi/service/kline-mql5/mql5/config"
	"goapi/service/kline-mql5/mql5/wsdata"
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

func main() {
	AppPort := flag.Int64("APP_PORT", conf.GetInt64("app.port"), "服务端口")
	flag.Parse()
	// pprof 分析
	pprof.Debug(":6071")
	bootstrap.SetupRedis(15)
	defer bootstrap.RedisClose()
	// 创建TCP监听
	logger.Info(*AppPort)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", *AppPort))
	if err != nil {
		logger.Error(err)
		return
	}
	defer func(listener net.Listener) {
		deferErr := listener.Close()
		if deferErr != nil {
			logger.Error(deferErr)
			return
		}
	}(listener)
	logger.Info("服务已启动，等待接收数据")
	CodeData := KlineCommon.GetData(kconfig.KlineSource)
	for _, obj := range CodeData {
		go SubPriceKLine(obj)
	}
	// 处理接收到的连接
	var tempDelay time.Duration
	for {
		rw, errs := listener.Accept()
		if errs != nil {
			if ne, ok := errs.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				logger.Error(errs)
				logger.Info(tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			logger.Error(errs)
			panic(errs)
		}
		tempDelay = 0
		// 在新的goroutine中处理连接
		handleConnection(rw)
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
		message, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				logger.Error(errors.New("客户端链接关闭"))
				break
			}
			logger.Error(err)
			break
		}
		var mql5data wsdata.Mql5
		err = json.Unmarshal(message, &mql5data)
		if err != nil {
			logger.Error(err)
			time.Sleep(time.Millisecond * 100)
			continue
		}
		info := checkExists(mql5data.Name)
		if len(info.ContractCode) <= 0 {
			continue
		}
		pushKey := fmt.Sprintf("%s:%s", "list-mql5", info.ContractCode)
		bytesData := bytes.Trim(message, "\x00")
		_, _ = redis.Client.Lpush(pushKey, string(bytesData))
	}
}

func checkExists(code string) KlineCommon.Info {
	CodeData := KlineCommon.GetData(kconfig.KlineSource)
	var info KlineCommon.Info
	for _, item := range CodeData {
		if item.Symbol == code {
			info = item
			return info
		}
	}
	return info
}

func SubPriceKLine(info KlineCommon.Info) {
	pushKey := fmt.Sprintf("%s:%s", "list-mql5", info.ContractCode)
	for {
		data := redis.Client.RPop(pushKey)
		if len(data) <= 0 {
			time.Sleep(time.Millisecond * 300)
		}
		var mql5data wsdata.Mql5
		err := json.Unmarshal([]byte(data), &mql5data)
		if err != nil {
			time.Sleep(time.Millisecond * 300)
			continue
		}
		mql5.SocketMql5(info, mql5data)
	}
}
