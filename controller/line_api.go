package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/pkg/errors"
	"github.com/shimokp/takizawa-garbage-bot/constant"
	"github.com/shimokp/takizawa-garbage-bot/manager/config"
	"github.com/shimokp/takizawa-garbage-bot/manager/database"
	"github.com/shimokp/takizawa-garbage-bot/manager/garbage"
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
		log.Println("UnmarshalError::", err)
		addString(&resp, err.Error())
	}

	//TODO: 一個でもエラーが出たら連鎖していろんなifに引っかかりそう
	if len(message.Events) > 0 {
		event := message.Events[0]

		err = sendToSlack(event)
		if err != nil {
			log.Println("SendToSlackError::", err)
			addString(&resp, err.Error())
		}

		err = replyMessage(event)
		if err != nil {
			log.Println("ReplyMessageError::", err)
			addString(&resp, err.Error())
		} else {
			log.Println("Reply Send!")
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"message": resp,
	})
}

func addString(base *string, text string) {
	*base = *base + "\n" + text
}

func replyMessage(event model.Event) error {
	sendMessage := switchMessage(event)

	if sendMessage == "" {
		return errors.New("sendMessage is empty")
	}

	bot, err := linebot.New(config.GetInstance().TGB_CHANNEL_SECRET, config.GetInstance().TGB_CHANNEL_ACCESS_TOKEN)
	if err != nil {
		return err
	}
	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(sendMessage)).Do(); err != nil {
		return err
	}

	return nil
}

var users []model.User

func switchMessage(event model.Event) string {

	if event.Type == "follow" {
		return constant.MESSAGE_FIRST_RESPONSE
	}

	user, err := model.GetUserByUserId(database.GetInstance().DB, event.Source.UserID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			//FIXME: use const
			return "地区が登録されていません"
		default:
			log.Println(err)
			return "エラーが発生しました"
		}
	}

	switch event.Message.Text {
	case "今日":
		return garbage.GetMessage(model.Today, user.Region)
	case "明日":
		return garbage.GetMessage(model.Tomorrow, user.Region)
	case "A":
		return registerUser(event.Source.UserID, model.A)
	case "B":
		return registerUser(event.Source.UserID, model.B)
	}

	return ""
}

func registerUser(userID string, region model.Region) string {
	//TODO:もし既に登録されている場合は更新する

	err := model.InsertUser(database.GetInstance().DB, userID, region)
	if err != nil {
		log.Println(err)
		return "エラーが発生しました"
	}
	return "登録しました"
}

func sendToSlack(event model.Event) error {
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
		" ``` "+` "}`, event.ReplyToken,
		event.Type,
		event.Timestamp,
		event.Source.Type,
		event.Source.UserID,
		event.Message.ID,
		event.Message.Type,
		event.Message.Text)

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
