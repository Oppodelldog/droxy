package proxyfile

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func getExecutableFilePath() (string, error) {
	executableDir, err := getExecutablePath()
	if err != nil {
		return "", err
	}

	commandFilepath := path.Join(executableDir, GetCommandName())
	if _, err := os.Stat(commandFilepath); os.IsNotExist(err) {
		return "", fmt.Errorf("could not find droxy command as expected at '%s'", commandFilepath)
	}

	return commandFilepath, nil
}

func getExecutablePath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}

	return filepath.Dir(ex), nil
}
