package symlinks

import (
	"docker-proxy-command/config"
	"os"

	"github.com/sirupsen/logrus"
)

func CreateSymlinks(commandBinaryFilePath string, configuration *config.Configuration) error {
	for _, command := range configuration.Command {

		if !command.HasPropertyName() {
			logrus.Warnf("skipped command because name is missing!")
			continue
		}

		if command.HasPropertyIsTemplate() && *command.IsTemplate {
			continue
		}

		err := CreateSymlink(commandBinaryFilePath, *command.Name)
		if err != nil {
			return err
		}
		logrus.Infof("created '%s'", *command.Name)
	}

	return nil
}

func CreateSymlink(commandBinaryFilePath, commandNameFilePath string) error {
	return os.Symlink(commandBinaryFilePath, commandNameFilePath)
}
