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

func (ctrl *ServerController) GetRecognition(c *gin.Context) {
	c.HTML(http.StatusOK, "recognition.html", nil)
}

func (ctrl *ServerController) PostVoiceText(c *gin.Context) {
	voice := Voice{}

	err := c.ShouldBind(&voice)
	if err != nil {
		panic(err)
	}

	fmt.Println("Input text ---> '", voice.Text, "'")

	if fileName, ok := ctrl.FileNames[voice.Text]; ok {
		ctrl.Main.VChs.Ch <- fileName
	}

}
