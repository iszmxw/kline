package test

import (
	"encoding/json"
	"fmt"
	"goapi/app/response"
	"testing"
)

func TestSource(t *testing.T) {
	source := response.AllowKlineSource()
	marshal, err := json.Marshal(source)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
	fmt.Println()
}
