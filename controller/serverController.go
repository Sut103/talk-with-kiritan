package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServerController struct {
	Main      *MainController
	FileNames map[string]string
}

type Voice struct {
	Text string `form:"text"`
}

func (sctrl *ServerController) GetRecognition(c *gin.Context) {
	c.HTML(http.StatusOK, "recognition.html", nil)
}

func (sctrl *ServerController) PostVoiceText(c *gin.Context) {
	voice := Voice{}

	err := c.ShouldBind(&voice)
	if err != nil {
		panic(err)
	}

	fmt.Println("Input text ---> '", voice.Text, "'")

	if fileName, ok := sctrl.FileNames[voice.Text]; ok {
		sctrl.Main.VChs.Ch <- fileName
	}

}
