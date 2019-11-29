package router

import (
	"talk-with-kiritan/config"
	"talk-with-kiritan/controller"

	"github.com/bwmarrin/discordgo"
)

func InitDiscordRouter(config config.Discord, mctrl *controller.MainController) (*discordgo.Session, error) {
	session, err := discordgo.New()
	if err != nil {
		return nil, err
	}
	session.Token = "Bot " + config.BotToken

	ctrl := mctrl.GetDiscordController()
	ctrl.Config = config
	session.AddHandler(ctrl.MessageRecive)

	return session, nil
}
