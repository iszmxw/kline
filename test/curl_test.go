package test

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func Test_Curl(t *testing.T) {
	// 从文件中读取内容
	data, err := ioutil.ReadFile("curl.txt")
	if err != nil {
		fmt.Println("无法读取文件:", err)
		return
	}
	// 将读取的内容存储在变量中
	str := string(data)
	str = strings.ReplaceAll(str, "-->", "")
	str = strings.Replace(str, "POST", "", 1)
	arr := strings.Split(str, "\n")
	res := ""
	for i, s := range arr {
		s = strings.Trim(s, " ")
		res += "\n"
		if i == 0 {
			res = res + fmt.Sprintf(`curl '%v' \`, s)
		} else if i == len(arr)-1 {
			res = res + fmt.Sprintf(`--data-raw '%v' \`, s)
		} else {
			res = res + fmt.Sprintf(`-H '%v' \`, s)
		}
	}
	fmt.Println(res + "\n--compressed")
	fmt.Println()
	fmt.Println()
	fmt.Println()

}

//
//curl 'https://web2.iqoptionproo.com/web/v1/loginout/login.json' \
//-H 'Accept: */*' \
//-H 'Accept-Language: en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7' \
//-H 'Cache-Control: no-cache' \
//-H 'Connection: keep-alive' \
//-H 'Content-Type: application/json;charset=UTF-8' \
//-H 'Origin: https://smo.ssmartoption.com' \
//-H 'Pragma: no-cache' \
//-H 'Referer: https://smo.ssmartoption.com/' \
//-H 'Sec-Fetch-Dest: empty' \
//-H 'Sec-Fetch-Mode: cors' \
//-H 'Sec-Fetch-Site: cross-site' \
//-H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36' \
//-H 'sec-ch-ua: "Chromium";v="116", "Not)A;Brand";v="24", "Google Chrome";v="116"' \
//-H 'sec-ch-ua-mobile: ?0' \
//-H 'sec-ch-ua-platform: "macOS"' \
//-H 'token: null' \
//--data-raw '{"languageCode":"CH","itemCode":"SMO","account":"zhiming","password":"a8d5e6bc6a4a4b81ca4224f10bb946ae","nation":"Indonesia"}' \
//--compressed
