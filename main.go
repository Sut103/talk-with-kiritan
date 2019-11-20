package main

import (
	"talk-with-kiritan/config"
	"talk-with-kiritan/controller"
	"talk-with-kiritan/router"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	mctrl := controller.GetMainController()

	dg, err := router.InitDiscord(config.Discord, mctrl)
	if err != nil {
		panic(err)
	}

	if err = dg.Open(); err != nil {
		panic(err)
	}

	r := router.InitServer(fileNames, mctrl)
	if err := r.Run(); err != nil {
		panic(err)
	}
}
