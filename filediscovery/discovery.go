package filediscovery

import (
	"bytes"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

type (
	FileDiscovery interface {
		Discover(fileName string) (string, error)
	}

	fileDiscovery struct {
		fileLocationProviders []FileLocationProvider
	}

	FileLocationProvider func(fileName string) (string, error)
)

func New(fileLocationProviders []FileLocationProvider) FileDiscovery {
	return &fileDiscovery{
		fileLocationProviders: fileLocationProviders,
	}
}

func (fd *fileDiscovery) Discover(fileName string) (string, error) {

	var possibleConfigFilePaths []string

	errorString := bytes.NewBufferString("")

	for _, configFileProvider := range fd.fileLocationProviders {
		possibleConfigFile, err := configFileProvider(fileName)
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
