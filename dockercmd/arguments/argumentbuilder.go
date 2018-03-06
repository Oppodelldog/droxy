package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercmd/builder"
)

// ArgumentBuilderInterface defines an interface for building an argument of the command line from configuration.
// It's intention is to convert one configuration value into one command line parameter, like for example:
// RemoveContainer:true will be converted into command parameter "--rm".
type ArgumentBuilderInterface interface {
	BuildArgument(commandDef *config.CommandDefinition, builder builder.Builder) error
}
