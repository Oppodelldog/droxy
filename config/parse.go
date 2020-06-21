package config

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

// Parse parses a toml configuration file into Configuration data model.
func Parse(filepath string) (Configuration, error) {
	fileContent, err := ioutil.ReadFile(filepath)
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
