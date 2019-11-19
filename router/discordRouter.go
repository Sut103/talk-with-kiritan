package router

import "github.com/bwmarrin/discordgo"

import "talk-with-kiritan/config"

func InitDiscord(config config.DiscordConfig) (*discordgo.Session, error) {
	session, err := discordgo.New()
	if err != nil {
		return nil, err
	}

	session.Token = "Bot " + config.Token

	return session, nil
}
