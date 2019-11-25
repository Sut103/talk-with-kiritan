package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"sort"
	"strings"
	"talk-with-kiritan/preprocessing"
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
	if ctrl.Main.Clock {
		voice := Voice{}

		err := c.ShouldBind(&voice)
		if err != nil {
			panic(err)
		}

		fmt.Println("Input text ---> '", voice.Text, "'")

		vtext := strings.ReplaceAll(voice.Text, " ", "")
		keys, err := preprocessing.GetKeys(vtext)
		if err != nil {
			panic(err)
		}

		sort.Slice(keys, func(i, j int) bool { return len(keys[i]) > len(keys[j]) })

		for _, key := range keys {
			if fileNames, ok := ctrl.LoadedFiles[key]; ok {
				count := len(fileNames)
				rand.Seed(time.Now().UnixNano())

				randNum := 0
				if count != 1 {
					randNum = rand.Intn(count - 1)
				}

				ctrl.Main.VChs.Ch <- fileNames[randNum]
				ctrl.Main.Clock = false
				break
			}
		}
	}
}
