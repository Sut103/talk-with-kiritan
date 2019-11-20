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

	ctrl := mctrl.GetDiscordController()
	session.AddHandler(ctrl.MessageRecive)

	return session, nil
}
