package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var data interface{}
	err := json.Unmarshal([]byte(`[{"d":[{"p":"_BRN","t":1655200595.681,"q":121.282}],"e":1}]`), &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if fmt.Sprintf("%T", data) == "[]interface {}" {
		dataMap := data.([]interface{})[0]
		if fmt.Sprintf("%T", dataMap) == "map[string]interface {}" {
			if fmt.Sprintf("%v", dataMap.(map[string]interface{})["e"]) == "1" {
				array := dataMap.(map[string]interface{})["d"]
				arrayMap := array.([]interface{})[0].(map[string]interface{})
				fmt.Println(fmt.Sprintf("%T",arrayMap["t"]))
			}
			//if dataMap.(map[string]interface{})["d"] {
			//
			//}
		}
	}
}
