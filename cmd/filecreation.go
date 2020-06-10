package cmd

import (
	"fmt"
	"os"

	"github.com/Oppodelldog/droxy/config"

	"github.com/Oppodelldog/droxy/proxyfile"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newCloneCommandWrapper() *cobra.Command {
	return createCommand(
		"clones",
		proxyfile.New(
			proxyfile.NewClonesStrategy(),
			config.NewLoader(),
		),
	)
}

func newHardlinkCommandWrapper() *cobra.Command {
	return createCommand(
		"hardlinks",
		proxyfile.New(
			proxyfile.NewHardlinkStrategy(),
			config.NewLoader(),
		),
	)
}

func newSymlinkCommandWrapper() *cobra.Command {
	return createCommand(
		"symlinks",
		proxyfile.New(
			proxyfile.NewSymlinkStrategy(),
			config.NewLoader(),
		),
	)
}

type ProxyFilesCreator interface {
	CreateProxyFiles(isForced bool) error
}

func createCommand(commandName string, proxyFilesCreator ProxyFilesCreator) *cobra.Command {
	cobraCommand := &cobra.Command{
		Use:   commandName,
		Short: fmt.Sprintf("creates command %s", commandName),
		Long:  `creates clones of droxy for all command in the current directory`,
		Run: func(cmd *cobra.Command, args []string) {
			isForced, _ := cmd.Flags().GetBool("force")
			logrus.Infof("creating '%s'...", commandName)

			err := proxyFilesCreator.CreateProxyFiles(isForced)
			if err != nil {
				logrus.Error(err)
				os.Exit(1)
			}
		},
	}
	cobraCommand.Flags().BoolP("force", "f", false, "removes existing files before creation")

	return cobraCommand
}
