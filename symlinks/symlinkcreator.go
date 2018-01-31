package symlinks

import (
	"docker-proxy-command/config"
	"os"
)

func CreateSymlinks(commandBinaryFilePath string, configuration *config.Configuration) error {
	for _, command := range configuration.Command {
		err := CreateSymlink(commandBinaryFilePath, command.Name)
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateSymlink(commandBinaryFilePath, commandNameFilePath string) error {
	return os.Symlink(commandBinaryFilePath, commandNameFilePath)
}
