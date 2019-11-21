package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ServerController struct {
	Main        *MainController
	LoadedFiles map[string][]string
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

	if fileNames, ok := ctrl.LoadedFiles[voice.Text]; ok {
		count := len(fileNames)
		rand.Seed(time.Now().UnixNano())

		randNum := 0
		if count != 1 {
			randNum = rand.Intn(count - 1)
		}

		ctrl.Main.VChs.Ch <- fileNames[randNum]
	}

}
