package helper

import (
	"os"
	"path/filepath"
)

func GetExecutablePath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(ex), nil
}
func GetExecutableInfo() (string, string) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fileInfo, err := os.Stat(ex)
	if err != nil {
		panic(err)
	}

	return fileInfo.Name(), exPath
}
