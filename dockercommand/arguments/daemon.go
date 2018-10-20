package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder"
)

//BuildInteractiveFlag sets the interactive flag, which enables user interaction
func BuildInteractiveFlag(commandDef *config.CommandDefinition, builder builder.Builder) error {
	if isInteractive, ok := commandDef.GetIsInteractive(); isInteractive && ok {
		builder.AddArgument("-i")
	}

	return nil
}
