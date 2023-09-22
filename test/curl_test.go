package test

import (
	"fmt"
	"github.com/atotto/clipboard"
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
	result := res + "\n--compressed"
	fmt.Println(result)
	//复制内容到剪贴板
	err = clipboard.WriteAll(result)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()

}
