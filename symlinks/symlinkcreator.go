package symlinks

import (
	"docker-proxy-command/config"
	"os"
	"fmt"
)

func CreateSymlinks(commandBinaryFilePath string, configuration *config.Configuration) error {
	for _, command := range configuration.Command {
		fmt.Printf(" - %s: ", command.Name)
		err := CreateSymlink(commandBinaryFilePath, command.Name)
		if err != nil {
			return err
		}
		fmt.Println("OK")
	}

	return nil
}

func CreateSymlink(commandBinaryFilePath, commandNameFilePath string) error {
	return os.Symlink(commandBinaryFilePath, commandNameFilePath)
}
