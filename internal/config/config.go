package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Mastodon MastodonConfig `json:"mastodon"`
}

type MastodonConfig struct {
	Server       string   `json:"server"`
	ClientID     string   `json:"clientId"`
	ClientSecret string   `json:"clientSecret"`
	AccessToken  string   `json:"accessToken"`
	TargetUsers  []string `json:"targetUsers"`
}

// LoadConfigFromFile reads and parses the configuration file
func LoadConfigFromFile(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %w", filename, err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file %s: %w", filename, err)
	}

	return &config, nil
}
