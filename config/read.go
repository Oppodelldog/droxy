package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

func readFromFile(filepath string) (Configuration, error) {
	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		return Configuration{}, err
	}

	cfg, err := parseFromBytes(fileContent)
	if err != nil {
		return Configuration{}, err
	}

	cfg.ConfigFilePath = filepath

	return cfg, nil
}

func parseFromBytes(bytes []byte) (Configuration, error) {
	var conf Configuration

	_, err := toml.Decode(string(bytes), &conf)

	return conf, err
}
