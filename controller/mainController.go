package controller

import "sync"

type MainController struct {
	VChs VoiceConnection
}

type VoiceConnection struct {
	Ch        chan string
	Condition bool
	Lock      sync.Mutex
}

func (mctrl *MainController) GetDiscordController() *DiscordController {
	return &DiscordController{Main: mctrl}
}

func (mctrl *MainController) GetServerController() *ServerController {
	return &ServerController{Main: mctrl}
}

func GetMainController() *MainController {
	ch := make(chan string)
	vchs := VoiceConnection{Ch: ch, Condition: false}
	return &MainController{VChs: vchs}
}
