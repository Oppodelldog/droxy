package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercmd/builder"
	"github.com/Oppodelldog/droxy/helper"
)

func BuildTerminalContext(commandDef *config.CommandDefinition, builder builder.Builder) error {
	_ = commandDef
	if helper.IsTerminalContext() {
		builder.AddArgument("-t")
	}

	return nil
}
