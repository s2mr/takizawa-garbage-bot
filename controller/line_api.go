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
	"github.com/shimokp/takizawa-garbage-bot/constant"
	"github.com/shimokp/takizawa-garbage-bot/manager/config"
	"github.com/shimokp/takizawa-garbage-bot/manager/database"
	"github.com/shimokp/takizawa-garbage-bot/manager/garbage"
	"github.com/shimokp/takizawa-garbage-bot/model"
)

func RootHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func PostHandler(c *gin.Context) {
	nickname, _ := c.GetPostForm("nickname")
	body, _ := c.GetPostForm("body")

	err := sendToSlackForRequest(nickname, body)
	if err != nil {
		log.Println(err)
	}

	c.HTML(http.StatusOK, "sent.html", gin.H{})
}

func CheckHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "check.html", gin.H{})
}

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

		err = sendToSlackForEvent(event)
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

func MultiHandler(c *gin.Context) {
	var resp = ""
	regionStr, _ := c.GetPostForm("region")
	dateTypeStr, _ := c.GetPostForm("dateType")

	log.Println(regionStr, dateTypeStr)

	err := sendMessage(model.ConvertStringToRegion(regionStr), model.ConvertStringToDateType(dateTypeStr))
	if err != nil {
		resp = err.Error()
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

func switchMessage(event model.Event) string {
	if event.Type == "follow" {
		return constant.MESSAGE_FIRST_RESPONSE
	}

	text := event.Message.Text

	isExists, _ := model.IsUserExists(database.GetInstance().DB, event.Source.UserID)
	if !isExists {
		/*ユーザが存在しない場合*/
		if text == "A" || text == "B" {
			return registerUser(event.Source.UserID, model.ConvertStringToRegion(text))
		} else {
			return constant.MESSAGE_COMMAND_NOTFOUND
		}
	}

	user, err := model.GetUserByUserId(database.GetInstance().DB, event.Source.UserID)
	if err != nil {
		log.Println(err)
		return "ユーザ取得時にエラーが発生しました。"
	}

	switch text {
	case "今日":
		return garbage.GetMessage(model.Today, user.Region)
	case "明日":
		return garbage.GetMessage(model.Tomorrow, user.Region)
	}

	if text == "A" || text == "B" {
		return updateUser(user.UserID, model.ConvertStringToRegion(text))
	}

	return constant.MESSAGE_COMMAND_NOTFOUND
}

func registerUser(userId string, region model.Region) string {
	err := model.InsertUser(database.GetInstance().DB, userId, region)
	if err != nil {
		log.Println(err)
		return "地区設定時にエラーが発生しました。"
	}
	return "地区を" + model.ConvertRegionToString(region) + "に設定しました。\n\n" + constant.MESSAGE_COMMAND_INSTRUCTION
}

func updateUser(userId string, region model.Region) string {
	err := model.UpdateUser(database.GetInstance().DB, userId, region)
	if err != nil {
		log.Println(err)
		return "更新時にエラーが発生しました。"
	}

	return "地区を" + model.ConvertRegionToString(region) + "に変更しました。"
}

func sendMessage(region model.Region, dateType model.DateType) error {
	users, err := model.GetUsersByRegion(database.GetInstance().DB, region)
	if err != nil {
		log.Println(err)
		return err
	}
	ids := model.GetUserIdsFromUsers(users)
	if len(ids) == 0 {
		return errors.New("user to send is empty")
	}
	bot, err := linebot.New(config.GetInstance().TGB_CHANNEL_SECRET, config.GetInstance().TGB_CHANNEL_ACCESS_TOKEN)
	if err != nil {
		log.Println(err)
		return err
	}
	str := garbage.GetMessage(dateType, region)
	if _, err := bot.Multicast(ids, linebot.NewTextMessage(str)).Do(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func getProfile(userId string) (model.Profile, error) {

	var profile model.Profile

	bot, err := linebot.New(config.GetInstance().TGB_CHANNEL_SECRET, config.GetInstance().TGB_CHANNEL_ACCESS_TOKEN)
	if err != nil {
		log.Println(err)
		return profile, err
	}
	res, err := bot.GetProfile(userId).Do()
	if err != nil {
		log.Println(err)
		return profile, err
	}

	profile.DisplayName = res.DisplayName
	profile.PictureURL = res.PictureURL
	profile.StatusMessage = res.StatusMessage

	return profile, nil
}

func sendToSlackForRequest(nickName string, body string) error {
	s := fmt.Sprintf(`
	{ "text" : "ご要望を受信しました！
ニックネーム: %s
本文: %s"}`,
		nickName,
		body)

	return sendToSlack(s)
}

func sendToSlackForEvent(event model.Event) error {
	prof, err := getProfile(event.Source.UserID)
	if err != nil {
		return err
	}

	s := fmt.Sprintf(`
	{ 	"text" : "name: %s
画像のURL: %s
ステータス: %s
本文: %s
 "}`,
		prof.DisplayName,
		prof.PictureURL,
		prof.StatusMessage,
		event.Message.Text)

	return sendToSlack(s)
}

func sendToSlack(s string) error {
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
