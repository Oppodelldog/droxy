package filediscovery

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func WorkingDirProvider() FileLocationProvider {

	return func(fileName string) (string, error) {
		dir, err := os.Getwd()
		if err != nil {
			return "", err
		}

		return path.Join(dir, fileName), nil
	}
}

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

func EnvVarFilePathProvider(envVar string) FileLocationProvider {
	return func(fileName string) (string, error) {
		_ = fileName
		if envConfigFile, ok := os.LookupEnv(envVar); ok {
			return envConfigFile, nil
		}

		return "", fmt.Errorf("env var '%s' not defined", envVar)
	}
}
