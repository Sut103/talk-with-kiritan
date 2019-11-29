package controller

import (
	"math/rand"
	"net/http"
	"sort"
	"strings"
	"sync"
	"talk-with-kiritan/preprocessing"
	"time"

	"github.com/gin-gonic/gin"
)

type ServerController struct {
	Main            *MainController
	LoadedFiles     []string
	ParsedFileNames map[string][]string
	Timer           Timer
}

type Timer struct {
	AllowSend bool
	Lock      sync.Mutex
}

type Voice struct {
	Text string `form:"text"`
}

type ResponseVoiceText struct {
	Input    string `json:"input"`
	Morpheme string `json:"morpheme"`
	FileName string `json:"file_name"`
}

type ResponseFileNames struct {
	Names []string `json:"names"`
}

func (ctrl *ServerController) GetRecognition(c *gin.Context) {
	c.HTML(http.StatusOK, "recognition.html", nil)
}

func (ctrl *ServerController) GetFileNames(c *gin.Context) {
	res := ResponseFileNames{ctrl.LoadedFiles}
	c.JSON(http.StatusOK, res)
}

func (ctrl *ServerController) PostVoiceText(c *gin.Context) {
	voice := Voice{}
	err := c.ShouldBind(&voice)
	if err != nil {
		panic(err)
	}

	res := ResponseVoiceText{Input: voice.Text}
	if ctrl.Timer.AllowSend && ctrl.Main.VChs.Condition {
		vtext := strings.ReplaceAll(voice.Text, " ", "")
		keys, err := preprocessing.GetKeys(vtext)
		if err != nil {
			panic(err)
		}

		sort.Slice(keys, func(i, j int) bool { return len(keys[i]) > len(keys[j]) })

		for _, key := range keys {
			if fileNames, ok := ctrl.ParsedFileNames[key]; ok {
				count := len(fileNames)
				rand.Seed(time.Now().UnixNano())

				randNum := 0
				if count != 1 {
					randNum = rand.Intn(count - 1)
				}

				sendFileName := fileNames[randNum]
				res.Morpheme = key
				res.FileName = sendFileName

				ctrl.Main.VChs.Ch <- sendFileName

				ctrl.Timer.Lock.Lock()
				ctrl.Timer.AllowSend = false
				ctrl.Timer.Lock.Unlock()

				break
			}
		}
	}
	c.JSON(http.StatusOK, res)
}
