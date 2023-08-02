package main

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
	C1 := Init()
	go SubPriceKLine(C1)
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
		handleConnection(rw, C1)
	}
}

// 处理Socket连接
func handleConnection(conn net.Conn, Client *Client) {
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
		if len(info.Symbol) <= 0 {
			continue
		}
		list := &wsdata.LinkedList{}
		// 入队数据
		bytesData := bytes.Trim(message, "\x00")
		list.Enqueue(bytesData)
		Client.SubTopicMap[info.Symbol] = &info
		Client.LinkedList[info.Symbol] = list
	}
}

type Client struct {
	SubTopicMap map[string]*KlineCommon.Info  // 订阅历史
	IsSub       map[string]bool               // 是否已经订阅了
	LinkedList  map[string]*wsdata.LinkedList // 订阅历史
}

func Init() *Client {
	C1 := new(Client)
	C1.IsSub = make(map[string]bool)
	C1.SubTopicMap = make(map[string]*KlineCommon.Info)
	C1.LinkedList = make(map[string]*wsdata.LinkedList)
	return C1
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

func SubPriceKLine(C *Client) {
	for {
		for key, val := range C.SubTopicMap {
			if C.IsSub[key] {
				time.Sleep(time.Millisecond * 500)
				continue
			}
			C.IsSub[key] = true
			logger.Info(key, val)
			go func() {
				for {
					for s, u := range C.LinkedList {
						if u == nil || u.Length == 0 {
							continue
						}
						bytesData := u.Dequeue().([]byte)
						logger.Info(s, string(bytesData))
						info := checkExists(s)
						var mql5data wsdata.Mql5
						err := json.Unmarshal(bytesData, &mql5data)
						if err != nil {
							logger.Error(err)
							continue
						}
						go mql5.SocketMql5(info, mql5data)
					}
					time.Sleep(time.Second)
				}
			}()
		}
	}
}
