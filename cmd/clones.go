package cmd

import (
	"fmt"

	"docker-proxy-command/config"
	"docker-proxy-command/helper"
	"os"
	"path"

	"docker-proxy-command/clones"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var forceClones bool

func NewCloneCommand() *cobra.Command {
	cloneCommand.Flags().BoolVarP(&forceClones, "force", "f", false, "removes existing files before creation")
	return cloneCommand
}

var cloneCommand = &cobra.Command{
	Use:   "clones",
	Short: "creates command clones",
	Long:  `creates clones of docker-proxy for all command in the current directory`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Load()
		createClones(cfg, forceClones)
	},
}

func createClones(cfg *config.Configuration, isForced bool) error {

	logrus.Info("creating clones...")

	executableDir, err := helper.GetExecutablePath()
	if err != nil {
		return err
	}

	commandFilepath := path.Join(executableDir, commandFileName)
	if _, err := os.Stat(commandFilepath); os.IsNotExist(err) {
		return fmt.Errorf("could not find docker-proxy command as expected at '%s'", commandFilepath)
	}

	return clones.CreateClones(commandFilepath, cfg, isForced)
}
