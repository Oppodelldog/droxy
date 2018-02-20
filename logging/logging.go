package logging

import (
	"github.com/Oppodelldog/droxy/config"
	"io"
	"os"
	"path"
	"path/filepath"
)

const logFileName = "droxy.log"

// GetLogWriter returns a logwriter which is used for debug logs
func GetLogWriter(cfg *config.Configuration) (io.WriteCloser, error) {
	configPath := filepath.Dir(cfg.GetConfigurationFilePath())
	logFilePath := path.Join(configPath, logFileName)
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	return file, nil
}
