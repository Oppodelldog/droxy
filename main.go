package main

import (
	"docker-proxy-command/cmd"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"docker-proxy-command/helper"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = cmd.Root
	cmd.Root.AddCommand(cmd.NewSymlinkCommand())
	cmd.Root.AddCommand(cmd.NewCloneCommand())
	cmd.Root.AddCommand(cmd.NewHardCommand())

	if len(os.Args) >= 2 && isSubCommand(os.Args[1], cmd.Root.Commands()) {
		err := rootCmd.Execute()

		if err != nil {
			logrus.Info(err)
		}
	} else if len(os.Args) >= 1 && filepath.Base(os.Args[0]) == "docker-proxy" {
		rootCmd.Help()
	} else {
		cmd.ProxyDockerCommand()
	}
}

func isSubCommand(s string, commands []*cobra.Command) bool {
	var subCommandNames []string
	for _, subCommand := range commands {
		subCommandNames = append(subCommandNames, subCommand.Name())

	}
	return helper.StringInSlice(s, subCommandNames)
}
