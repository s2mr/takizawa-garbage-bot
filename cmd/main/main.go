package main

import (
	"log"

	"github.com/shimokp/takizawa-garbage-bot/manager"
	"github.com/shimokp/takizawa-garbage-bot/model"
)

func main() {
	//bot, err := linebot.New(secret.CHANNEL_SECRET, secret.CH ®NNEL_ACCESS_TOKEN)
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

	log.Println(manager.GetMessage(model.Tomorrow, model.B))
	log.Println("SUCCEEDED")

	//log.Println(manager.GetGarbageName(time.Date(2018, 2, 7, 0, 0, 0, 0, &time.Location{}), model.B))
	//log.Println("SUCCEEDED")
}
