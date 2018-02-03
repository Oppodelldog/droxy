package helper

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

const commandFileName = "docker-proxy"

func GetExecutablePath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(ex), nil
}

func GetExecutableFilePath() (string, error) {
	executableDir, err := GetExecutablePath()
	if err != nil {
		return "", err
	}

	commandFilepath := path.Join(executableDir, commandFileName)
	if _, err := os.Stat(commandFilepath); os.IsNotExist(err) {
		return "", fmt.Errorf("could not find docker-proxy command as expected at '%s'", commandFilepath)
	}

	return commandFilepath, nil
}
