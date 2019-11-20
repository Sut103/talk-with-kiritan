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

var (
	fileNames map[string]string //トリミングしたファイル名と元のファイル名
)

func init() {
	extension := ".wav"
	ignoreSymbols := []string{"。", "、", ",", ".", "・", "_", "＿", "!", "！", "?", "？", " ", "　", "…"}
	fileNames = map[string]string{}

	fmt.Println("Loading sound file ...")
	files, err := ioutil.ReadDir("sounds")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fileName := file.Name()
		trimmedFileName := strings.TrimRight(fileName, extension)
		if trimmedFileName+extension == fileName { // 拡張子のバリデーション
			for _, ignoreSymbol := range ignoreSymbols {
				trimmedFileName = strings.ReplaceAll(trimmedFileName, ignoreSymbol, "")
			}
			fileNames[trimmedFileName] = fileName

		}
	}
	fmt.Println("Sound file was Loaded!")
}

func InitMainRouter(config config.Config) (*discordgo.Session, *gin.Engine, error) {
	ctrl := controller.GetMainController()

	dg, err := InitDiscordRouter(config.Discord, ctrl)
	if err != nil {
		return nil, nil, err
	}

	g := InitServerRouter(fileNames, ctrl)

	return dg, g, err
}
