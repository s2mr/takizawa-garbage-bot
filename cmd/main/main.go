package main

import (
	"log"

	"time"

	"github.com/shimokp/takizawa-garbage-bot/manager"
	"github.com/shimokp/takizawa-garbage-bot/model"
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

	log.Println(manager.GetGarbageName(time.Now(), model.A))

	log.Println("SUCCEEDED")
}
