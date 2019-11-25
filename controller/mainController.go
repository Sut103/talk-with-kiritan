package controller

import "sync"

type MainController struct {
	VChs  VoiceConnection
	Timer Timer
}

type VoiceConnection struct {
	Ch chan string
}

type Timer struct {
	AllowSend bool
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
	vchs := VoiceConnection{Ch: ch}
	timer := Timer{}
	return &MainController{VChs: vchs, Timer: timer}
}
