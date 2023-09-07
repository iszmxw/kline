package main

import (
	"fmt"
	"goapi/pkg/utils"
)

func main() {
	str := "itemCode" //qUQwZERQ
	fmt.Println(utils.GetRandomEncryption())
	fmt.Println(utils.GetRandomEncryptionToLowerCase())
	fmt.Println(utils.GetEncryptionString("path"))
	fmt.Println(utils.GetEncryptionString("itemCode"))
	fmt.Println(utils.GetEncryptionString("qUQwZERQ"))
	fmt.Println(utils.GetEncryptionToLowerCase("itemcode"))
	fmt.Println(utils.GetEncryptionToLowerCase("edisngti"))
	fmt.Println(str)
}
