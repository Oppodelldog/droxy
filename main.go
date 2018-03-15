package main

import (
	"os"
	"path/filepath"

	"github.com/Oppodelldog/droxy/cmd"

	"github.com/Oppodelldog/droxy/helper"

	"fmt"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockerrun"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = cmd.NewRoot()

	if len(os.Args) >= 2 && isSubCommand(os.Args[1], rootCmd.Commands()) {
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

		dockerRunCommandBuilder := dockerrun.NewCommandBuilder()
		configLoader := config.NewLoader()
		commandresultHandler := cmd.NewCommandResultHandler()
		commandRunner := cmd.NewCommandRunner()
		executableNameParser := cmd.NewExecutableNameParser()
		exitCode := cmd.ExecuteCommand(dockerRunCommandBuilder, configLoader, commandresultHandler, commandRunner, executableNameParser)
		os.Exit(exitCode)
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
