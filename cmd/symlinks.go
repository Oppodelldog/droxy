package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"docker-proxy-command/helper"
	"path"
	"os"
	"docker-proxy-command/symlinks"
	"docker-proxy-command/config"
	"github.com/sirupsen/logrus"
)

var Symlinks = &cobra.Command{
	Use:   "symlinks",
	Short: "creates command symlinks",
	Long:  `creates symlinks for all command in the current directory`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Load()
		createSymlinks(cfg)
	},
}

const commandFileName = "docker-proxy"

func createSymlinks(cfg *config.Configuration) error {

	logrus.Info("creating symlinks...")

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
