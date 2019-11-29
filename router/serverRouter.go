package router

import (
	"talk-with-kiritan/config"
	"talk-with-kiritan/controller"
	"time"

	"github.com/gin-gonic/gin"
)

// InitServer サーバの初期化
func InitServerRouter(config config.Server, loadedFiles map[string][]string, mctrl *controller.MainController) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	ctrl := mctrl.GetServerController()
	ctrl.LoadedFiles = loadedFiles

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
