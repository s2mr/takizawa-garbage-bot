package main

import (
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/shimokp/takizawa-garbage-bot/secret"
)

func main() {
	bot, err := linebot.New(secret.CHANNEL_SECRET, secret.CHANNEL_ACCESS_TOKEN)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	//var messages []linebot.Message
	//
	//message := linebot.NewTextMessage("hello")
	//messages = append(messages, message)

	// log.Printf("%#v(%v)", text, reflect.TypeOf(text))

	if _, err := bot.PushMessage(secret.USER_ID, linebot.NewTextMessage("hello"), linebot.NewTextMessage("hello")).Do(); err != nil {
		log.Println(err)
		panic(err)
	}

	log.Println("SUCCEEDED")
}
