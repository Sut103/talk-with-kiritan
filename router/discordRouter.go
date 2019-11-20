package router

import (
	"talk-with-kiritan/config"
	"talk-with-kiritan/controller"

	"github.com/bwmarrin/discordgo"
)

func InitDiscordRouter(config config.DiscordConfig, mctrl *controller.MainController) (*discordgo.Session, error) {
	session, err := discordgo.New()
	if err != nil {
		return nil, err
	}
	session.Token = "Bot " + config.Token

	dctrl := controller.GetDiscordController(mctrl)
	session.AddHandler(dctrl.MessageRecive)

	return session, nil
}
