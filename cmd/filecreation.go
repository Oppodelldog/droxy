package cmd

import (
	"fmt"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/helper"
	"github.com/Oppodelldog/droxy/proxyfile"

	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// NewCloneCommandWrapper creates and returns clones sub-command
func NewCloneCommandWrapper() *FileCreationSubCommandWrapper {
	return NewFileCreationSubCommand("clones", proxyfile.NewClonesStrategy())
}

// NewHardlinkCommandWrapper creates and returns hardlinks sub-command
func NewHardlinkCommandWrapper() *FileCreationSubCommandWrapper {
	return NewFileCreationSubCommand("hardlinks", proxyfile.NewHardlinkStrategy())
}

// NewSymlinkCommandWrapper creates and returns symlinks sub-command
func NewSymlinkCommandWrapper() *FileCreationSubCommandWrapper {
	return NewFileCreationSubCommand("symlinks", proxyfile.NewSymlinkStrategy())
}

// FileCreationSubCommandWrapper wraps a cobra command fields to hold parse flag options
type FileCreationSubCommandWrapper struct {
	cobraCommand *cobra.Command
	isForced     bool
}

// GetCommand returns the wrapped cobra command
func (w *FileCreationSubCommandWrapper) GetCommand() *cobra.Command {
	return w.cobraCommand
}

func (w *FileCreationSubCommandWrapper) createCommand(commandName string, strategy proxyfile.FileCreationStrategy) *cobra.Command {
	w.cobraCommand = &cobra.Command{
		Use:   commandName,
		Short: fmt.Sprintf("creates command %s", commandName),
		Long:  `creates clones of droxy for all command in the current directory`,
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

// NewFileCreationSubCommand creates a subcommand which will execute proxy file creation
func NewFileCreationSubCommand(commandName string, strategy proxyfile.FileCreationStrategy) *FileCreationSubCommandWrapper {

	commandWrapper := new(FileCreationSubCommandWrapper)
	commandWrapper.createCommand(commandName, strategy)

	return commandWrapper
}
