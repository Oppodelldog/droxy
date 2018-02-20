package config

import (
	"bytes"
	"github.com/Oppodelldog/droxy/helper"
	"fmt"
	"os"
	"path"

	"github.com/pkg/errors"
)

const configFileName = "droxy.toml"

// DiscoverConfigFile tries to find a config file in different locations:
// * current working directory
// * directory the executable is located in
// * filepath from an environment variable
func DiscoverConfigFile() (string, error) {

	var possibleConfigFilePaths []string
	var configFileProviders []func() (string, error)

	errorString := bytes.NewBufferString("")

	configFileProviders = append(configFileProviders, workingDirProvider)
	configFileProviders = append(configFileProviders, executableDirProvider)
	configFileProviders = append(configFileProviders, envVarFilePathProvider)

	for _, configFileProvider := range configFileProviders {
		possibleConfigFile, err := configFileProvider()
		if err != nil {
			errorString.WriteString(err.Error())
			errorString.WriteString("\n")
		}

		possibleConfigFilePaths = append(possibleConfigFilePaths, possibleConfigFile)
	}

	for _, configFilePath := range possibleConfigFilePaths {
		if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
			errorString.WriteString(fmt.Sprintf("could not find config file at '%s'\n", configFilePath))
		} else {
			return configFilePath, nil
		}
	}

	return "", errors.New(errorString.String())
}

func workingDirProvider() (string, error) {

	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return path.Join(dir, configFileName), nil
}

func executableDirProvider() (string, error) {

	dir, err := helper.GetExecutablePath()
	if err != nil {
		return "", err
	}

	return path.Join(dir, configFileName), nil
}

func envVarFilePathProvider() (string, error) {
	if envConfigFile, ok := os.LookupEnv("DROXY_CONFIG"); ok {
		return envConfigFile, nil
	}

	return "", errors.New("env var DROXY_CONFIG not defined")
}
