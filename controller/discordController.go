package controller

import (
	"github.com/bwmarrin/discordgo"
)

type DiscordController struct {
	VC *discordgo.VoiceConnection
}

func GetDiscordController() *DiscordController {
	return &DiscordController{}
}

func (dctrl *DiscordController) MessageRecive(s *discordgo.Session, event *discordgo.MessageCreate) {
	discordChannel, err := s.State.Channel(event.ChannelID)
	if err != nil {
		panic(err)
	}

	discordGuild, err := s.State.Guild(discordChannel.GuildID)
	if err != nil {
		panic(err)
	}

	if event.Content == "きりたん砲全門斉射！" { //VC参加
		for _, vs := range discordGuild.VoiceStates {
			if vs.UserID == event.Author.ID {
				dctrl.VC, err = s.ChannelVoiceJoin(discordGuild.ID, vs.ChannelID, false, true)
			}
			if err != nil {
				panic(err)
			}
			s.ChannelMessageSend(discordChannel.ID, "きりたん砲の味噌となれっ!!")
		}
	}
}
