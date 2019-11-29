package router

import (
	"fmt"
	"strings"
	"talk-with-kiritan/config"
	"talk-with-kiritan/controller"
	"talk-with-kiritan/preprocessing"
	"time"

	"github.com/gin-gonic/gin"
)

// InitServer サーバの初期化
func InitServerRouter(config config.Server, loadedFiles []string, mctrl *controller.MainController) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	ctrl := mctrl.GetServerController()
	ctrl.LoadedFiles = loadedFiles
	parsedFileNames, err := parseAudioFileNames(config.AudioFileExtension, loadedFiles)
	if err != nil {
		panic(err)
	}
	ctrl.ParsedFileNames = parsedFileNames

	accounts := gin.Accounts{}
	for _, user := range config.Users {
		accounts[user.User] = user.Pass
	}

	auth := r.Group("/", gin.BasicAuth(accounts))

	root := auth.Group("/")
	{
		root.GET("/recognition", ctrl.GetRecognition)
		root.POST("/postVoiceText", ctrl.PostVoiceText)
	}

	go clock(ctrl, config.Idle)

	return r
}

func clock(ctrl *controller.ServerController, idle int) {
	for {
		ctrl.Timer.Lock.Lock()
		ctrl.Timer.AllowSend = true
		ctrl.Timer.Lock.Unlock()

		time.Sleep(time.Second * time.Duration(idle))
	}
}

func parseAudioFileNames(fileExtension string, fileNames []string) (map[string][]string, error) {
	parsedFileNames := map[string][]string{}
	for _, fileName := range fileNames {
		trimmedFileName := strings.TrimRight(fileName, fileExtension)

		keys, err := preprocessing.GetKeys(trimmedFileName)
		if err != nil {
			return nil, err
		}

		for _, key := range keys {
			parsedFileNames[key] = append(parsedFileNames[key], fileName)
		}
	}

	fmt.Printf("%d keys", len(parsedFileNames))

	return parsedFileNames, nil
}
