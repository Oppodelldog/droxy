package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercmd/builder"
)

func BuildInteractiveFlag(commandDef *config.CommandDefinition, builder *builder.Builder) error {
	if isInteractive, ok := commandDef.GetIsInteractive(); isInteractive && ok {
		builder.AddArgument("-i")
	}

	return nil
}
