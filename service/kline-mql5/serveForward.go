package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"goapi/config"
	"goapi/pkg/logger"
	"goapi/pkg/pprof"
	KlineCommon "goapi/service/kline-common"
	kconfig "goapi/service/kline-mql5/mql5/config"
	"goapi/service/kline-mql5/mql5/wsdata"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

/**

转发服务，主要是用来转发 mql5 的数据到各个服务器，用http的方式转发

 **/

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init(kconfig.ServiceLog)
}

func main() {
	// pprof 分析
	pprof.Debug(":6072")
	// 创建TCP监听
	listener, err := net.Listen("tcp", ":5555")
	if err != nil {
		logger.Error(err)
		return
	}
	defer listener.Close()
	logger.Info("服务已启动，等待接收数据")
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
	defer conn.Close()
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
		bytesData := bytes.Trim(message, "\x00")
		// 转发服务，将数据分发到不同的服务器
		Forward(info.ContractCode, string(bytesData))
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

// 转发数据

func Forward(index, item string) {
	serves := []string{
		//"http://47.243.109.20:7104",
		"http://172.31.0.4:7104",
		"https://goclientio.smosignal.com",
		"https://huobiapi.smooption.com",
		"https://goclient.smarttradess.com",
	}
	for _, HOST := range serves {
		baseURL := HOST + "/app/v3/other/pushList.json"
		go pushToClient(baseURL, index, item)
	}
	//baseURL := "http://47.242.235.4:7104/app/v3/other/pushList.json"
	//go pushToClient(baseURL, index, item)
}

type Res struct {
	Code int `json:"code"`
	Data struct {
		Result bool `json:"result"`
	} `json:"data"`
	Msg   string `json:"msg"`
	ReqId string `json:"reqId"`
}

func pushToClient(baseURL, index, item string) {
	params := url.Values{}
	params.Set("index", index)
	params.Set("item", item)
	urls := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	response, err := http.Get(urls)
	if err != nil {
		logger.Error(err)
		return
	}
	if response.StatusCode != 200 {
		logger.Info(response.Status, "-----", urls)
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Error(err)
		return
	}
	var res Res
	err = json.Unmarshal(body, &res)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(res)
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logger.Error(err)
		}
	}(response.Body)
}
