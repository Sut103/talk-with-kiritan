package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/recognition", getRecognition)
	r.POST("/postVoiceText", postVoiceText)

	r.Run()
}

func getRecognition(c *gin.Context) {
	c.HTML(http.StatusOK, "recognition.html", nil)
}

type Voice struct {
	Text string `form:"text"`
}

func postVoiceText(c *gin.Context) {
	voice := Voice{}

	err := c.ShouldBind(&voice)
	if err != nil {
		panic(err)
	}

	fmt.Println(voice.Text)
}
