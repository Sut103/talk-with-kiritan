package router

import (
	"fmt"
	"io/ioutil"
	"strings"
	"talk-with-kiritan/config"
	"talk-with-kiritan/controller"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

func InitMainRouter(config config.Config) (*discordgo.Session, *gin.Engine, error) {
	ctrl := controller.GetMainController()

	loadedFiles, err := loadAudioFiles(config)
	if err != nil {
		return nil, nil, err
	}

	dg, err := InitDiscordRouter(config.Discord, ctrl)
	if err != nil {
		return nil, nil, err
	}

	g := InitServerRouter(loadedFiles, ctrl)

	return dg, g, err
}

func loadAudioFiles(config config.Config) (map[string][]string, error) {
	extension := ".wav"
	loadedFiles := map[string][]string{}

	fmt.Println("Loading sound file ...")
	files, err := ioutil.ReadDir("sounds")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		fileName := file.Name()
		trimmedFileName := strings.TrimRight(fileName, extension)
		if trimmedFileName+extension == fileName { // 拡張子のバリデーション
			loadedFiles[trimmedFileName] = append(loadedFiles[trimmedFileName], fileName)
		}
	}
	fmt.Println("Sound file was Loaded!")

	return loadedFiles, nil
}
