package controller

type MainController struct {
	VChs VoiceChannels
}

type VoiceChannels struct {
	Ch   chan string
	Exit chan string
}

func (mctrl *MainController) GetDiscordController() *DiscordController {
	return &DiscordController{Main: mctrl}
}

func (mctrl *MainController) GetServerController() *ServerController {
	return &ServerController{Main: mctrl}
}

func GetMainController() *MainController {
	ch := make(chan string)
	exit := make(chan string)
	vchs := VoiceChannels{Ch: ch, Exit: exit}
	return &MainController{vchs}
}
