package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder"
)

// BuildTerminalContext sets -t if terminal context was detected.
func BuildTerminalContext(commandDef *config.CommandDefinition, builder builder.Builder) error {
	_ = commandDef

	if isTerminalContext() {
		builder.AddArgument("-t")
	}

	return nil
}
