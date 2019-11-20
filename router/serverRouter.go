package router

import (
	"talk-with-kiritan/controller"

	"github.com/gin-gonic/gin"
)

// InitServer サーバの初期化
func InitServerRouter(fileNames map[string]string, mctrl *controller.MainController) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	ctrl := mctrl.GetServerController()
	ctrl.FileNames = fileNames

	r.GET("/recognition", ctrl.GetRecognition)
	r.POST("/postVoiceText", ctrl.PostVoiceText)

	return r
}
