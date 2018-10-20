package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder"
)

// BuildEntryPoint sets the docker entrypoint
func BuildEntryPoint(commandDef *config.CommandDefinition, builder builder.Builder) error {
	if entryPoint, ok := commandDef.GetEntryPoint(); ok {
		builder.SetEntryPoint(entryPoint)
	}

	return nil
}
