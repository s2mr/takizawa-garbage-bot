package manager

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shimokp/takizawa-garbage-bot/controller"
)

func Init() {
	r := gin.Default()
	r.POST("/callback", controller.CallbackHandler)
	port := os.Getenv("PORT")
	r.Run(":" + port) // listen and serve on 0.0.0.0:8080
}
