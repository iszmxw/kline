package ping

import (
	"fmt"
	"github.com/go-ping/ping"
	"goapi/pkg/address"
	"goapi/pkg/logger"
	"goapi/service/binary-monitoring/pkg"
	"os"
	"os/signal"
	"time"
)

func Check(addr string) {
	for true {
		pingServe(addr)
		fmt.Println("休息一会继续")
		time.Sleep(time.Minute * 1)
	}
}

func pingServe(addr string) {
	pinger, err := ping.NewPinger(addr)
	if err != nil {
		panic(err)
	}
	pinger.Timeout = time.Second * 10
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			pinger.Stop()
		}
	}()
	pinger.OnRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
	}
	pinger.OnDuplicateRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v (DUP!)\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.Ttl)
	}
	pinger.OnFinish = func(stats *ping.Statistics) {
		msg := ""
		msg = msg + fmt.Sprintf("\n--- %s ping statistics ---\n", stats.Addr)
		msg = msg + fmt.Sprintf("%d 传输的数据包, %d 收到的数据包, %v%% 丢包\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
		msg = msg + fmt.Sprintf("往返 min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
		fmt.Println(msg)
		if 100 == stats.PacketLoss {
			res := address.GetMessageByIp("")
			pkg.Notice(fmt.Sprintf("---服务器【%s】\n\n---访问不通【%s】\n\n具体情况如下：\n\n%s\n\n", res.Query, addr, msg))
		}
	}
	fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	err = pinger.Run()
	if err != nil {
		logger.Error(err)
	}
}
