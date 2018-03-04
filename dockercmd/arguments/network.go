package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercmd/builder"
)

func BuildNetwork(commandDef *config.CommandDefinition, builder *builder.Builder) error {
	if network, ok := commandDef.GetNetwork(); ok {
		builder.SetNetwork(network)
	}

	return nil
}
