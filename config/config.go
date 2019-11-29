package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Discord Discord `json:"discord"`
	Server  Server  `json:"server"`
}
type Discord struct {
	BotToken          string `json:"bot_token"`
	OrderMessageJoin  string `json:"order_message_join"`
	OrderMessageLeave string `json:"order_message_leave"`
	MessageJoin       string `json:"message_join"`
	MessageLeave      string `json:"message_leave"`
	AudioFileJoin     string `json:"audio_file_join"`
	AudioFileLeave    string `json:"audio_file_leave"`
}
type Users struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}
type Server struct {
	AudioFileExtension string  `json:"audio_file_extension"`
	Idle               int     `json:"idle"`
	Users              []Users `json:"users"`
}

func GetConfig() (Config, error) {
	config := Config{}

	file, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		return config, err
	}

	if err := json.Unmarshal(file, &config); err != nil {
		return config, err
	}

	return config, nil
}
