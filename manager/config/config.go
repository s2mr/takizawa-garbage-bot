package config

import (
	"os"
	"reflect"
)

type configManager struct {
	PORT                     string
	TGB_CHANNEL_SECRET       string
	TGB_CHANNEL_ACCESS_TOKEN string
	TGB_USER_ID              string
	DATABASE_URL             string
}

var sharedInstance *configManager = newConfigManager()

func newConfigManager() *configManager {
	port := os.Getenv("PORT")
	cs := os.Getenv("TGB_CHANNEL_SECRET")
	cat := os.Getenv("TGB_CHANNEL_ACCESS_TOKEN")
	ui := os.Getenv("TGB_USER_ID")
	durl := os.Getenv("DATABASE_URL")

	//FIXME: sliceを使わなくてもいいようにしたい
	slice := []string{port, cs, cat, ui, durl}
	for i := 0; i < len(slice); i++ {
		if slice[i] == "" {
			panic("[FATAL]" + reflect.ValueOf(configManager{}).Type().Field(i).Name + " is not assign")
		}
	}

	return &configManager{port, cs, cat, ui, durl}
}

func GetInstance() *configManager {
	return sharedInstance
}
