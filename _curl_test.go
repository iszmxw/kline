package main

import (
	"encoding/json"
	"fmt"
	"goapi/pkg/utils"
	"io/ioutil"
	"strings"
	"testing"
)

func Test_Curl(t *testing.T) {
	// 从文件中读取内容
	data, err := ioutil.ReadFile("./test/curl.txt")
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
			// 是否解密 url 参数
			s = deSrc(s)
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

func deSrc(s string) string {
	message := []byte(s)
	tempMap := make(map[string]interface{})
	err := json.Unmarshal(message, &tempMap)
	if err != nil {
		panic(err)
	}
	result := conversionData(tempMap)
	message, _ = json.Marshal(result)
	return string(message)
}
func conversionData(data map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for key, val := range data {
		vType := fmt.Sprintf("%T", val)
		if vType == "map[string]interface {}" {
			tempRes := conversionData(val.(map[string]interface{}))
			result[utils.GetEncryptionString(key)] = tempRes
		} else {
			result[utils.GetEncryptionString(key)] = val
		}
	}
	return result
}
