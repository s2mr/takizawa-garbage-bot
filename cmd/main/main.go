package main

import (
	"log"

	"github.com/shimokp/takizawa-garbage-bot/manager"
)

func main() {
	//bot, err := linebot.New(secret.CHANNEL_SECRET, secret.CHANNEL_ACCESS_TOKEN)
	//
	//if err != nil {
	//	log.Println(err)
	//	panic(err)
	//}
	//
	//if _, err := bot.PushMessage(secret.USER_ID, linebot.NewTextMessage("hello"), linebot.NewTextMessage("hello")).Do(); err != nil {
	//	log.Println(err)
	//	panic(err)
	//}

	log.Println(manager.GetInstance())

	log.Println("SUCCEEDED")
}
