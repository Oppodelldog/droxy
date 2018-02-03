package logging

import (
	"docker-proxy-command/config"
	"io"
	"os"
	"path"
	"path/filepath"
)

const logFileName = "docker-proxy.log"

func GetLogWriter(cfg *config.Configuration) (io.WriteCloser, error) {
	configPath := filepath.Dir(cfg.GetConfigurationFilePath())
	logFilePath := path.Join(configPath, logFileName)
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	return file, nil
}
