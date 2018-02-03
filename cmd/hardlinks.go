package cmd

import (
	"fmt"

	"docker-proxy-command/config"
	"docker-proxy-command/helper"
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"docker-proxy-command/proxyfile"
)

var forceHardlinks bool

func NewHardCommand() *cobra.Command {
	hardlinkCommand.Flags().BoolVarP(&forceHardlinks, "force", "f", false, "removes existing files before creation")
	return hardlinkCommand
}

var hardlinkCommand = &cobra.Command{
	Use:   "hardlinks",
	Short: "creates command hardlinks",
	Long:  `creates hardlinks for all command in the current directory`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Load()
		createHardlinks(cfg, forceHardlinks)
	},
}

func createHardlinks(cfg *config.Configuration, isForced bool) error {

	logrus.Info("creating hardlinks...")

	executableDir, err := helper.GetExecutablePath()
	if err != nil {
		return err
	}

	commandFilepath := path.Join(executableDir, commandFileName)
	if _, err := os.Stat(commandFilepath); os.IsNotExist(err) {
		return fmt.Errorf("could not find docker-proxy command as expected at '%s'", commandFilepath)
	}

	profileFileCreator := proxyfile.New(proxyfile.NewHardlinkStrategy())

	return profileFileCreator.CreateProxyFiles(commandFilepath, cfg, isForced)
}
