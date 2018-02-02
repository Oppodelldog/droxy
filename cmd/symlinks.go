package cmd

import (
	"fmt"

	"docker-proxy-command/config"
	"docker-proxy-command/helper"
	"docker-proxy-command/symlinks"
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var forceSymlinks bool

func NewSymlinkCommand() *cobra.Command {
	symlinkCommand.Flags().BoolVarP(&forceSymlinks, "force", "f", false, "removes existing files before creation")
	return symlinkCommand
}

var symlinkCommand = &cobra.Command{
	Use:   "symlinks",
	Short: "creates command symlinks",
	Long:  `creates symlinks for all command in the current directory`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Load()
		createSymlinks(cfg, forceSymlinks)
	},
}

const commandFileName = "docker-proxy"

func createSymlinks(cfg *config.Configuration, isForced bool) error {

	logrus.Info("creating symlinks...")

	executableDir, err := helper.GetExecutablePath()
	if err != nil {
		return err
	}

	commandFilepath := path.Join(executableDir, commandFileName)
	if _, err := os.Stat(commandFilepath); os.IsNotExist(err) {
		return fmt.Errorf("could not find docker-proxy command as expected at '%s'", commandFilepath)
	}

	return symlinks.CreateSymlinks(commandFilepath, cfg, isForced)
}
