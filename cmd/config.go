package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path"
)

const configFileName = "docker-proxy.yml"

func GetConfigFilePath() (string, error) {

	errorString := bytes.NewBufferString("")
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	configFilePath := path.Join(dir, configFileName)

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		errorString.WriteString(fmt.Sprintf("could not find config file at '%s'\n", configFilePath))
	} else {
		return configFilePath, nil
	}

	if envConfigFile, ok := os.LookupEnv("DOCKER_PROXY_CONFIG"); !ok {
		errorString.WriteString(fmt.Sprintf("no env var 'DOCKER_PROXY_CONFIG' set\n"))
	} else {
		if _, err := os.Stat(envConfigFile); os.IsNotExist(err) {
			errorString.WriteString(fmt.Sprintf("could not find config file (defined by env var DOCKER_PROXY_CONFIG) at '%s'\n", envConfigFile))
		} else {
			return envConfigFile, nil
		}
	}

	var finalError error

	if errorString.Len() > 0 {
		finalError = errors.New(errorString.String())
	}
	return "", finalError
}
