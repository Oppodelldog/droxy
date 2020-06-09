package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder"
)

// BuildCommand sets the docker command (aka CMD).
func BuildCommand(commandDef config.CommandDefinition, builder builder.Builder) error {
	if command, ok := commandDef.GetCommand(); ok {
		builder.SetCommand(command)
	}

	return nil
}
