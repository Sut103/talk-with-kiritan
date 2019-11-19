package controller

import (
	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

type DiscordController struct {
	VC   *discordgo.VoiceConnection
	Main *MainController
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

	//VC参加
	if event.Content == "きりたん砲全門斉射！" {
		for _, vs := range discordGuild.VoiceStates {
			if vs.UserID == event.Author.ID {
				dctrl.VC, err = s.ChannelVoiceJoin(discordGuild.ID, vs.ChannelID, false, true)
			}
			if err != nil {
				panic(err)
			}
			s.ChannelMessageSend(discordChannel.ID, "きりたん砲の味噌となれっ!!")

			//音声ファイルのリクエスト受付を開始
			go playVoiceRoop(s, dctrl)

		}
	}

	//VC退出
	if event.Content == "おつかれさまです" {
		s.ChannelMessageSend(discordChannel.ID, "それでは")
		if err := dctrl.VC.Disconnect(); err != nil {
			panic(err)
		}
		dctrl.Main.Exit <- "exit"
	}
}

func playVoiceRoop(s *discordgo.Session, dctrl *DiscordController) {
	for {
		select {
		case soundFilename := <-dctrl.Main.Ch:
			dgvoice.PlayAudioFile(dctrl.VC, "sounds/"+soundFilename, make(<-chan bool))

		case status := <-dctrl.Main.Exit:
			if status == "exit" {
				break
			}
		}
	}
}
