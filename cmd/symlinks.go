package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"docker-proxy-command/helper"
	"path"
	"os"
	"docker-proxy-command/config"
	"github.com/sirupsen/logrus"
	"docker-proxy-command/symlinks"
)

var force bool

func NewCommand() *cobra.Command {
	symlinkCommand.Flags().BoolVarP(&force, "force", "f", false, "removes existing files before creation")
	return symlinkCommand
}

var symlinkCommand = &cobra.Command{
	Use:   "symlinks",
	Short: "creates command symlinks",
	Long:  `creates symlinks for all command in the current directory`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Load()
		createSymlinks(cfg, force)
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

	return symlinks.CreateSymlinks(commandFilepath, cfg,isForced)
}
