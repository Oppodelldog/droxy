package symlinks

import (
	"docker-proxy-command/config"
	"os"

	"github.com/sirupsen/logrus"
)

func CreateSymlinks(commandBinaryFilePath string, configuration *config.Configuration, isForced bool) error {
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

		err := CreateSymlink(commandBinaryFilePath, *command.Name)
		if err != nil {
			logrus.Errorf("error creating symlink '%s': %v", *command.Name, err)
			continue
		}

		logrus.Infof("created '%s'", *command.Name)
	}

	return nil
}

func CreateSymlink(commandBinaryFilePath, commandNameFilePath string) error {
	return os.Symlink(commandBinaryFilePath, commandNameFilePath)
}
