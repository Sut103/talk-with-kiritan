package router

import (
	"talk-with-kiritan/config"
	"talk-with-kiritan/controller"

	"github.com/bwmarrin/discordgo"
)

func InitDiscord(config config.DiscordConfig) (*discordgo.Session, error) {
	session, err := discordgo.New()
	if err != nil {
		return nil, err
	}
	session.Token = "Bot " + config.Token

	dctrl := controller.GetDiscordController()
	session.AddHandler(dctrl.MessageRecive)

	return session, nil
}
