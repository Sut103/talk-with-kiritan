package router

import "github.com/bwmarrin/discordgo"

func InitDiscord() (*discordgo.Session, error) {
	dg, err := discordgo.New()
	if err != nil {
		return nil, err
	}
	return dg, nil
}
