package filediscovery

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// WorkingDirProvider provides the working directory as a possible file location
func WorkingDirProvider() FileLocationProvider {

	return func(fileName string) (string, error) {
		dir, err := os.Getwd()
		if err != nil {
			return "", err
		}

		return path.Join(dir, fileName), nil
	}
}

// ExecutableDirProvider provides the executables directory as a possible file location
func ExecutableDirProvider() FileLocationProvider {

	return func(fileName string) (string, error) {
		dir, err := getExecutablePath()
		if err != nil {
			return "", err
		}

		return path.Join(dir, fileName), nil
	}
}

func getExecutablePath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(ex), nil
}

// EnvVarFilePathProvider provides a filePath in the given environment variable.
// In contrast to other FileLocationProviders, this file location provider expects a complete filePath in the given
// environment variable.
func EnvVarFilePathProvider(envVar string) FileLocationProvider {
	return func(fileName string) (string, error) {
		_ = fileName
		if envConfigFile, ok := os.LookupEnv(envVar); ok {
			return envConfigFile, nil
		}

		return "", fmt.Errorf("env var '%s' not defined", envVar)
	}
}
