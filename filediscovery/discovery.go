package filediscovery

import (
	"bytes"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

type (
	// FileDiscovery defines logic to discover a file.
	FileDiscovery interface {
		Discover(fileName string) (string, error)
	}

	fileDiscovery struct {
		fileLocationProviders []FileLocationProvider
	}

	// FileLocationProvider provides a possible file location to FileDiscovery
	FileLocationProvider func(fileName string) (string, error)
)

// New creates a new FileDiscovery and takes a list of FileLocationProviders which specify possible location a given file
// will be searched in.
func New(fileLocationProviders []FileLocationProvider) FileDiscovery {
	return &fileDiscovery{
		fileLocationProviders: fileLocationProviders,
	}
}

// Discover tries to find the given fileName in all FileLocationProviders. The providers are checked in given sequence.
// the first matching result will be returned. If the file could not be found and error is returned as if any other
// error occurs.
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
