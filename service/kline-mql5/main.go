package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"goapi/bootstrap"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/logger"
	"goapi/pkg/pprof"
	"goapi/service/kline-mql5/mql5"
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
	Client := new(mql5.WebSocketOp)
	Client.Init()
	//go Client.Read()
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
		handleConnection(rw, Client)
	}
}

// 处理Socket连接
func handleConnection(conn net.Conn, ws *mql5.WebSocketOp) {
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
		ws.SystemTs = time.Now().UnixMilli()
		ws.SocketMql5(bytes.Trim(message, "\x00"))
		//ws.ReadChan <- bytes.Trim(message, "\x00")
	}
}
