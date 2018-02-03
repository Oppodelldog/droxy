package cmd

import (
	"docker-proxy-command/config"
	"docker-proxy-command/helper"
	"docker-proxy-command/proxyfile"
	"fmt"

	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewCloneCommandWrapper() *FileCreationSubCommandWrapper {
	return NewFileCreationSubCommand("clones", proxyfile.NewClonesStrategy())
}
func NewHardlinkCommandWrapper() *FileCreationSubCommandWrapper {
	return NewFileCreationSubCommand("hardlinks", proxyfile.NewHardlinkStrategy())
}
func NewSymlinkCommandWrapper() *FileCreationSubCommandWrapper {
	return NewFileCreationSubCommand("symlinks", proxyfile.NewHardlinkStrategy())
}

type FileCreationSubCommandWrapper struct {
	cobraCommand *cobra.Command
	isForced     bool
}

func (w *FileCreationSubCommandWrapper) GetCommand() *cobra.Command {
	return w.cobraCommand
}

func (w *FileCreationSubCommandWrapper) CreateCommand(commandName string, strategy proxyfile.FileCreationStrategy) *cobra.Command {
	w.cobraCommand = &cobra.Command{
		Use:   commandName,
		Short: fmt.Sprintf("creates command %s", commandName),
		Long:  `creates clones of docker-proxy for all command in the current directory`,
		Run: func(cmd *cobra.Command, args []string) {
			cfg := config.Load()

			logrus.Infof("creating '%s'...", commandName)

			commandFilePath, err := helper.GetExecutableFilePath()
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

func NewFileCreationSubCommand(commandName string, strategy proxyfile.FileCreationStrategy) *FileCreationSubCommandWrapper {

	commandWrapper := new(FileCreationSubCommandWrapper)
	commandWrapper.CreateCommand(commandName, strategy)

	return commandWrapper
}
