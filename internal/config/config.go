package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(user string) error {
	if c.CurrentUserName == user {
		return nil
	}
	c.CurrentUserName = user
	return write(*c)
}

func Read() (Config, error) {
	configFileDir, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(configFileDir)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func getConfigFilePath() (string, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configFileDir := filepath.Join(userHome, configFileName)

	return configFileDir, nil
}

func write(cfg Config) error {
	configFileDir, err := getConfigFilePath()
	if err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(cfg, "", "\t")
	if err != nil {
		return err
	}

	err = os.WriteFile(configFileDir, jsonData, 0755)
	if err != nil {
		return err
	}
	return nil
}
