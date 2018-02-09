package config

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

// Parse parses a toml configuration file into Configuration data model
func Parse(filepath string) (*Configuration, error) {
	fileContent, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return parseFromBytes(fileContent)
}

func parseFromBytes(bytes []byte) (*Configuration, error) {
	var conf Configuration

	_, err := toml.Decode(string(bytes), &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
