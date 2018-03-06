package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercmd/builder"
)

type ArgumentBuilderInterface interface {
	BuildArgument(commandDef *config.CommandDefinition, builder builder.Builder) error
}
