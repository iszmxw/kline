package main

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

func main() {
	res := strings.ToUpper(uuid.New().String())
	arr := strings.Split(res, "-")
	str := arr[0] + arr[1] + arr[2]
	fmt.Println(str)
}
