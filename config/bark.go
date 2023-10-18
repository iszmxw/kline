package config

import "goapi/pkg/config"

// redis 参数
func init() {
	config.Add("bark", config.StrMap{
		"url":       config.Env("bark.url", "https://api.day.app/"),
		"keyMac":    config.Env("bark.key.mac", "YvemvqzPAgvc2ASqwcP2KH"),
		"keyIphone": config.Env("bark.key.iphone", "nPuFoXfYsMpVrUaqU8to99"),
		"logo":      config.Env("bark.logo", "https://avatars.githubusercontent.com/u/135967790?s=400&u=0efb3cc947e9f0c2165c11f65f374524cb48915d&v=4"),
	})
}
