package controller

type MainController struct {
	VChs VoiceChannels
}

type VoiceChannels struct {
	Ch   chan string
	Exit chan string
}

func GetDiscordController(mctrl *MainController) *DiscordController {
	return &DiscordController{Main: mctrl}
}

func GetServerController(mctrl *MainController) *ServerController {
	return &ServerController{Main: mctrl}
}

func GetMainController() *MainController {
	ch := make(chan string)
	exit := make(chan string)
	return &MainController{Ch: ch, Exit: exit}
}
