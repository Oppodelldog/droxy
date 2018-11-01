package cmd

import (
	"github.com/Oppodelldog/droxy/cmd/proxyexecution"
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand"
)

func executeDroxyCommand(args []string) int {
	dockerRunCommandBuilder := dockercommand.NewCommandBuilder()
	configLoader := config.NewLoader()
	commandResultHandler := proxyexecution.NewCommandResultHandler()
	commandRunner := proxyexecution.NewCommandRunner()
	executableNameParser := proxyexecution.NewExecutableNameParser()

	return proxyexecution.ExecuteCommand(args, dockerRunCommandBuilder, configLoader, commandResultHandler, commandRunner, executableNameParser)
}
