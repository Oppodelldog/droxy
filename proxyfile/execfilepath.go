package proxyfile

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

var errDroxyCommandNotFound = errors.New("could not find droxy command")

func getExecutableFilePath() (string, error) {
	executableDir, err := getExecutablePath()
	if err != nil {
		return "", err
	}

	commandFilepath := path.Join(executableDir, GetCommandName())
	if _, err := os.Stat(commandFilepath); os.IsNotExist(err) {
		return "", fmt.Errorf("%w as expected at '%s'", errDroxyCommandNotFound, commandFilepath)
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
