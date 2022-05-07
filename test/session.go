package main

import (
	"goapi/config"
	"goapi/pkg/logger"
	"io"
	"net/http"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("session")
}

func main() {

	url := "https://olymptrade.com/platform"
	var c = make(chan interface{})
	go func() {
		Client := &http.Client{}
		get, err := Client.Get(url)
		if err != nil {
			return
		}
		c <- get.Header.Values("Set-Cookie")
	}()
	for {
		select {
		case res, _ := <-c:
			getSession(res)
		}
	}

}
// 6Lc5ACgTAAAAABL3s8j9VkkUn4Engv4QtdCLd9qI
// 6Lc5ACgTAAAAABL3s8j9VkkUn4Engv4QtdCLd9qI

func getSession(Cookies interface{}) {
	url := "https://olymptrade.com/platform"
	//logger.Info(Cookies)
	req, _ := http.NewRequest("GET", url, nil)
	//for _, item := range Cookies.([]string) {
	//	logger.Info(item)
	//}
	//logger.Info(strings.Join(Cookies.([]string), " "))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
	req.Header.Set("Cookie", "checked=1; guest_id=1000006125155658318958528064978641650421679517729511400270316226; enterdate=2022-04-20+05%3A27%3A59; lang=zh_CN; _ga=GA1.2.b4675346426104.1650421680385; ga_block_once=yes; hubspotutk=eda8a4de2ccf8404b2a19e255fda6368; __hssrc=1; _fbp=fb.1.1650421684654.709923444; ecp=hi_AXZujaPKIo1wlwthz1Ab3B2Oxa3CEtp8; ece=hi_AXZujaPKIo1wlwthz1Ab3B2Oxa3CEtp8; ecc=hi_AXZujaPKIo1wlwthz1Ab3B2Oxa3CEtp8; uhdwidv2=hi_AXZujaPKIo1wlwthz1Ab3B2Oxa3CEtp8; session=1629167146143495028583624161042912173895462127666685937664207996; __hstc=95761603.eda8a4de2ccf8404b2a19e255fda6368.1650421682745.1650421682745.1650424793854.2; __hssc=95761603.1.1650424793854; jwt_auth=eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpYXQiOjE2NTA0MjQ3OTYsIm5iZiI6MTY1MDQyNDc5NiwiZXhwIjoxNjUwNTExMTk2LCJ1c2VyX2lkIjo4MzYxNDEzMCwiZW1haWwiOiI0MzUzMTE5MzlAcXEuY29tIiwiMmZhdSI6MH0.L2aPODEz5t6YwpOxqeFslzKcvgaOad8yprVGfv5yuWJcbOtuJsJvWL_yszUpgtRrZWQkASYK8KwIGdH9tOJ9izt5qkqz9D7T4cDwCZMK5CSpd-Ys4oeKIpgrxLyJ78hoiR1VcvFzYX1M1ju02h1TOv2b9HTU-uCNluTcpgKs8RTk7BYL2-ZNvI6MFwuz2pLgk08D1IkWQNxL4gnvLNCLIsUzg_qJjCvG13Mt3fgP_3Frfg-I0Wtsot6E80gzbCQQVICcM3S5TYXy_g6NA6svWtpCinO8iJqC8SYtvOt4PkKNJsx1xGP5nLj6fKFaEB6YNFtpSLy0S6SCJBHj2TORcJQAf51-6hXqTF4BTkEQG86T5QGi51fjl0qJJU19BnzydstWfkhrfuj4jLAMXRunv4rXyBicnGCv8igsUi2ERoTQl6o7BNEDZGo8bclJ4v-uPe60OLBaj2CAFZLohqDEPAnJClG4GZQnPiTCA_9WeCP1xav05kPgdgn-yTzVFnDHSR3Qu4d1wAsRmqfdgy6fNzhxCQ4ChGltCZaZrwxOQMmxZjPmnbzV-PG0QqICpu0TVkna_jsf2_98OK_PataRmEyEelds_9efWA5MEn0tqAtudA8T8JFntSqS7x1hVQfDyA12511jE6o24aZmbemU4aZSQvoHINUwBcMPZa4Cayc")

	resp, err := (&http.Client{}).Do(req)
	//resp, err := http.Get(serviceUrl + "/topic/query/false/lsj")
	if err != nil {
		logger.Error(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Error(err)
		}
	}(resp.Body)
	logger.Info(resp.Header)
}
