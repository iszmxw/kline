package main

import (
	"encoding/json"
	"fmt"
	"goapi/pkg/helpers"
	"log"
	"net"
)

/**

{
    "name": "EURUSD",
    "bid": 1.09648,
    "ask": 1.09648,
    "flags": 6,
    "last": 0,
    "time": "2023.07.10 12:52",
    "timeMsc": 1688993565655,
    "volume": 0,
    "volumeReal": 0
}
*/

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
	// 创建TCP监听
	listener, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}

	defer listener.Close()

	fmt.Println("Server started. Listening on port 7777...")

	// 处理接收到的连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting connection: ", err)
		}

		// 在新的goroutine中处理连接
		go handleConnection(conn)
	}
}

// 处理Socket连接
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// 处理接收到的数据
	buffer := make([]byte, 3060)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}

		// 处理接收到的数据
		data := buffer[:n]
		if len(data) > 0 {
			var mql5 Mql5
			err = json.Unmarshal(data, &mql5)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Println(helpers.Decimal((mql5.Bid+mql5.Ask)/2, 5), mql5)
		}
		//fmt.Println("Received data:", string(data))
	}
}
