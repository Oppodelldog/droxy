package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder"
)

const containerLabel = "droxy"

// LabelContainer labels the container.
func LabelContainer(commandDef *config.CommandDefinition, builder builder.Builder) error {
	_ = commandDef

	builder.AddLabel(containerLabel)

	return nil
}
