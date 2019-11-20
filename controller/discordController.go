package controller

import (
	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

type DiscordController struct {
	VC   *discordgo.VoiceConnection
	Main *MainController
}

func (ctrl *DiscordController) MessageRecive(s *discordgo.Session, event *discordgo.MessageCreate) {
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
				ctrl.VC, err = s.ChannelVoiceJoin(discordGuild.ID, vs.ChannelID, false, true)
			}
			if err != nil {
				panic(err)
			}
			s.ChannelMessageSend(discordChannel.ID, "きりたん砲の味噌となれっ!!")

			//音声ファイルのリクエスト受付を開始
			go playAudioLoop(s, ctrl)

		}
	}

	//VC退出
	if event.Content == "おつかれさまです" {
		s.ChannelMessageSend(discordChannel.ID, "それでは")
		if err := ctrl.VC.Disconnect(); err != nil {
			panic(err)
		}
		ctrl.Main.VChs.Exit <- "exit"
	}
}

func playAudioLoop(s *discordgo.Session, ctrl *DiscordController) {
	for {
		select {
		case soundFilename := <-ctrl.Main.VChs.Ch:
			dgvoice.PlayAudioFile(ctrl.VC, "sounds/"+soundFilename, make(<-chan bool))

		case status := <-ctrl.Main.VChs.Exit:
			if status == "exit" {
				break
			}
		}
	}
}
