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
	proxyFilesCreator := newProxyFilesCreator(proxyfile.NewClonesStrategy(), config.NewLoader())

	return newFileCreationSubCommand("clones", proxyFilesCreator)
}

func newHardlinkCommandWrapper() *fileCreationSubCommandWrapper {
	proxyFilesCreator := newProxyFilesCreator(proxyfile.NewHardlinkStrategy(), config.NewLoader())

	return newFileCreationSubCommand("hardlinks", proxyFilesCreator)
}

func newSymlinkCommandWrapper() *fileCreationSubCommandWrapper {
	proxyFilesCreator := newProxyFilesCreator(proxyfile.NewSymlinkStrategy(), config.NewLoader())

	return newFileCreationSubCommand("symlinks", proxyFilesCreator)
}

type fileCreationSubCommandWrapper struct {
	cobraCommand *cobra.Command
	isForced     bool
}

func (w *fileCreationSubCommandWrapper) getCommand() *cobra.Command {
	return w.cobraCommand
}

func (w *fileCreationSubCommandWrapper) createCommand(
	commandName string,
	proxyFilesCreator ProxyFilesCreator,
) *cobra.Command {
	w.cobraCommand = &cobra.Command{
		Use:   commandName,
		Short: fmt.Sprintf("creates command %s", commandName),
		Long:  `creates clones of droxy for all command in the current directory`,
		Run: func(cmd *cobra.Command, args []string) {
			logrus.Infof("creating '%s'...", commandName)

			err := proxyFilesCreator.CreateProxyFiles(w.isForced)
			if err != nil {
				logrus.Error(err)
				os.Exit(1)
			}
		},
	}

	w.cobraCommand.Flags().BoolVarP(&w.isForced, "force", "f", false, "removes existing files before creation")

	return w.cobraCommand
}

func newFileCreationSubCommand(commandName string, proxyFilesCreator ProxyFilesCreator) *fileCreationSubCommandWrapper {
	commandWrapper := new(fileCreationSubCommandWrapper)
	commandWrapper.createCommand(commandName, proxyFilesCreator)

	return commandWrapper
}

type ProxyFilesCreator interface {
	CreateProxyFiles(isForced bool) error
}

func newProxyFilesCreator(strategy proxyfile.FileCreationStrategy, configLoader config.Loader) ProxyFilesCreator {
	return proxyfile.New(strategy, configLoader)
}
