package clones

import (
	"docker-proxy-command/config"
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func CreateClones(commandBinaryFilePath string, configuration *config.Configuration, isForced bool) error {
	for _, command := range configuration.Command {

		if !command.HasPropertyName() {
			logrus.Warnf("skipped command because name is missing!")
			continue
		}

		if command.HasPropertyIsTemplate() && *command.IsTemplate {
			continue
		}

		if _, err := os.Stat(*command.Name); err == nil {
			if isForced {
				err := os.Remove(*command.Name)
				if err != nil {
					panic(err)
				}
			} else {
				logrus.Warnf("command symlink already exists for '%s'", *command.Name)
				continue
			}
		}

		err := createClone(commandBinaryFilePath, *command.Name)
		if err != nil {
			logrus.Errorf("error creating symlink '%s': %v", *command.Name, err)
			continue
		}

		logrus.Infof("created '%s'", *command.Name)
	}

	return nil
}

func createClone(commandBinaryFilePath, commandNameFilePath string) error {

	cleanSrc := filepath.Clean(commandBinaryFilePath)
	cleanDst := filepath.Clean(commandNameFilePath)
	if cleanSrc == cleanDst {
		return nil
	}
	sf, err := os.Open(cleanSrc)
	if err != nil {
		return err
	}
	defer sf.Close()
	if err := os.Remove(cleanDst); err != nil && !os.IsNotExist(err) {
		return err
	}
	df, err := os.OpenFile(cleanDst, os.O_CREATE|os.O_WRONLY, 0766)
	if err != nil {
		return err
	}
	defer df.Close()

	_, err = io.Copy(df, sf)

	return err

}
