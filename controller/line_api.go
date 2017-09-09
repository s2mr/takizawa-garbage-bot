package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/pkg/errors"
	"github.com/shimokp/takizawa-garbage-bot/manager/config"
	"github.com/shimokp/takizawa-garbage-bot/model"
)

func CallbackHandler(c *gin.Context) {
	var resp = ""

	bufBody := new(bytes.Buffer)
	bufBody.ReadFrom(c.Request.Body)
	defer c.Request.Body.Close()
	log.Println("RequestBuffer::", bufBody.String())

	var message = model.MessageText{}
	err := json.Unmarshal(bufBody.Bytes(), &message)
	if err != nil {
		log.Println(err)
		addString(&resp, err.Error())
	}

	err = sendToSlack(message)
	if err != nil {
		log.Println(err)
		addString(&resp, err.Error())
	}

	err = returnMessage(message)
	if err != nil {
		log.Println(err)
		addString(&resp, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": resp,
	})
}

// TODO: ちゃんと動いてる？？
func addString(base *string, text string) {
	*base = *base + "\n" + text
}

func returnMessage(message model.MessageText) error {
	if len(message.Events) == 0 {
		return errors.New("out of range")
	}
	event := message.Events[0]

	bot, err := linebot.New(config.GetInstance().TGB_CHANNEL_SECRET, config.GetInstance().TGB_CHANNEL_ACCESS_TOKEN)
	if err != nil {
		return err
	}
	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("リプライ！")).Do(); err != nil {
		return err
	}

	return nil
}

func sendToSlack(message model.MessageText) error {

	s := fmt.Sprintf(`
	{ 	"text" : " `+
		" ``` "+
		`replyToken: %s
type: %s
timeStamp: %s
---Source---
type: %s
userId: %s
iD: %s
---Message---
type: %s
text: %s`+
		" ``` "+` "}`, message.Events[0].ReplyToken,
		message.Events[0].Type,
		message.Events[0].Timestamp,
		message.Events[0].Source.Type,
		message.Events[0].Source.UserID,
		message.Events[0].Message.ID,
		message.Events[0].Message.Type,
		message.Events[0].Message.Text)

	body := strings.NewReader(s)

	url := os.Getenv("SlackToMe")
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	log.Println("Slack Responce Status::", resp.Status)

	bufBody := new(bytes.Buffer)
	_, err = bufBody.ReadFrom(resp.Body)
	if err != nil {
		return err
	}
	log.Println("Slack Responce::", bufBody.String())
	defer resp.Body.Close()
	return nil
}
