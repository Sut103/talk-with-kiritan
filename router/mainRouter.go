package router

import (
	"fmt"
	"io/ioutil"
	"strings"
	"talk-with-kiritan/config"
	"talk-with-kiritan/controller"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	mecab "github.com/shogo82148/go-mecab"
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
			keys, err := getKeys(trimmedFileName)
			if err != nil {
				return nil, err
			}

			for _, key := range keys {
				loadedFiles[key] = append(loadedFiles[key], fileName)
			}
		}
	}
	fmt.Println("Sound file was Loaded!")

	return loadedFiles, nil
}

func getKeys(fileName string) ([]string, error) {
	tagger, err := mecab.New(map[string]string{})
	if err != nil {
		return nil, err
	}
	defer tagger.Destroy()

	result, err := tagger.ParseToNode(fileName)
	if err != nil {
		return nil, err
	}

	keys := []string{}
	result = result.Next()
	for ; !result.Next().IsZero(); result = result.Next() {
		feature := result.Feature()
		if allowAdd(feature) {
			keys = append(keys, getOrigin(feature))
	}
	}

	return keys, nil
}

func allowAdd(feature string) bool {
	splitedFeature := strings.Split(feature, ",")

	if splitedFeature[0] == "記号" {
		return false
	}

	if splitedFeature[0] == "助動詞" {
		return false
	}

	if splitedFeature[0] == "助詞" {
		return false
	}

	return true
}

func getOrigin(feature string) string {
	splitedFeature := strings.Split(feature, ",")

	return splitedFeature[6]
}
