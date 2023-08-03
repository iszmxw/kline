package utils

import (
	"goapi/pkg/address"
	"goapi/pkg/redis"
	"net"
	"os"
	"strings"
)

// 获取服务器的Uuid

func ServeUuid() string {
	addresses, _ := getServerMACAddresses()
	return Md5(addresses)
}

func ServeIp() string {
	key := "ServeIp:" + ServeUuid()
	info, _ := redis.Client.SelectDbGet(0, key)
	if len(info) == 0 {
		res := address.GetMessageByIp("")
		if len(res.Query) > 0 {
			_, _ = redis.Client.SelectDbAdd(0, key, res.Query, 0)
			info = res.Query
		}
	}
	return info
}

func getServerMACAddresses() (string, error) {
	var macAddresses string
	macAddresses, _ = os.Hostname()
	interfaces, err := net.Interfaces()
	if err != nil {
		return macAddresses, err
	}
	for _, iface := range interfaces {
		// 过滤掉回环接口和虚拟接口
		if strings.HasPrefix(iface.Name, "lo") || strings.HasPrefix(iface.Name, "vir") {
			continue
		}
		mac := iface.HardwareAddr.String()
		if mac != "" {
			macAddresses = macAddresses + mac
		}
	}
	return macAddresses, nil
}
