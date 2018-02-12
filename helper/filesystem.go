package helper

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// GetExecutablePath returns the path of the directory where docker-proxy binary is located
func GetExecutablePath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(ex), nil
}

// GetExecutableFilePath returns the file-path of the docker-proxy binary
func GetExecutableFilePath() (string, error) {
	executableDir, err := GetExecutablePath()
	if err != nil {
		return "", err
	}

	commandFilepath := path.Join(executableDir, GetCommandName())
	if _, err := os.Stat(commandFilepath); os.IsNotExist(err) {
		return "", fmt.Errorf("could not find docker-proxy command as expected at '%s'", commandFilepath)
	}

	return commandFilepath, nil
}
