package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"talk-with-kiritan/controller"

	"github.com/gin-gonic/gin"
)

var (
	fileNames map[string]string //トリミングしたファイル名と元のファイル名
)

func init() {
	extension := ".wav"
	ignoreSymbols := []string{"。", "、", ",", ".", "・", "_", "＿", "!", "！", "?", "？", " ", "　", "…"}

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

			fileNames := map[string]string{}
			fileNames[trimmedFileName] = fileName
		}
	}

	fmt.Println("Sound file was Loaded!")
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/recognition", controller.GetRecognition)
	r.POST("/postVoiceText", controller.PostVoiceText)

	if err := r.Run(); err != nil {
		panic(err)
	}
}
