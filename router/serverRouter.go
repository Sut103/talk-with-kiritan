package router

import (
	"talk-with-kiritan/controller"

	"github.com/gin-gonic/gin"
)

// InitServer サーバの初期化
func InitServer() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/recognition", controller.GetRecognition)
	r.POST("/postVoiceText", controller.PostVoiceText)

	return r
}
