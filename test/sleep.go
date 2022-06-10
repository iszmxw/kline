package main

import (
	"encoding/hex"
	"fmt"
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/logger"
	kline_common "goapi/service/kline-common"
	"goapi/service/kline-common/heartbeat"
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
	// 初始化 SQL
	bootstrap.SetupDB()
}

func main() {
	bootstrap.SetupRedis(0)
	defer bootstrap.RedisClose()
	CodeData := kline_common.GetData("qxbroker")
	// 启动监控，监控数据源活跃状态，设置参数每个30秒检测一次
	heartbeat.Start(30, "qxbroker", CodeData)
	time.Sleep(time.Minute * 10)
	// {
	//    "signal_server": "gwt5.mql5.com:443",
	//    "ping": 1,
	//    "key": "eeb4988a5bcbc12438bbf5042281edf8427730ebd9cff70751e56764176ff25a",
	//    "token": "s2CJRjBXzyudyR7hgLhFCWZdJkaxAUTSnJT5szHNTyikxSRWCXBN3edpajSD3dfx",
	//    "version": 5,
	//    "enabled": true
	//}
	// 登录字符串 hex
	logByte1, _ := hex.DecodeString("480000000200000000008782f84164007972417a785a4d6169465058655259586653387772624d5963546571324b4557514473506e444461333466486877346750465a354d6d43677875774e45673246")
	logByte2, _ := hex.DecodeString("48000000020000000000a8b3534264007a4a4e6a67516b5879384c6e594a59387376434b776e44664b5a53356a415053706454747245557575667233654a4d4b6d466862557845376472483952704d52")
	logByte3, _ := hex.DecodeString("480000000200000000002adf59426400373450504e575643376d734b387667684d70545263425457544b447a4e324357467746524a4c723251327945654e444878386d3866744643694b505773764864")
	logByte4, _ := hex.DecodeString("48000000020000000000bba55b4264007332434a526a42587a79756479523768674c684643575a644a6b6178415554536e4a5435737a484e5479696b785352574358424e33656470616a534433646678")
	logByte5, _ := hex.DecodeString("4800000002000000000029bd614264006b695868775a557a465752716842797171526471417973785679706d7741663370666d527a556d444543457a434b4546384279664d7737444a7a755276465658")
	fmt.Println(logByte1)
	fmt.Println(logByte2)
	fmt.Println(logByte3)
	fmt.Println(logByte4)
	fmt.Println(logByte5)
	// H.........].kBd.4hdERp3MccwgLgQw7Be2fQbgTrpXeD3AaqQPD4LNBdqwCQKUcd9nuSs7g7Dz9XFq

}

func ja(a, b string) {
	//let c = new Uint8Array(8 + (b && b.byteLength || 0));
	//b && c.set(new Uint8Array(b), 8);
	//b = new DataView(c.buffer);
	//b.setUint16(0, va, !0);
	//b.setUint32(2, (new Date).getTime(), !0);
	//b.setUint16(6, a, !0);
	//return c.buffer
}

// token=== ZbxztS2w5ELhyrycBTLBRHsZXv8MgYQjEKBeVicNWScTufdbTFtREpNcRFvRDnpe
//func token(token string) {
//let b = new ArrayBuffer(64)
//w.setCharString(b, a, 0)
//return b
//}
