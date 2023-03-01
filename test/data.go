package main

import (
	"fmt"
	goadjust "goapi/pkg/adjust"
	"net/url"
)

func main() {
	req := make(url.Values)
	// Add client fields to request
	req.Add("s2s", "1")
	req.Add("app_token", "majbv8e0m8e8")
	req.Add("event_token", "y0ec7c")
	req.Add("gps_adid", "482bd535-7fcd-49ea-9678-9b9d15e81737")
	req.Add("callback_params", "{\"userId\":\"7FF235ACCD4\"}")
	//req.Add("idfa", "D2CADB5F-410F-4963-AC0C-2A78534BDF1E")
	u := "https://s2s.adjust.com/event?" + req.Encode()
	send, err := goadjust.Get(u, "2dc8c77d0503c9aab03cdeb6ebc5849b")
	if err != nil {
		return
	}
	fmt.Println(send)
}
