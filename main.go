package main

import (
	"talk-with-kiritan/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/recognition", controller.GetRecognition)
	r.POST("/postVoiceText", controller.PostVoiceText)

	r.Run()
}
