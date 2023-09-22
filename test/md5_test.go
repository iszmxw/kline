package test

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func Test_md5(t *testing.T) {
	body := "{\"test\":\"test\"}"
	apiKey := "K-xxxxxxx" // Your merchant key
	sign := Md5UtilsHash(body, apiKey)
	fmt.Println(sign)

}
func Md5UtilsHash(bodyJson, apiKey string) string {
	data := []byte(bodyJson + apiKey)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) // Convert []byte to decimal
	return md5str1
}
