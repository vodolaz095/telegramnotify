package commands

import (
	"encoding/json"
	"io/ioutil"
)

const defaultFilePermissions = 0600

// Sink depicts single pair of bot token and telegram channel we can send messages too
type Sink struct {
	Token  string `json:"token"`
	ChatID int64  `json:"chatID"`
}

// Config depicts manifold of sinks to be used by notification system
type Config map[string]Sink

// LoadConfigFromFile loads config from file
func LoadConfigFromFile(pathToFile string) (cfg Config, err error) {
	data, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &cfg)
	return
}

// Save saves config to file
func (cfg *Config) Save(pathToFile string) (err error) {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return
	}
	err = ioutil.WriteFile(pathToFile, data, defaultFilePermissions)
	return
}
