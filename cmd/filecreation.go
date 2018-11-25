package cmd

import (
	"fmt"
	"github.com/Oppodelldog/droxy/config"
	"os"

	"github.com/Oppodelldog/droxy/proxyfile"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newCloneCommandWrapper() *fileCreationSubCommandWrapper {
	return newFileCreationSubCommand("clones", proxyfile.NewClonesStrategy(), config.NewLoader())
}

func newHardlinkCommandWrapper() *fileCreationSubCommandWrapper {
	return newFileCreationSubCommand("hardlinks", proxyfile.NewHardlinkStrategy(), config.NewLoader())
}

func newSymlinkCommandWrapper() *fileCreationSubCommandWrapper {
	return newFileCreationSubCommand("symlinks", proxyfile.NewSymlinkStrategy(), config.NewLoader())
}

type fileCreationSubCommandWrapper struct {
	cobraCommand *cobra.Command
	isForced     bool
}

func (w *fileCreationSubCommandWrapper) getCommand() *cobra.Command {
	return w.cobraCommand
}

func (w *fileCreationSubCommandWrapper) createCommand(commandName string, strategy proxyfile.FileCreationStrategy, configLoader config.Loader) *cobra.Command {
	w.cobraCommand = &cobra.Command{
		Use:   commandName,
		Short: fmt.Sprintf("creates command %s", commandName),
		Long:  `creates clones of droxy for all command in the current directory`,
		Run: func(cmd *cobra.Command, args []string) {
			logrus.Infof("creating '%s'...", commandName)

			profileFileCreator := proxyfile.New(strategy, configLoader)

			err := profileFileCreator.CreateProxyFiles(w.isForced)
			if err != nil {
				logrus.Error(err)
				os.Exit(1)
			}
		},
	}

	w.cobraCommand.Flags().BoolVarP(&w.isForced, "force", "f", false, "removes existing files before creation")

	return w.cobraCommand
}

func newFileCreationSubCommand(commandName string, strategy proxyfile.FileCreationStrategy, configLoader config.Loader) *fileCreationSubCommandWrapper {

	commandWrapper := new(fileCreationSubCommandWrapper)
	commandWrapper.createCommand(commandName, strategy, configLoader)

	return commandWrapper
}
