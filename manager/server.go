package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/shimokp/takizawa-garbage-bot/controller"
	"github.com/shimokp/takizawa-garbage-bot/manager/config"
)

func Init() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", controller.RootHandler)
	r.POST("/callback", controller.CallbackHandler)
	r.POST("/multi", controller.MultiHandler)
	r.Run(":" + config.GetInstance().PORT) // listen and serve on 0.0.0.0:8080
}
