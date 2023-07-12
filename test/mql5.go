package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"goapi/pkg/helpers"
	"io"
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
	listener, err := net.Listen("tcp", ":7104")
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}

	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatal("Error listener.Close: ", err)
		}
	}(listener)

	fmt.Println("Server started. Listening on port 7104...")

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
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("Error conn.Close: ", err)
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
			fmt.Println("Error reading data:", err)
			break
		}

		go processData(data)
	}
}

func processData(data []byte) {
	var mql5 Mql5
	err := json.Unmarshal(data, &mql5)
	if err != nil {
		fmt.Println("Error unmarshalling data:", err)
		return
	}
	if mql5.Name == "USDJPY" {
		fmt.Println(helpers.Decimal((mql5.Bid+mql5.Ask)/2, 5), mql5)
	}
}
