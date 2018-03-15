package main

import (
	"os"
	"path/filepath"

	"github.com/Oppodelldog/droxy/cmd"

	"github.com/Oppodelldog/droxy/helper"

	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = cmd.Root
	symlinkCommandWrapper := cmd.NewSymlinkCommandWrapper()
	hardlinkCommandWrapper := cmd.NewHardlinkCommandWrapper()
	cloneCommandWrapper := cmd.NewCloneCommandWrapper()

	cmd.Root.AddCommand(symlinkCommandWrapper.GetCommand())
	cmd.Root.AddCommand(hardlinkCommandWrapper.GetCommand())
	cmd.Root.AddCommand(cloneCommandWrapper.GetCommand())

	if len(os.Args) >= 2 && isSubCommand(os.Args[1], cmd.Root.Commands()) {
		err := rootCmd.Execute()
		if err != nil {
			logrus.Info(err)
		}
	} else if len(os.Args) >= 1 && filepath.Base(os.Args[0]) == helper.GetCommandName() {
		err := rootCmd.Help()
		if err != nil {
			logrus.Info(err)
		}
	} else {
		if shallRevealItsDroxy() {
			fmt.Println("YES-IT-IS")
			os.Exit(0)
		}
		cmd.ExecuteCommand()
	}
}

func shallRevealItsDroxy() bool {
	for _, arg := range os.Args {
		if arg == "--is-it-droxy" {

			return true
		}
	}

	return false
}

func isSubCommand(s string, commands []*cobra.Command) bool {
	var subCommandNames []string
	for _, subCommand := range commands {
		subCommandNames = append(subCommandNames, subCommand.Name())

	}
	return helper.StringInSlice(s, subCommandNames)
}
