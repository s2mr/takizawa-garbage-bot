package tgb

import (
	"github.com/gin-gonic/gin"
	"github.com/shimokp/takizawa-garbage-bot/controller"
)

func Init() {
	r := gin.Default()
	r.GET("/callback", controller.CallbackHandler)
	r.Run("localhost:80") // listen and serve on 0.0.0.0:8080
}
