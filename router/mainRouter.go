package router

import (
	"fmt"
	"io/ioutil"
	"strings"
	"talk-with-kiritan/config"
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
