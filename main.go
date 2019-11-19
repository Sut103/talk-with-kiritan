package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"talk-with-kiritan/router"
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
	dg, err := router.InitDiscord()
	if err != nil {
		panic(err)
	}

	if err = dg.Open(); err != nil {
		panic(err)
	}

	r := router.InitServer()
	if err := r.Run(); err != nil {
		panic(err)
	}
}
