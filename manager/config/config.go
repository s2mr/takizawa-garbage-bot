package config

import "os"

type configManager struct {
	TGB_CHANNEL_SECRET       string
	TGB_CHANNEL_ACCESS_TOKEN string
	TGB_USER_ID              string
}

var sharedInstance *configManager = newConfigManager()

func newConfigManager() *configManager {
	cs := os.Getenv("TGB_CHANNEL_SECRET")
	cat := os.Getenv("TGB_CHANNEL_ACCESS_TOKEN")
	ui := os.Getenv("TGB_USER_ID")

	if cs == "" || cat == "" || ui == "" {
		panic("[FATAL]TGB_CHANNEL_SECRET: " + cs + " TGB_CHANNEL_ACCESS_TOKEN: " + cat + " TGB_USER_ID: " + ui)
	}

	return &configManager{cs, cat, ui}
}

func GetInstance() *configManager {
	return sharedInstance
}
