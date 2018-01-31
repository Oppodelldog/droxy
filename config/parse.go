package config

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

func Parse(filepath string) (*Configuration, error) {
	if fileContent, err := ioutil.ReadFile(filepath); err == nil {
		return parseFromBytes(fileContent)
	} else {
		return nil, err
	}
}

func parseFromBytes(bytes []byte) (*Configuration, error) {
	var conf Configuration

	_, err := toml.Decode(string(bytes), &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
