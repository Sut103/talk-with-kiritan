package controller

import (
	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	"talk-with-kiritan/config"
	"time"
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
			//参加メッセージ
			s.ChannelMessageSend(discordChannel.ID, ctrl.Config.MessageJoin)

			//音声ファイルのリクエスト受付を開始
			go playAudioLoop(s, ctrl)

			//　参加時Voice
			if voice := ctrl.Config.AudioFileJoin; voice != "" {
				time.Sleep(time.Second * 2)
				dgvoice.PlayAudioFile(ctrl.VC, "sounds/"+voice, make(<-chan bool))
				time.Sleep(time.Second * 2)
			}

			//音声ファイル名の送信を許可
			ctrl.Main.VChs.Lock.Lock()
			ctrl.Main.VChs.Condition = true
			ctrl.Main.VChs.Lock.Unlock()
		}
	}

	//VC退出
	if event.Content == ctrl.Config.OrderMessageLeave {
		//音声ファイル名の送信を遮断
		ctrl.Main.VChs.Lock.Lock()
		ctrl.Main.VChs.Condition = false
		ctrl.Main.VChs.Lock.Unlock()

		//退出時Voice
		if voice := ctrl.Config.AudioFileLeave; voice != "" {
			time.Sleep(time.Second * 2)
			dgvoice.PlayAudioFile(ctrl.VC, "sounds/"+voice, make(<-chan bool))
			time.Sleep(time.Second * 2)
		}

		//退出メッセージ
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
