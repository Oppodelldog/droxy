package helper

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// GetExecutablePath returns the path of the directory where droxy binary is located
func GetExecutablePath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(ex), nil
}

// GetExecutableFilePath returns the file-path of the droxy binary
func GetExecutableFilePath() (string, error) {
	executableDir, err := GetExecutablePath()
	if err != nil {
		return "", err
	}

	commandFilepath := path.Join(executableDir, GetCommandName())
	if _, err := os.Stat(commandFilepath); os.IsNotExist(err) {
		return "", fmt.Errorf("could not find droxy command as expected at '%s'", commandFilepath)
	}

	return commandFilepath, nil
}
