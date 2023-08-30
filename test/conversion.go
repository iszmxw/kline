package main

import (
	"fmt"
	"math/rand"
	"strings"
)

/*
*
1. `getRandomEncryption()`：生成一个随机字符映射字符串。该方法将一个字符列表分成两半，然后随机选择一个字符对，将它们拼接起来，重复这个过程直到所有字符都被使用。返回生成的字符映射字符串。

2. `getRandomEncryptionMap(String enc)`：根据给定的字符映射字符串生成一个字符映射的Map。将字符映射字符串分成两个字符一组，然后将每组的第一个字符作为key，第二个字符作为value，将它们放入Map中。返回生成的字符映射Map。

3. `getEncryptionString(String string)`：根据之前生成的字符映射Map，将给定的字符串进行加密。遍历字符串中的每个字符，如果是小写字母或大写字母，则通过字符映射Map获取对应的加密字符，否则保持原字符不变。返回加密后的字符串。

4. `main(String[] args)`：测试用的主方法。在这个方法中，先定义了一个字符串`str`，然后调用`getRandomEncryption()`生成一个随机字符映射字符串并赋值给`enc`，接着调用`getEncryptionString()`将字符串"nGCuSf"进行加密并赋值给`res`，最后打印输出`res`和`enc`。

这段代码的作用是将给定的字符串进行加密，加密规则是将字符映射表中的字符替换成对应的加密字符。
*/
var (
	encoding = "aIbkczdZefgrhijJlEmDnCoLpWqTstuvwAxMyPBGFHKQNVOSRYUX"
	charList = []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o",
		"p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D",
		"E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S",
		"T", "U", "V", "W", "X", "Y", "Z",
	}
	charMap = map[byte]byte{}
)

func getRandomEncryption() string {
	enc := strings.Builder{}
	length := len(charList) / 2
	for i := 0; i < length; i++ {
		first := charList[0]
		random := rand.Intn(len(charList)-1) + 1
		nextChar := charList[random]
		charList = append(charList[:0], charList[1:]...)
		charList = append(charList[:random-1], charList[random:]...)
		enc.WriteString(first + nextChar)
	}
	return enc.String()
}

func getRandomEncryptionMap(enc string) map[byte]byte {
	str := enc
	charMap := make(map[byte]byte)
	for i := 0; i < len(enc)/2; i++ {
		first := str[:2]
		str = str[2:]
		charMap[first[0]] = first[1]
		charMap[first[1]] = first[0]
	}
	return charMap
}

func getEncryptionString(str string) string {
	if len(charMap) == 0 {
		charMap = getRandomEncryptionMap(encoding)
	}
	enc := strings.Builder{}
	for _, c := range str {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			enc.WriteByte(charMap[byte(c)])
		} else {
			enc.WriteRune(c)
		}
	}
	return enc.String()
}

func main() {
	//str := "iszmxw"
	enc := getRandomEncryption()
	res := getEncryptionString("itemCode")
	hsfDnLZf := getEncryptionString("hsfDnLZf")
	iszmxw := getEncryptionString("iszmxw")
	htcDMA := getEncryptionString("htcDMA")
	fmt.Println(res, iszmxw)
	fmt.Println(hsfDnLZf, htcDMA)
	fmt.Println(enc)
}
