package config

import (
	"encoding/json"
	"io/ioutil"
)

type DiscordConfig struct {
	Token string `json:"token"`
}
type Config struct {
	Discord DiscordConfig `json:"discord"`
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
