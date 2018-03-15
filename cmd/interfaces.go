package cmd

import (
	"os/exec"

	"github.com/Oppodelldog/droxy/config"
)

type (
	//ConfigLoader loads configuration
	ConfigLoader interface {
		Load() *config.Configuration
	}
	//CommandBuilder builds a executable command object
	CommandBuilder interface {
		BuildCommandFromConfig(commandName string, cfg *config.Configuration) (*exec.Cmd, error)
	}
	//CommandRunner runs a command
	CommandRunner interface {
		RunCommand(cmd *exec.Cmd) error
	}
	//CommandResultHandler handles the result of an executed command
	CommandResultHandler interface {
		HandleCommandResult(*exec.Cmd, error) int
	}
	//ExecutableNameParser parsed the name of the current executed file from cli arguments
	ExecutableNameParser interface {
		ParseCommandNameFromCommandLine() string
	}
)
