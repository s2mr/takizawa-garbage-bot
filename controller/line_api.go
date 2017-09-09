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
	"github.com/shimokp/takizawa-garbage-bot/model"
)

func CallbackHandler(c *gin.Context) {
	var resp = "ok"

	bufBody := new(bytes.Buffer)
	bufBody.ReadFrom(c.Request.Body)
	defer c.Request.Body.Close()
	log.Println("RequestBuffer::", bufBody.String())

	var message = model.MessageText{}
	err := json.Unmarshal(bufBody.Bytes(), &message)
	if err != nil {
		log.Println(err)
		resp = err.Error()
	}

	err = sendToSlack(message)
	if err != nil {
		log.Println(err)
		resp = resp + err.Error()
	}

	c.JSON(http.StatusOK, gin.H{
		"message": resp,
	})
}

func sendToSlack(message model.MessageText) error {

	s := fmt.Sprintf(`
	{ 	"text" : " `+
		" ``` "+
		`replayToken: %s
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
