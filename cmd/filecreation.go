package cmd

import (
	"fmt"
	"os"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/proxyfile"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newCloneCommandWrapper() *fileCreationSubCommandWrapper {
	return newFileCreationSubCommand("clones", proxyfile.NewClonesStrategy())
}

func newHardlinkCommandWrapper() *fileCreationSubCommandWrapper {
	return newFileCreationSubCommand("hardlinks", proxyfile.NewHardlinkStrategy())
}

func newSymlinkCommandWrapper() *fileCreationSubCommandWrapper {
	return newFileCreationSubCommand("symlinks", proxyfile.NewSymlinkStrategy())
}

type fileCreationSubCommandWrapper struct {
	cobraCommand *cobra.Command
	isForced     bool
}

func (w *fileCreationSubCommandWrapper) getCommand() *cobra.Command {
	return w.cobraCommand
}

func (w *fileCreationSubCommandWrapper) createCommand(commandName string, strategy proxyfile.FileCreationStrategy) *cobra.Command {
	w.cobraCommand = &cobra.Command{
		Use:   commandName,
		Short: fmt.Sprintf("creates command %s", commandName),
		Long:  `creates clones of droxy for all command in the current directory`,
		Run: func(cmd *cobra.Command, args []string) {
			cfg := config.NewLoader().Load()

			logrus.Infof("creating '%s'...", commandName)

			commandFilePath, err := proxyfile.GetExecutableFilePath()
			if err != nil {
				logrus.Error(err)
				os.Exit(1)
			}

			profileFileCreator := proxyfile.New(strategy)

			err = profileFileCreator.CreateProxyFiles(commandFilePath, cfg, w.isForced)
			if err != nil {
				logrus.Error(err)
				os.Exit(1)
			}
		},
	}

	w.cobraCommand.Flags().BoolVarP(&w.isForced, "force", "f", false, "removes existing files before creation")

	return w.cobraCommand
}

func newFileCreationSubCommand(commandName string, strategy proxyfile.FileCreationStrategy) *fileCreationSubCommandWrapper {

	commandWrapper := new(fileCreationSubCommandWrapper)
	commandWrapper.createCommand(commandName, strategy)

	return commandWrapper
}
