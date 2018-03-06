package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockerrun/builder"
	"github.com/Oppodelldog/droxy/helper"
)

// BuildTerminalContext sets -t if terminal context was detected
func BuildTerminalContext(commandDef *config.CommandDefinition, builder builder.Builder) error {
	_ = commandDef
	if helper.IsTerminalContext() {
		builder.AddArgument("-t")
	}

	return nil
}
