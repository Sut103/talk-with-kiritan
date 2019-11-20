package main

import (
	"talk-with-kiritan/config"
	"talk-with-kiritan/router"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	dg, r, err := router.InitMainRouter(config)
	if err != nil {
		panic(err)
	}

	if err = dg.Open(); err != nil {
		panic(err)
	}

	if err := r.Run(); err != nil {
		panic(err)
	}
}
