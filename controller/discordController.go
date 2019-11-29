package controller

import (
	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	"talk-with-kiritan/config"
)

type DiscordController struct {
	VC     *discordgo.VoiceConnection
	Main   *MainController
	close  chan bool
	Config config.Discord
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
	if event.Content == ctrl.Config.OrderMessageJoin {
		for _, vs := range discordGuild.VoiceStates {
			if vs.UserID == event.Author.ID {
				ctrl.VC, err = s.ChannelVoiceJoin(discordGuild.ID, vs.ChannelID, false, true)
				if err != nil {
					panic(err)
				}
			}
		}
		if ctrl.VC != nil {
			s.ChannelMessageSend(discordChannel.ID, ctrl.Config.MessageJoin)

			//音声ファイルのリクエスト受付を開始
			go playAudioLoop(s, ctrl)

			ctrl.Main.VChs.Lock.Lock()
			ctrl.Main.VChs.Condition = true //音声ファイル名の送信を許可
			ctrl.Main.VChs.Lock.Unlock()
		}
	}

	//VC退出
	if event.Content == ctrl.Config.OrderMessageLeave {
		ctrl.Main.VChs.Lock.Lock()
		ctrl.Main.VChs.Condition = false //音声ファイル名の送信を遮断
		ctrl.Main.VChs.Lock.Unlock()

		s.ChannelMessageSend(discordChannel.ID, ctrl.Config.MessageLeave)
		if err := ctrl.VC.Disconnect(); err != nil {
			panic(err)
		}
		ctrl.close <- true
	}
}

func playAudioLoop(s *discordgo.Session, ctrl *DiscordController) {
	for {
		select {
		case soundFilename := <-ctrl.Main.VChs.Ch:
			dgvoice.PlayAudioFile(ctrl.VC, "sounds/"+soundFilename, make(<-chan bool))

		case <-ctrl.close:
			break
		}
	}
}
