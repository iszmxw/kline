package main

import (
	"fmt"
	"goapi/service/kline-hsb/hsb/utils/zip"
)

func main() {
	resp := "48000000020000000000cb65445264007270695061334d704c637a4862516159533266684c544e6167396a5366784d76684b7379346e775438417738396e68634e5459485141677174684662414e3237"
	message, err := zip.GZipDecompress([]byte(resp))
	if err != nil {
		return
	}
	fmt.Println(string(message))
}
