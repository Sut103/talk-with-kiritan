package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Voice struct {
	Text string `form:"text"`
}

func GetRecognition(c *gin.Context) {
	c.HTML(http.StatusOK, "recognition.html", nil)
}

func PostVoiceText(c *gin.Context) {
	voice := Voice{}

	err := c.ShouldBind(&voice)
	if err != nil {
		panic(err)
	}

	fmt.Println(voice.Text)
}
