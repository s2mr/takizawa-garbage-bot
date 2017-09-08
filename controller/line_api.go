package controller

import (
	"log"
	"net/http"
	"strings"

	"bytes"

	"os"

	"github.com/gin-gonic/gin"
)

func CallbackHandler(c *gin.Context) {

	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(c.Request.Body)
	defer c.Request.Body.Close()
	text := bufbody.String()
	log.Println(text)

	body := strings.NewReader(`{"text":` + text + `}`)

	url := os.Getenv("SlackToMe")
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
