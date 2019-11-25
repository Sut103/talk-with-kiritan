package router

import (
	"talk-with-kiritan/controller"
	"time"

	"github.com/gin-gonic/gin"
)

// InitServer サーバの初期化
func InitServerRouter(loadedFiles map[string][]string, mctrl *controller.MainController) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	ctrl := mctrl.GetServerController()
	ctrl.LoadedFiles = loadedFiles

	auth := r.Group("/", gin.BasicAuth(gin.Accounts{
		"tohoku": "zunko",
	}))

	root := auth.Group("/")
	{
		root.GET("/recognition", ctrl.GetRecognition)
		root.POST("/postVoiceText", ctrl.PostVoiceText)
	}

	go clock(ctrl)

	return r
}

func clock(ctrl *controller.ServerController) {
	for {
		ctrl.Timer.Lock.Lock()
		ctrl.Timer.AllowSend = true
		ctrl.Timer.Lock.Unlock()

		time.Sleep(time.Second * 45)
	}
}
