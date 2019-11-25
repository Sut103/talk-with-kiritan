package controller

type MainController struct {
	VChs  VoiceConnection
	Clock bool
}

type VoiceConnection struct {
	Ch chan string
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
	clock := false
	return &MainController{VChs: vchs, Clock: clock}
}
