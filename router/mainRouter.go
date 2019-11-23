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

type TaggerNode struct {
	Surface       string
	PartsOfSpeech string
	Origin        string
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

	noSymbolFileName := ""
	keys := []string{}
	result = result.Next()
	for ; !result.Next().IsZero(); result = result.Next() {
		resultStruct := nodeToStruct(result)
		if !resultStruct.isSymbol() { //記号を取り除いた元ファイル名の生成
			noSymbolFileName += resultStruct.getSurface()
		}
		if resultStruct.allowAdd() {
			keys = append(keys, resultStruct.getOrigin())
		}
	}

	noSokuonFileName := trimSokuon(noSymbolFileName)
	noLongVowelFileName := trimLongVowel(noSokuonFileName)
	keys = append(keys, noSymbolFileName)
	keys = append(keys, noSokuonFileName)
	keys = append(keys, noLongVowelFileName)

	return keys, nil
}

func (t TaggerNode) allowAdd() bool {

	if t.PartsOfSpeech == "記号" {
		return false
	}

	if t.PartsOfSpeech == "助動詞" {
		return false
	}

	if t.PartsOfSpeech == "助詞" {
		return false
	}

	return true
}

func (t TaggerNode) getOrigin() string {
	return t.Origin
}

func (t TaggerNode) getSurface() string {
	return t.Surface
}

func nodeToStruct(node mecab.Node) TaggerNode {
	ret := TaggerNode{}

	ret.Surface = node.Surface()

	feature := node.Feature()
	splitedFeature := strings.Split(feature, ",")
	ret.PartsOfSpeech = splitedFeature[0]
	ret.Origin = splitedFeature[6]

	return ret
}

func (t TaggerNode) isSymbol() bool {
	if t.PartsOfSpeech == "記号" {
		return true
	}
	return false
}

func trimLongVowel(fileName string) string {
	return strings.ReplaceAll(fileName, "ー", "")
}

func trimSokuon(fileName string) string {
	return strings.ReplaceAll(fileName, "っ", "")
}
