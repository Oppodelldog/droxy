package cmd

import (
	"docker-proxy-command/config"
	"docker-proxy-command/helper"
	"docker-proxy-command/symlinks"
	"fmt"
	"os"
	"path"
)

const commandFileName = "docker-proxy"

func CreateSymlinks(cfg *config.Configuration) error {
	executableDir, err := helper.GetExecutablePath()
	if err != nil {
		return err
	}

	commandFilepath := path.Join(executableDir, commandFileName)
	if _, err := os.Stat(commandFilepath); os.IsNotExist(err) {
		return fmt.Errorf("could not find docker-proxy command as expected at '%s'", commandFilepath)
	}

	return symlinks.CreateSymlinks(commandFilepath, cfg)
}
