package main

import (
	"fmt"
	goadjust "goapi/pkg/adjust"
	"log"
	"time"
)

func main() {
	client := goadjust.New("tiwyto5q38jk", "")
	resp, err := client.TrackEvent("event_token", string(goadjust.IDFA), "6ce9c0a255c6aa171a715c6be0dd8bf9", time.Now())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.TrackerName)
}
